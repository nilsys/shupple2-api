package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/dto"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

const (
	listUsersPath              = "/wp-json/wp/v2/users/"
	listPostsPath              = "/wp-json/wp/v2/posts/"
	listLocationsPath          = "/wp-json/wp/v2/locations/"
	listCategoriesPath         = "/wp-json/wp/v2/categories/"
	listLocationCategoriesPath = "/wp-json/wp/v2/location__cat/"
	listComicPath              = "/wp-json/wp/v2/comic/"
	listFeaturePath            = "/wp-json/wp/v2/features/"
	listVlogPath               = "/wp-json/wp/v2/vlog/"

	maxPerPage = 100
)

type (
	WordpressQueryRepositoryImpl struct {
		config.Wordpress
		Client http.Client
	}
)

var WordpressQueryRepositorySet = wire.NewSet(
	NewWordpressQueryRepositoryImpl,
)

func NewWordpressQueryRepositoryImpl(config config.Wordpress) repository.WordpressQueryRepository {
	return &WordpressQueryRepositoryImpl{
		config,
		http.Client{},
	}
}

func (r *WordpressQueryRepositoryImpl) FindUsersByIDs(ids []int) ([]*wordpress.User, error) {
	var res []*wordpress.User
	return res, r.gets(listUsersPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindPostsByIDs(ids []int) ([]*wordpress.Post, error) {
	var res []*wordpress.Post
	return res, r.gets(listPostsPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindLocationsByIDs(ids []int) ([]*wordpress.Location, error) {
	var res []*wordpress.Location
	return res, r.gets(listLocationsPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindCategoriesByIDs(ids []int) ([]*wordpress.Category, error) {

	var resp dto.WordpressCategories
	if err := r.gets(listCategoriesPath, ids, &resp); err != nil {
		return nil, errors.Wrap(err, "failed to get wordpress category")
	}

	result, err := resp.ToEntities()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert wordpress category dto")
	}
	return result, nil
}

func (r *WordpressQueryRepositoryImpl) FindCategoriesByParentID(parentID, limit int) ([]*wordpress.Category, error) {
	if limit == 0 {
		limit = maxPerPage
	}

	wURL := r.BaseURL
	wURL.Path = path.Join(wURL.Path, listCategoriesPath)

	q := wURL.Query()
	q.Set("parent", fmt.Sprint(parentID))
	q.Set("per_page", fmt.Sprint(limit))
	wURL.RawQuery = q.Encode()

	var resp dto.WordpressCategories
	if err := r.getJSON(wURL.String(), &resp); err != nil {
		return nil, errors.Wrap(err, "failed to get wordpress category")
	}

	result, err := resp.ToEntities()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert wordpress category dto")
	}
	return result, nil
}

func (r *WordpressQueryRepositoryImpl) FindLocationCategoriesByIDs(ids []int) ([]*wordpress.LocationCategory, error) {
	var res []*wordpress.LocationCategory
	return res, r.gets(listLocationCategoriesPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindComicsByIDs(ids []int) ([]*wordpress.Comic, error) {
	var res []*wordpress.Comic
	return res, r.gets(listComicPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindFeaturesByIDs(ids []int) ([]*wordpress.Feature, error) {
	var res []*wordpress.Feature
	return res, r.gets(listFeaturePath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindVlogsByIDs(ids []int) ([]*wordpress.Vlog, error) {
	var res []*wordpress.Vlog
	return res, r.gets(listVlogPath, ids, &res)
}

// http通信するだけなのでどこにでも置けるが、便宜的にココに置く
func (r *WordpressQueryRepositoryImpl) DownloadAvatar(avatarURL string) ([]byte, error) {
	resp, err := r.Client.Get(avatarURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get avatar")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, errors.Wrap(err, "failed to read avatar")
}

func (r *WordpressQueryRepositoryImpl) gets(wPath string, ids []int, result interface{}) error {
	if len(ids) == 0 {
		return nil
	}

	if len(ids) > maxPerPage {
		return serror.New(nil, serror.CodeInvalidParam, "too many ids")
	}

	wURL := r.BaseURL
	wURL.Path = path.Join(wURL.Path, wPath)

	q := wURL.Query()
	q.Set("include", util.JoinIntSlice(ids, ","))
	q.Set("per_page", fmt.Sprint(len(ids)))
	wURL.RawQuery = q.Encode()

	return r.getJSON(wURL.String(), result)
}

func (r *WordpressQueryRepositoryImpl) getJSON(url string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.SetBasicAuth(r.User, r.Password)

	resp, err := r.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to get wordpress resource")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "wordpress returns %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return errors.Wrap(err, "failed to decode json")
	}

	return nil
}
