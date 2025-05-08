package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
)

type ConfigLoader interface {
	Load(path string, cfg any) error
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Path   string `yaml:"path"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// --------- CleanENV ------------

type CleanenvLoader struct {}

func (l CleanenvLoader) Load(path string, cfg any) error {
	return cleanenv.ReadConfig(path, cfg)
}

// -------- Viper ---------

type ViperLoader struct {}

func (l ViperLoader) Load(path string, cfg any) error {
	v := viper.New()

	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	return v.Unmarshal(cfg)
}

// ------ init -------

func InitConfig(loader ConfigLoader, path string) *Config {
	var cfg Config

	if err := loader.Load(path, &cfg); err != nil {
		panic("ошибка загрузки конфигурации: " + err.Error())
	}

	return &cfg
}
