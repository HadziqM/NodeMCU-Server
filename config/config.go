package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbUrl string `json:"db_url"`
	Host  string `json:"host"`
	Use   bool   `json:"db_use"`
}

type Sensor struct {
	Val    string `json:"value"`
	Device string `json:"device"`
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
func LogIgnore(e error) {
	if e != nil {
		log.Println(e)
	}
}
