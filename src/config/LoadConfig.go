package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Cfg Config

func LoadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	d := yaml.NewDecoder(file)

	if err = d.Decode(&Cfg); err != nil {
		return err
	}

	return nil
}
