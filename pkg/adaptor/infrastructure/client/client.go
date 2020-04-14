package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/pkg/errors"

	"go.uber.org/zap"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
)

// HTTPClient
type (
	Client interface {
		GetJSON(url string, opt *Option, v interface{}) error
	}

	clientImpl struct {
		client http.Client
	}

	Option struct {
		QueryParams url.Values
		Headers     url.Values
		BasicAuth   *BasicAuth
	}

	BasicAuth struct {
		Username string `validate:"required"`
		Password string `validate:"required"`
	}

	Config struct {
		Timeout time.Duration
	}
)

func NewClient(c *Config) Client {
	if c == nil {
		c = &Config{}
	}
	return &clientImpl{
		client: http.Client{
			Timeout: c.Timeout,
		},
	}
}

// TODO: エラー時
func (c *clientImpl) GetJSON(url string, opt *Option, v interface{}) error {
	if opt == nil {
		opt = &Option{}
	}

	req, err := c.createRequest(http.MethodGet, url, nil, opt)
	if err != nil {
		return errors.Wrapf(err, "failed to create get request to %s", url)
	}

	resp, err := c.do(req)
	if err != nil {
		return errors.Wrap(err, "get request failed")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(v); err != nil {
		return errors.Wrap(err, "failed to parse json")
	}

	return nil
}

func (c *clientImpl) createRequest(method string, url string, body io.Reader, opt *Option) (*http.Request, error) {
	url = c.buildURL(url, opt.QueryParams)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	if opt.BasicAuth != nil {
		req.SetBasicAuth(opt.BasicAuth.Username, opt.BasicAuth.Password)
	}

	for k, vs := range opt.Headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	return req, nil
}

func (c *clientImpl) do(req *http.Request) (*http.Response, error) {
	logger.Debug("http request",
		zap.String("method", req.Method),
		zap.String("host", req.URL.Host),
		zap.String("path", req.URL.Path),
		zap.Reflect("query", req.URL.Query()),
		zap.Reflect("headers", req.Header))

	res, err := c.client.Do(req)
	if err != nil {
		// MEMO: code仮置き
		return nil, serror.New(err, serror.CodeUndefined, "Failed http")
	}

	logger.Debug("http response",
		zap.String("status", res.Status),
		zap.Reflect("headers", res.Header),
	)

	if res.StatusCode/100 != 2 {
		return res, c.err(res)
	}

	return res, nil
}

func (c *clientImpl) buildURL(url string, queryParams url.Values) string {
	added := queryParams.Encode()
	if strings.Contains(url, "?") {
		return url + "&" + added
	}

	return url + "?" + added
}

func (c *clientImpl) err(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	if resp.StatusCode/100 == 4 {
		// MEMO: code仮置き
		return serror.New(err, serror.CodeNotFound, string(body))
	}

	// MEMO: code仮置き
	return serror.New(err, serror.CodeNotFound, string(body))
}
