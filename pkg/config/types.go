package config

import (
	"net/url"
	"time"

	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v3"
)

type (
	FilePath string

	Config struct {
		Env             Env
		Version         string          `validate:"required" yaml:"version"`
		Development     *Development    `validate:"omitempty" yaml:"development"`
		Database        string          `validate:"required" yaml:"database"`
		Migrate         Migrate         `validate:"" yaml:"migrate"`
		Logger          *Logger         `validate:"" yaml:"logger"`
		Stayway         Stayway         `validate:"required" yaml:"stayway"`
		Wordpress       Wordpress       `validate:"required" yaml:"wordpress"`
		AWS             AWS             `validate:"required" yaml:"aws"`
		Slack           Slack           `validate:"required" yaml:"slack"`
		GoogleAnalytics GoogleAnalytics `validate:"required" yaml:"google_analytics"`
		Payjp           Payjp           `validate:"required" yaml:"payjp"`
		CfProject       CfProject       `validate:"required" yaml:"cf_project"`
		Facebook        Facebook        `validate:"" yaml:"facebook"`

		// scripts配下のスクリプト固有の設定
		Scripts yaml.Node `validate:"" yaml:"scripts"`
	}

	Migrate struct {
		Auto     bool   `yaml:"auto"`
		FilesDir string `yaml:"files_dir"`
	}

	Development struct {
		CognitoID string `validate:"required" yaml:"cognito_id"`
	}

	Logger struct {
		JSON  bool
		Level zap.AtomicLevel
	}

	Stayway struct {
		Metasearch StaywayMetasearch `validate:"required" yaml:"metasearch"`
		Media      StaywayMedia      `validate:"required" yaml:"media"`
	}

	StaywayMetasearch struct {
		BaseURL URL `validate:"required" yaml:"base_url"`
	}

	StaywayMedia struct {
		BaseURL  URL `validate:"required" yaml:"base_url"`
		FilesURL URL `validate:"required" yaml:"files_url"`
	}

	Auth struct {
		Username string `validate:"required" yaml:"username"`
		Password string `validate:"required" yaml:"password"`
	}

	Wordpress struct {
		BaseURL     URL    `validate:"required" yaml:"base_url"`
		BasicAuth   Auth   `validate:"required" yaml:"basic_auth"`
		APIAuth     Auth   `validate:"required" yaml:"api_auth"`
		CallbackKey string `validate:"required" yaml:"callback_key"`
	}

	AWS struct {
		Endpoint          string        `validate:"" yaml:"endpoint"`
		Region            string        `validate:"required" yaml:"region"`
		FilesBucket       string        `validate:"required" yaml:"files_bucket"`
		UserPoolID        string        `validate:"required" yaml:"user_pool_id"`
		ClientID          string        `validate:"required" yaml:"client_id"`
		UploadExpire      time.Duration `validate:"required" yaml:"upload_expire"`
		FromEmail         string        `validate:"required" yaml:"from_email"`
		PersistMediaQueue string        `validate:"required" yaml:"persist_media_queue"`
	}

	GoogleAnalytics struct {
		ViewID int `validate:"required" yaml:"view_id"`
	}

	// TODO: 他のアプリが追加されると思うからSlackの下にアプリ毎にconfig作る
	Slack struct {
		Token         string `validate:"required" yaml:"token"`
		CallbackKey   string `validate:"required" yaml:"callback_key"`
		ReportChannel string `validate:"required" yaml:"report_channel"`
	}

	URL struct {
		url.URL
	}

	Payjp struct {
		SecretKey string `validate:"required" yaml:"secret_key"`
	}

	// MEMO: facebookアプリが増えた場合には追記
	Facebook struct {
		BatchApp FacebookApp `validate:"" yaml:"batch"`
	}

	FacebookApp struct {
		AppID     string `validate:"" yaml:"app_id"`
		AppSecret string `validate:"" yaml:"app_secret"`
	}

	CfProject struct {
		SystemFee int `validate:"required" yaml:"system_fee"`
	}
)

func (c Config) IsDev() bool {
	return c.Development != nil
}

func (u *URL) Byte() []byte {
	return []byte(u.String())
}

func (u *URL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if str == "" {
		return nil
	}

	parsed, err := url.Parse(str)
	if err != nil {
		return err
	}
	u.URL = *parsed
	return nil
}

func URLRequiredValidation(sl validator.StructLevel) {
	u := sl.Current().Interface().(URL)

	zero := url.URL{}
	if u.URL == zero {
		sl.ReportError(u, "URL", "URL", "required", "")
	}
}
