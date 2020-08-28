package service

import (
	"bytes"
	"fmt"
	"io/ioutil"

	path2 "path"

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

const (
	cfProjectPathPrefix    = "/projects"
	fixTargetCfProjectPath = "/project"
)

var SitemapServiceSet = wire.NewSet(
	wire.Struct(new(SitemapServiceImpl), "*"),
	wire.Bind(new(SitemapService), new(*SitemapServiceImpl)),
)

func (s *SitemapServiceImpl) Show(path string) ([]byte, string, error) {
	origURL := s.WordpressConfig.BaseURL
	fmt.Println(origURL)
	origURL.Path = path // prefixがpathに入っているのでJoinではなく代入で良い
	fmt.Println(origURL)

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

	// CfProjectは/tourismのprefixが必要ない
	cfProjectURL := s.MediaConfig.BaseURL
	cfProjectURL.Path = cfProjectPathPrefix

	// 修正対象
	fixTargetCfProjectURL := s.MediaConfig.BaseURL
	fixTargetCfProjectURL.Path = path2.Join(fixTargetCfProjectURL.Path, fixTargetCfProjectPath)
	body = bytes.ReplaceAll(body, fixTargetCfProjectURL.Byte(), cfProjectURL.Byte())

	return body, file.ContentType, nil
}
