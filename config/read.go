package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() (Config, error) {
	var config Config

	// Open YAML file
	file, err := os.Open(os.Getenv("CONFIG_PATH")) //todo: add env var to docker-compose
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			return Config{}, err
		}
	}

	return config, nil
}
