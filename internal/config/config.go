package book_inventory_system_config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerAddress string `yaml:"server_address"`
	Admins        string `yaml:"admins"`
	Authors       string `yaml:"authors"`
	Books         string `yaml:"books"`
	Genres        string `yaml:"genres"`
	Instances     string `yaml:"instances"`
	Languages     string `yaml:"languages"`
	Productions   string `yaml:"productions"`
	Readers       string `yaml:"readers"`
	Users         string `yaml:"users"`
}

func New(cfgPath string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(cfgPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
