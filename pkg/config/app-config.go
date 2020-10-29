package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Debug bool   `envconfig:"DEBUG"`
	Addr  string `envconfig:"ADDR" default:":9095"`
	Stage string `envconfig:"STAGE" default:"dev"`
	//MongoDatasource        	string 	`envconfig:"MONGODATASOURCE"`
}

func ApplicationConfig() (*Config, error) {
	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse environment config.")
	}
	return cfg, nil
}
