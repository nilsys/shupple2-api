package config

import (
	"os"

	"github.com/pkg/errors"

	validator "gopkg.in/go-playground/validator.v9"
	yaml "gopkg.in/yaml.v2"
)

// build時にldflagで注入する
var Version = "unknown"

func GetConfig(filename ConfigFilePath) (*Config, error) {
	f, err := os.Open(string(filename))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	var config Config
	if err = yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := validator.New().Struct(&config); err != nil {
		return nil, err
	}

	config.Version = Version

	return &config, nil
}
