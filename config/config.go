package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbUrl string `json:"db_url"`
	Host  string `json:"host"`
}

func Load() (Config, error) {
	var config Config
	file, err := os.ReadFile("./config.json")
	if err != nil {
		return config, err
	}
	json.Unmarshal(file, &config)
	return config, nil
}

func LogErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}