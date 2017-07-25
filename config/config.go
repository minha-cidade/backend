package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Address        string `yaml:"address"`
	DatabaseInfo   string `yaml:"databaseInfo"`
	AllowedOrigins []string `yaml:"allowedOrigins"`
}

var config Config

func init() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln("Falha lendo arquivo de configuração:", err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalln("Falha processando arquivo de configuração:", err)
	}
}

func Get() Config {
	return config
}
