package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Address        string `json:"address"`
	DatabaseInfo   string `json:"databaseInfo"`
}

var config Config

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("Falha lendo arquivo de configuração:", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalln("Falha processando arquivo de configuração:", err)
	}
}

func Get() Config {
	return config
}
