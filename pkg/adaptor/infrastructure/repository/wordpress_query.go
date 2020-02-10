package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

const (
	listPostsPath              = "/wp-json/wp/v2/posts/"
	listLocationsPath          = "/wp-json/wp/v2/touristSpots/"
	listCategoriesPath         = "/wp-json/wp/v2/categories/"
	listLocationCategoriesPath = "/wp-json/wp/v2/touristSpot__cat/"
	listComicPath              = "/wp-json/wp/v2/comic/"
	listFeaturePath            = "/wp-json/wp/v2/features/"
	listVlogPath               = "/wp-json/wp/v2/vlog/"
	getUserPathFormat          = "/wp-json/wp/v2/users/%d/"
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

func (r *WordpressQueryRepositoryImpl) FindPostsByIDs(ids []int) ([]*wordpress.Post, error) {
	var res []*wordpress.Post
	return res, r.gets(listPostsPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindLocationsByIDs(ids []int) ([]*wordpress.Location, error) {
	var res []*wordpress.Location
	return res, r.gets(listLocationsPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindCategoriesByIDs(ids []int) ([]*wordpress.Category, error) {
	var res []*wordpress.Category
	return res, r.gets(listCategoriesPath, ids, &res)
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

func (r *WordpressQueryRepositoryImpl) FindUserByID(id int) (*wordpress.User, error) {
	u := r.Host
	u.Path = path.Join(u.Path, fmt.Sprintf(getUserPathFormat, id))

	var res wordpress.User
	return &res, r.getJSON(u.String(), &res)
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

	wURL := r.Host
	wURL.Path = path.Join(wURL.Path, wPath)

	q := wURL.Query()
	q.Set("include", util.JoinIntSlice(ids, ","))
	wURL.RawQuery = q.Encode()

	return r.getJSON(wURL.String(), result)
}

func (r *WordpressQueryRepositoryImpl) getJSON(url string, result interface{}) error {
	resp, err := r.Client.Get(url)
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
