package config

import (
	"net/url"
	"time"

	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
	yaml "gopkg.in/yaml.v3"
)

type (
	FilePath string

	Config struct {
		Version     string
		Development *Development `validate:"omitempty" yaml:"development"`
		Database    string       `validate:"required" yaml:"database"`
		Migrate     Migrate      `validate:"" yaml:"migrate"`
		Logger      *Logger      `validate:"" yaml:"logger"`
		Stayway     Stayway      `validate:"required" yaml:"stayway"`
		Wordpress   Wordpress    `validate:"required" yaml:"wordpress"`
		AWS         AWS          `validate:"required" yaml:"aws"`
		Slack       Slack        `validate:"required" yaml:"slack"`
		Env         Env

		// scripts配下のスクリプト固有の設定
		Scripts yaml.Node `validate:"" yaml:"scripts"`
	}

	Migrate struct {
		Auto     bool   `yaml:"auto"`
		FilesDir string `yaml:"files_dir"`
	}

	Env struct {
		IsDev bool
		IsStg bool
		IsPrd bool
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

	Wordpress struct {
		BaseURL     URL    `validate:"required" yaml:"base_url"`
		User        string `validate:"required" yaml:"user"`
		Password    string `validate:"required" yaml:"password"`
		CallbackKey string `validate:"required" yaml:"callback_key"`
	}

	AWS struct {
		Endpoint     string        `validate:"" yaml:"endpoint"`
		Region       string        `validate:"required" yaml:"region"`
		FilesBucket  string        `validate:"required" yaml:"files_bucket"`
		UserPoolID   string        `validate:"" yaml:"user_pool_id"`
		ClientID     string        `validate:"" yaml:"client_id"`
		UploadExpire time.Duration `validate:"required" yaml:"upload_expire"`
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
)

func (c Config) IsDev() bool {
	return c.Development != nil
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

func NewDevEnv() Env {
	return Env{
		IsDev: true,
		IsStg: false,
		IsPrd: false,
	}
}

func NewStgEnv() Env {
	return Env{
		IsDev: false,
		IsStg: true,
		IsPrd: false,
	}
}

func NewPrdEnv() Env {
	return Env{
		IsDev: false,
		IsStg: false,
		IsPrd: true,
	}
}
