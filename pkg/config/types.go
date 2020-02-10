package config

import (
	"net/url"

	"go.uber.org/zap"
)

type (
	ConfigFilePath string

	Config struct {
		Version     string
		Development *Development `validate:"omitempty" yaml:"development"`
		Database    string       `validate:"required" yaml:"database"`
		Logger      *Logger      `validate:"" yaml:"logger"`
		Wordpress   Wordpress    `validate:"required" yaml:"wordpress"`
		AWS         AWS          `validate:"required" yaml:"aws"`
	}

	Development struct {
		UserID int `validate:"required" yaml:"user_id"`
	}

	Logger struct {
		JSON  bool
		Level zap.AtomicLevel
	}

	Wordpress struct {
		Host url.URL `validate:"required" yaml:"host"`
	}

	AWS struct {
		Endpoint     string `validate:"" yaml:"endpoint"`
		Region       string `validate:"required" yaml:"region"`
		AvatarBucket string `validate:"required" yaml:"avatar_bucket"`
	}
)

func (c Config) IsDev() bool {
	return c.Development != nil
}
