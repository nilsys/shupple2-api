package config

import "go.uber.org/zap"

type (
	ConfigFilePath string

	Config struct {
		Version     string
		Development *Development `validate:"omitempty" yaml:"development"`
		Database    string       `validate:"required" yaml:"database"`
		Logger      *Logger      `validate:"" yaml:"logger"`
	}

	Development struct {
		UserID int `validate:"required" yaml:"user_id"`
	}

	Logger struct {
		JSON  bool
		Level zap.AtomicLevel
	}
)

func (c Config) IsDev() bool {
	return c.Development != nil
}
