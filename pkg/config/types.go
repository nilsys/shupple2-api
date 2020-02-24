package config

import (
	"net/url"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v3"
)

type (
	ConfigFilePath string

	Config struct {
		Version     string
		Development *Development  `validate:"omitempty" yaml:"development"`
		Database    string        `validate:"required" yaml:"database"`
		Logger      *Logger       `validate:"" yaml:"logger"`
		Stayway     StaywayConfig `validate:"required" yaml:"stayway"`
		Wordpress   Wordpress     `validate:"required" yaml:"wordpress"`
		AWS         AWS           `validate:"required" yaml:"aws"`

		// scripts配下のスクリプト固有の設定
		Scripts yaml.Node `validate:"" yaml:"scripts"`
	}

	Development struct {
		UserID int `validate:"required" yaml:"user_id"`
	}

	Logger struct {
		JSON  bool
		Level zap.AtomicLevel
	}

	StaywayConfig struct {
		BaseURL string `validate:"required" yaml:"base_url"`
	}

	Wordpress struct {
		BaseURL  URL    `validate:"required" yaml:"base_url"`
		User     string `validate:"required" yaml:"user"`
		Password string `validate:"required" yaml:"password"`
	}

	AWS struct {
		Endpoint     string `validate:"" yaml:"endpoint"`
		Region       string `validate:"required" yaml:"region"`
		AvatarBucket string `validate:"required" yaml:"avatar_bucket"`
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

	parsed, err := url.Parse(str)
	if err != nil {
		return err
	}
	u.URL = *parsed
	return nil
}
