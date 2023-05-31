package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config struct
type Config struct {
	BotToken string `required:"true" split_words:"true"`
	AiToken  string `required:"true" split_words:"true"`
	Channel  int64  `required:"true"`
}

// New config
func New() *Config {
	return &Config{}
}

// GetEnv configuration init
func (cnf *Config) GetEnv() error {
	if err := envconfig.Process("", cnf); err != nil {
		return err
	}
	return nil
}
