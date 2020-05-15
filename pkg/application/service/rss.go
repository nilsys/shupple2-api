package service

import (
	"bytes"
	"io/ioutil"
	"net/url"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	RSSService interface {
		Show() ([]byte, string, error)
	}

	RSSServiceImpl struct {
		repository.WordpressQueryRepository
		WordpressConfig config.Wordpress
		MediaConfig     config.StaywayMedia
	}
)

var RssServiceSet = wire.NewSet(
	wire.Struct(new(RSSServiceImpl), "*"),
	wire.Bind(new(RSSService), new(*RSSServiceImpl)),
)

const (
	wpContentPath = "/wp-content/uploads"
	feedQuery     = "?feed=smartnews"
)

func (s *RSSServiceImpl) Show() ([]byte, string, error) {
	// /tourism?feed=smartnews
	originURL := s.WordpressConfig.BaseURL
	q, _ := url.ParseQuery(originURL.RawQuery)
	q.Add("feed", "smartnews")
	originURL.RawQuery = q.Encode()

	file, err := s.WordpressQueryRepository.FetchResource(originURL.String())
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to get smartnews rss")
	}
	defer file.Body.Close()

	body, err := ioutil.ReadAll(file.Body)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to read smartnews rss body")
	}

	//https://stayway.jp/tourism/wp-content/uploads/2020/05/DSC_0326.ぶれ修正jpg.jpg -> https://files.stayway.jp/wp-content/uploads/2020/05/DSC_0326.ぶれ修正jpg.jpg
	body = bytes.ReplaceAll(body, []byte(s.MediaConfig.BaseURL.String()+wpContentPath), []byte(s.MediaConfig.FilesURL.String()+wpContentPath))
	//https://admin.stayway.jp/tourism/wp-content/uploads/2018/04/AdobeStock_302203471.jpeg -> https://files.stayway.jp/wp-content/uploads/2018/04/AdobeStock_302203471.jpeg
	body = bytes.ReplaceAll(body, []byte(s.WordpressConfig.BaseURL.String()+wpContentPath), []byte(s.MediaConfig.FilesURL.String()+wpContentPath))
	//https://admin.stayway.jp/tourism/?feed=smartnews の ?feed=smartnews -> /smartnews
	body = bytes.ReplaceAll(body, []byte(feedQuery), []byte("/smartnews"))
	//https://admin.stayway.jp/tourism -> https://stayway.jp/tourism
	body = bytes.ReplaceAll(body, s.WordpressConfig.BaseURL.Byte(), s.MediaConfig.BaseURL.Byte())
	return body, file.ContentType, nil
}
