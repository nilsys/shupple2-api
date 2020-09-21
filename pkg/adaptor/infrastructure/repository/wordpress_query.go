package repository

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/reflectwalk"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/dto"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
	"go.uber.org/zap"
)

const (
	usersPath              = "/wp-json/wp/v2/users/"
	postsPath              = "/wp-json/wp/v2/posts/"
	locationsPath          = "/wp-json/wp/v2/locations/"
	postTagsPath           = "/wp-json/wp/v2/tags/"
	categoriesPath         = "/wp-json/wp/v2/categories/"
	locationCategoriesPath = "/wp-json/wp/v2/location__cat/"
	comicPath              = "/wp-json/wp/v2/comic/"
	featurePath            = "/wp-json/wp/v2/features/"
	vlogPath               = "/wp-json/wp/v2/vlog/"
	cfProjectPath          = "/wp-json/wp/v2/cf_project/"
	cfReturnGiftPath       = "/wp-json/wp/v2/cf_return_gift/"
	mediaPath              = "/wp-json/wp/v2/media/"

	maxPerPage = 100

	wpContentPrefix = "/wp-content"

	wpAuthorizationHeader = "WP-Authorization"
)

type (
	WordpressQueryRepositoryImpl struct {
		config.Wordpress
		client           http.Client
		isDev            bool
		wordpressBaseURL string
		mediaBaseURL     string
		filesBaseURL     string
	}
)

var WordpressQueryRepositorySet = wire.NewSet(
	NewWordpressQueryRepositoryImpl,
	wire.Bind(new(repository.WordpressQueryRepository), new(*WordpressQueryRepositoryImpl)),
)

func NewWordpressQueryRepositoryImpl(config *config.Config) *WordpressQueryRepositoryImpl {
	return &WordpressQueryRepositoryImpl{
		config.Wordpress,
		http.Client{},
		config.IsDev(),
		config.Wordpress.BaseURL.String(),
		config.Stayway.Media.BaseURL.String(),
		config.Stayway.Media.FilesURL.String(),
	}
}

