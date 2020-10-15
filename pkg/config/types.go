package config

import (
	"net/url"

	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

type (
	FilePath string

	Config struct {
		Env         Env
		Version     string       `validate:"required" yaml:"version"`
		Development *Development `validate:"omitempty" yaml:"development"`
		Database    string       `validate:"required" yaml:"database"`
		Shupple     Shupple      `validate:"required" yaml:"shupple"`
		Migrate     Migrate      `validate:"" yaml:"migrate"`
		Logger      *Logger      `validate:"" yaml:"logger"`
	}

	Migrate struct {
		Auto     bool   `yaml:"auto"`
		FilesDir string `yaml:"files_dir"`
	}

	Development struct {
		FirebaseID string `validate:"required" yaml:"firebase_id"`
	}

	Logger struct {
		JSON  bool
		Level zap.AtomicLevel
	}

	Shupple struct {
		FilesURL URL `validate:"required" yaml:"files_url"`
	}

	URL struct {
		url.URL
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
