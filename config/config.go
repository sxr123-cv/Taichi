package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	MySQL MySQL `toml:"MySQL"`
	Redis Redis `toml:"Redis"`
}
type MySQL struct {
	Usr      string `toml:"usr"`
	Pwd      string `toml:"pwd"`
	Port     string `toml:"port"`
	Addr     string `toml:"addr"`
	Database string `toml:"database"`
}

type Redis struct {
	Pwd  string `toml:"pwd"`
	Port string `toml:"port"`
	Addr string `toml:"addr"`
	Db   int    `toml:"db"`
}

func ReadConfig() (*Config, error) {
	var conf Config
	var filePath = "K:\\Taichi\\taichi.toml"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	err = toml.NewDecoder(file).Decode(&conf)
	return &conf, nil
}