func (r *WordpressQueryRepositoryImpl) FindUserByID(id int) (*wordpress.User, error) {
	var res wordpress.User
	return &res, r.getSingleResource(usersPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindPostByID(id int) (*wordpress.Post, error) {
	var res wordpress.Post
	return &res, r.getSingleResource(postsPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindLocationByID(id int) (*wordpress.Location, error) {
	var res wordpress.Location
	return &res, r.getSingleResource(locationsPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindPostTagsByIDs(ids []int) ([]*wordpress.PostTag, error) {
	var res []*wordpress.PostTag

	return res, r.getList(postTagsPath, ids, &res)
}

func (r *WordpressQueryRepositoryImpl) FindCategoryByID(id int) (*wordpress.Category, error) {
	var resp dto.WordpressCategory
	if err := r.getSingleResource(categoriesPath, id, &resp); err != nil {
		return nil, errors.Wrap(err, "failed to get wordpress category")
	}

	return resp.ToEntity()
}

func (r *WordpressQueryRepositoryImpl) FindLocationCategoryByID(id int) (*wordpress.LocationCategory, error) {
	var res wordpress.LocationCategory
	return &res, r.getSingleResource(locationCategoriesPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindComicByID(id int) (*wordpress.Comic, error) {
	var res wordpress.Comic
	return &res, r.getSingleResource(comicPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindFeatureByID(id int) (*wordpress.Feature, error) {
	var res wordpress.Feature
	return &res, r.getSingleResource(featurePath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindVlogByID(id int) (*wordpress.Vlog, error) {
	var res wordpress.Vlog
	return &res, r.getSingleResource(vlogPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindCfProjectByID(id int) (*wordpress.CfProject, error) {
	var res wordpress.CfProject
	return &res, r.getSingleResource(cfProjectPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindCfReturnGiftsByCfProjectID(id int) ([]*wordpress.CfReturnGift, error) {
	wURL := r.BaseURL
	wURL.Path = path.Join(wURL.Path, cfReturnGiftPath)

	q := wURL.Query()
	q.Set("meta[cf_project]", fmt.Sprint(id))
	wURL.RawQuery = q.Encode()

	var result []*wordpress.CfReturnGift
	return result, errors.Wrap(r.GetJSON(wURL.String(), &result), "failed to get cf_return_gifts associated to cf_project")
}

func (r *WordpressQueryRepositoryImpl) FindCfReturnGiftByID(id int) (*wordpress.CfReturnGift, error) {
	var res wordpress.CfReturnGift
	return &res, r.getSingleResource(cfReturnGiftPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FindMediaByID(id int) (*wordpress.Media, error) {
	var res wordpress.Media
	return &res, r.getSingleResource(mediaPath, id, &res)
}

func (r *WordpressQueryRepositoryImpl) FetchMediaBodyByID(id int) (*model.MediaBody, error) {
	media, err := r.FindMediaByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get media body(id=%d)", id)
	}

	return r.FetchResource(media.SourceURL)
}

func (r *WordpressQueryRepositoryImpl) FetchResource(avatarURL string) (*model.MediaBody, error) {
	resp, err := r.httpGet(avatarURL, false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get avatar")
	}

	return &model.MediaBody{
		ContentType: resp.Header.Get(echo.HeaderContentType),
		Body:        resp.Body,
	}, nil
}

func (r *WordpressQueryRepositoryImpl) getList(wPath string, ids []int, result interface{}) error {
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

	return errors.Wrapf(r.GetJSON(wURL.String(), result), "failed to get %s", wURL.Path)
}

func (r *WordpressQueryRepositoryImpl) getSingleResource(wPath string, id int, result interface{}) error {
	wURL := r.BaseURL
	wURL.Path = path.Join(wURL.Path, wPath, strconv.Itoa(id))

	return errors.Wrapf(r.GetJSON(wURL.String(), result), "failed to get %s", wURL.Path)
}

func (r *WordpressQueryRepositoryImpl) GetJSON(url string, result interface{}) error {
	// query stringがあるとキャッシュが無効になるように設定されているのでここで付与
	if strings.Contains(url, "?") {
		url += "&cache_busting"
	} else {
		url += "?cache_busting"
	}

	resp, err := r.httpGet(url, true)
	if err != nil {
		return errors.Wrap(err, "failed to get wordpress resource")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return serror.New(nil, serror.CodeNotFound, "not found")
		}
		return serror.New(nil, serror.CodeUndefined, "wordpress returns %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return errors.Wrap(err, "failed to decode json")
	}

	if err := r.replaceURLs(result); err != nil {
		return errors.Wrap(err, "failed to replace urls in wordpress resource")
	}

	return nil
}

func (r *WordpressQueryRepositoryImpl) httpGet(url string, apiAuth bool) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	if strings.HasPrefix(url, r.wordpressBaseURL) {
		req.SetBasicAuth(r.BasicAuth.Username, r.BasicAuth.Password)
	}

	if apiAuth {
		/* NOTE:
		WPにはアクセス制限のためのBasic認証とAPIのユーザー認証がある
		これらは同じヘッダAuthorizationを使用するため、conflictしてしまう。
		そこで、ユーザー認証のためのAuthorizationヘッダはWP-Authorizationヘッダとして送り、nginxでBasic認証後にAuthorizationヘッダに乗せるという処理をしている
		*/
		apiAuthValue := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(r.APIAuth.Username, ":", r.APIAuth.Password)))
		req.Header.Set(wpAuthorizationHeader, "Basic "+apiAuthValue)
	}

	if r.isDev {
		dump, _ := httputil.DumpRequest(req, false)
		logger.Debug("httpn response", zap.String("req", string(dump)))
	}

	resp, err := r.client.Do(req)

	if r.isDev && err == nil {
		dump, _ := httputil.DumpResponse(resp, false)
		logger.Debug("http response", zap.String("res", string(dump)))
	}

	return resp, err
}

// Wordpressが/を\/にエスケープして返すせいでjsonパース後にしかURLを置換できない
func (r *WordpressQueryRepositoryImpl) replaceURLs(v interface{}) error {
	return reflectwalk.Walk(v, wordpressReplaceURLsWalker{r})
}

func (r *WordpressQueryRepositoryImpl) replaceDomain(str string) string {
	return strings.ReplaceAll(str, r.wordpressBaseURL, r.mediaBaseURL)
}

// 移行前後は、負荷の問題でpluginのURL置換機能をonにできないので、こちら側で置換する。
// TODO: WP Offload Media LiteのURL置換機能をonにして、このメソッドを削除する
var (
	mediaURLRegexp = regexp.MustCompile(`https://([-a-z]+\.)?stayway.jp/tourism/wp-content/uploads/\S+\.[A-Za-z]+`)
)

func (r *WordpressQueryRepositoryImpl) replaceMediaURL(str string) string {
	return mediaURLRegexp.ReplaceAllStringFunc(str, func(url string) string {
		start := strings.Index(url, wpContentPrefix)
		if start < 0 { // ありえないが念の為
			return url
		}
		url = url[start:]

		return r.filesBaseURL + url
	})
}

type wordpressReplaceURLsWalker struct {
	*WordpressQueryRepositoryImpl
}

func (w wordpressReplaceURLsWalker) Struct(reflect.Value) error {
	return nil
}

func (w wordpressReplaceURLsWalker) StructField(field reflect.StructField, fv reflect.Value) error {
	if field.Type.Kind() != reflect.String || !fv.CanSet() {
		return nil
	}

	fv.SetString(w.replaceMediaURL(w.replaceDomain(fv.String())))
	return nil
}
