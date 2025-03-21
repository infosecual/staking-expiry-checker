package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Pollers PollersConfig `mapstructure:"pollers"`
	Db      DbConfig      `mapstructure:"db"`
	Btc     BTCConfig     `mapstructure:"btc"`
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func (cfg *Config) Validate() error {
	if err := cfg.Pollers.Validate(); err != nil {
		return err
	}

	if err := cfg.Db.Validate(); err != nil {
		return err
	}

	if err := cfg.Btc.Validate(); err != nil {
		return err
	}

	if err := cfg.Metrics.Validate(); err != nil {
		return err
	}

	return nil
}

// New returns a fully parsed Config object from a given file directory
func New(cfgFile string) (*Config, error) {
	_, err := os.Stat(cfgFile)
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()
	/*
		Below code will replace nested fields in yml into `_` and any `-` into `__` when you try to override this config via env variable
		To give an example:
		1. `some.config.a` can be overriden by `SOME_CONFIG_A`
		2. `some.config-a` can be overriden by `SOME_CONFIG__A`
		This is to avoid using `-` in the environment variable as it's not supported in all os terminal/bash
		Note: vipner package use `.` as delimitter by default. Read more here: https://pkg.go.dev/github.com/spf13/viper#readme-accessing-nested-keys
	*/
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "__"))

	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err = cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
