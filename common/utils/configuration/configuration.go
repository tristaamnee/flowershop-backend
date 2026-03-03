package configuration

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var (
	config  *Configuration
	once    sync.Once
	loadErr error
)

type Configuration struct {
	Db  Db  `json:"db"`
	Jwt Jwt `json:"jwt"`
}

type Db struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Database string `json:"database"`
}

type Jwt struct {
	EncryptionKey string `json:"encrypt_key"`
}

func loadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		loadErr = err
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		loadErr = err
	}
}

func Get(key string) (interface{}, error) {
	once.Do(loadConfig)

	if loadErr != nil {
		return nil, loadErr
	}

	switch key {
	case "db":
		return config.Db, nil
	case "jwt":
		return config.Jwt, nil
	default:
		return nil, errors.New("invalid key")
	}
}
