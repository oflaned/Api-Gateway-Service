package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Port    string        `yaml:"port"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Username    string `json:"username"`
	Passwd      string `json:"passwd"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Database    string `json:"database"`
	MaxAttempts int    `json:"max_attempts"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			cleanenv.GetDescription(instance, nil)
		}
	})
	return instance
}
