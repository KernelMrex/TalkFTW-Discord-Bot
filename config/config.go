package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	BotConfig BotConfig `json:"bot"`
}

func LoadConfigFromJsonFile(path string) (*Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return &Configuration{}, err
	}
	conf := &Configuration{}
	if err := json.NewDecoder(file).Decode(conf); err != nil {
		return &Configuration{}, err
	}
	return conf, err
}
