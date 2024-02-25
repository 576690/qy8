package core

import (
	"log"
	"os"

	"github.com/576690/qy8/backend/config"
	"github.com/576690/qy8/backend/global"
	"gopkg.in/yaml.v2"
)

func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("yamlFile.Get err :%s ", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}
