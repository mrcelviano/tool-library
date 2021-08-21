package tool_library

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var configuration Configuration

type Configuration struct {
	DataBase     DataBase          `json:"dataBase"`
	AccessConfig AuthConfiguration `json:"accessConfiguration"`
}

func GetConfigurationService() Configuration {
	return configuration
}

func InitConfiguration(filePath string) {
	log.Println("Init configuration...")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Can`t reade configuration file")
		os.Exit(1)
	}
	err = json.Unmarshal(file, &configuration)
	if err != nil {
		log.Println("Can`t unmarshal configuration file to structure")
		os.Exit(1)
	}
	log.Println("Configuration init successful!")
}
