package service

import (
	"bytes"
	"io/ioutil"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	SitemapService interface {
		Show(path string) ([]byte, string, error)
	}

	SitemapServiceImpl struct {
		repository.WordpressQueryRepository
		WordpressConfig config.Wordpress
		MediaConfig     config.StaywayMedia
	}
)

var SitemapServiceSet = wire.NewSet(
	wire.Struct(new(SitemapServiceImpl), "*"),
	wire.Bind(new(SitemapService), new(*SitemapServiceImpl)),
)

func (s *SitemapServiceImpl) Show(path string) ([]byte, string, error) {
	origURL := s.WordpressConfig.BaseURL
	origURL.Path = path // prefixがpathに入っているのでJoinではなく代入で良い

	file, err := s.WordpressQueryRepository.FetchResource(origURL.String())
	if err != nil {
		return nil, "", errors.Wrapf(err, "failed to get sitemap(%s)", path)
	}
	defer file.Body.Close()

	body, err := ioutil.ReadAll(file.Body)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to read sitemap body")
	}

	body = bytes.ReplaceAll(body, s.WordpressConfig.BaseURL.Byte(), s.MediaConfig.BaseURL.Byte())

	return body, file.ContentType, nil
}
