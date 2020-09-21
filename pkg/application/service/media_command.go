package service

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"gopkg.in/vansante/go-ffprobe.v2"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

const (
	ffprobeTimeout    = 3 * time.Second
	ffmpegTimeout     = 3 * time.Minute
	imIdentifyTimeout = 1 * time.Second
	imConvertTimeout  = 10 * time.Second
)

type (
	MediaCommandService interface {
		PreparePersist(uuid, dest string, mediaType model.MediaType) error
		Persist(request *model.PersistMediaRequest) error
	}

	MediaCommandServiceImpl struct {
		AWSConfig  config.AWS
		AWSSession *session.Session
		repository.MediaQueryRepository
		repository.MediaCommandRepository
	}
)

var MediaCommandServiceSet = wire.NewSet(
	wire.Struct(new(MediaCommandServiceImpl), "*"),
	wire.Bind(new(MediaCommandService), new(*MediaCommandServiceImpl)),
)

func (s *MediaCommandServiceImpl) PreparePersist(uuid, dest string, mediaType model.MediaType) error {
	return s.MediaCommandRepository.SavePersistRequest(&model.PersistMediaRequest{
		UUID:        uuid,
		Destination: dest,
		MediaType:   mediaType,
	})
}

// NOTE: linuxにゴリゴリ依存するので抽象化諦めた
func (s *MediaCommandServiceImpl) Persist(request *model.PersistMediaRequest) error {
	uploadedMedia, err := s.MediaQueryRepository.GetUploadedMedia(request.UUID)
	if err != nil {
		return errors.Wrap(err, "failed to download uploaded media")
	}
	defer uploadedMedia.Body.Close()

	tmpfile, err := ioutil.TempFile("", "media")
	if err != nil {
		return errors.Wrap(err, "failed to create temporary file to download media")
	}
	defer os.Remove(tmpfile.Name())

	if _, err := io.Copy(tmpfile, uploadedMedia.Body); err != nil {
		return errors.Wrap(err, "failed to write uploaded video to file")
	}
	if err := s.rewind(tmpfile); err != nil {
		return errors.Wrap(err, "failed to rewind file")
	}

	resizedMedia, err := s.resize(request.UUID, tmpfile, request.MediaType)
	if err != nil {
		return errors.Wrap(err, "failed to resize media")
	}
	defer resizedMedia.Close()

	mediaBody := &model.MediaBody{
		ContentType: uploadedMedia.ContentType,
		Body:        resizedMedia,
	}
	return s.MediaCommandRepository.Save(mediaBody, request.Destination)
}

func (s *MediaCommandServiceImpl) resize(uuid string, file *os.File, mediaType model.MediaType) (io.ReadCloser, error) {
	switch mediaType {
	case model.MediaTypeUserIcon, model.MediaTypeUserHeader, model.MediaTypeReviewImage:
		return s.resizeImage(uuid, file, model.MaxMediaSize(mediaType))
	case model.MediaTypeReviewVideo:
		return s.resizeVideo(uuid, file, model.MaxMediaSize(mediaType))
	}

	return nil, errors.Errorf("unsupported media: %s", mediaType)
}

func (s *MediaCommandServiceImpl) resizeImage(uuid string, src *os.File, maxSize model.Size) (io.ReadCloser, error) {
	origSize, err := s.getImageSize(src)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get image size")
	}

	resultSize, resized := s.decideResultSize(*origSize, maxSize)
	if !resized {
		if err := s.rewind(src); err != nil {
			return nil, errors.Wrap(err, "failed to rewind file")
		}

		return src, nil
	}
	logger.Debug("image resizing", zap.String("uuid", uuid))

	imDest := src.Name() + ".resized"

	ctx, cancelFn := context.WithTimeout(context.Background(), imConvertTimeout)
	defer cancelFn()

	cmd := exec.CommandContext(ctx, "convert", "-resize", resultSize.JoinByX()+"!", src.Name(), imDest)
	if err := cmd.Run(); err != nil {
		return nil, s.handleCmdError(ctx, cmd, err)
	}

	return os.Open(imDest)
}

func (s *MediaCommandServiceImpl) getImageSize(file *os.File) (*model.Size, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), imIdentifyTimeout)
	defer cancelFn()

	cmd := exec.CommandContext(ctx, "identify", "-format", "%[width],%[height]", file.Name()+"[0]")
	output, err := cmd.Output()
	if err != nil {
		return nil, s.handleCmdError(ctx, cmd, err)
	}

	parts := strings.Split(string(output), ",")
	if len(parts) != 2 {
		return nil, errors.Errorf("strange identify output: %s", output)
	}

	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.Wrapf(err, "strange identify output: %s", output)
	}

	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.Wrapf(err, "strange identify output: %s", output)
	}

	return &model.Size{
		Width:  uint(width),
		Height: uint(height),
	}, nil
}

