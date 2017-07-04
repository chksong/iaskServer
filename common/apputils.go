package common

import (
	"os"
	"log"
	"encoding/json"
)

type configuration struct {
	Server string
}

//////////////////////////////////////////////



var AppConfig configuration

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()

	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
		return
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)

	if err != nil {
		log.Fatalf("[loadAppConfig] := %s\n", err)
	}

	log.Println(AppConfig)
}


