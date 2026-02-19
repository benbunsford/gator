package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	data := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)

	return data, nil
}
