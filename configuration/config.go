package configuration

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Port int `envconfig:"PORT" `
}

func NewConfig() (*Config, error) {
	var cfg Config
	fmt.Println(cfg)
	fmt.Println(&cfg)

	err := envconfig.Process("myapp", &cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to build config from env")
	}
	return &cfg, nil
}