func (s *MediaCommandServiceImpl) resizeVideo(uuid string, src *os.File, maxSize model.Size) (io.ReadCloser, error) {
	origSize, extension, err := s.getVideoSizeAndExtension(src)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get video size")
	}

	resultSize, resized := s.decideResultSize(*origSize, maxSize)
	if !resized {
		if err := s.rewind(src); err != nil {
			return nil, errors.Wrap(err, "failed to rewind file")
		}

		return src, nil
	}
	logger.Debug("video resizing", zap.String("uuid", uuid))

	// H264は縦横が偶数でないといけないらしいので、ファイル形式に依らず一律で変えてしまう
	resultSize.Width = s.truncateToEven(resultSize.Width)
	resultSize.Height = s.truncateToEven(resultSize.Height)

	ffmpegDest := src.Name() + ".compressed"

	ctx, cancelFn := context.WithTimeout(context.Background(), ffmpegTimeout)
	defer cancelFn()

	cmd := exec.CommandContext(ctx, "ffmpeg", "-i", src.Name(), "-s", resultSize.JoinByX(), "-f", extension, ffmpegDest)
	if err := cmd.Run(); err != nil {
		return nil, s.handleCmdError(ctx, cmd, err)
	}

	return os.Open(ffmpegDest)
}

func (s *MediaCommandServiceImpl) getVideoSizeAndExtension(file *os.File) (*model.Size, string, error) {
	if err := s.rewind(file); err != nil {
		return nil, "", errors.Wrap(err, "failed to rewind video file")
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), ffprobeTimeout)
	defer cancelFn()

	videoAttributes, err := ffprobe.ProbeReader(ctx, file)
	if err != nil {
		return nil, "", errors.Wrap(err, "ffprobe failed")
	}

	videoStreams := videoAttributes.StreamType(ffprobe.StreamVideo)
	if len(videoStreams) == 0 {
		return nil, "", errors.New("no video stream")
	}
	size := &model.Size{
		Width:  uint(videoStreams[0].Width),
		Height: uint(videoStreams[0].Height),
	}

	formats := strings.Split(videoAttributes.Format.FormatName, ",")
	if len(formats) == 0 {
		return nil, "", errors.New("unknown video formats")
	}

	return size, formats[0], nil
}

func (s *MediaCommandServiceImpl) handleCmdError(ctx context.Context, cmd *exec.Cmd, err error) error {
	logger.Debug("failed ffmpeg command", zap.Stringer("cmd", cmd))

	if err := ctx.Err(); err != nil {
		logger.Error("context failure; maybe timeout", zap.Error(err))
	}

	if exitErr, ok := err.(*exec.ExitError); ok && len(exitErr.Stderr) > 0 {
		logger.Error("ffmpeg exited with failure", zap.String("error", string(exitErr.Stderr)))
	}

	return errors.Wrap(err, "ffmpeg failed")
}

func (s *MediaCommandServiceImpl) rewind(file *os.File) error {
	_, err := file.Seek(0, io.SeekStart)
	return err
}

/*
* 縦横それぞれを最大サイズに収めるための比率をそれぞれ出し、小さい方の比率を元サイズにかける。
* ratio = min(maxW/origW, maxH/origH)
* resultW = origW * ratio, resultH = origH * ratio
* これをfloatの計算を避けて行う
 */
func (s *MediaCommandServiceImpl) decideResultSize(origSize, maxSize model.Size) (retSize model.Size, resized bool) {
	if origSize.Width <= maxSize.Width && origSize.Height <= maxSize.Height {
		return origSize, false
	}

	if maxSize.Width*origSize.Height < maxSize.Height*origSize.Width { // <=> maxW/origW < maxH/origH
		resultH := origSize.Height * maxSize.Width / origSize.Width
		return model.Size{
			Width:  maxSize.Width,
			Height: resultH,
		}, true
	}

	resultW := origSize.Width * maxSize.Height / origSize.Height
	return model.Size{
		Width:  resultW,
		Height: maxSize.Height,
	}, true
}

func (s *MediaCommandServiceImpl) truncateToEven(i uint) uint {
	if i%2 == 0 {
		return i
	}
	return i - 1
}
