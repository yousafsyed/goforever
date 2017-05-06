package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Run      Command `json:"Run"`
	LogsFile string  `json:"LogsFile"`
}

type Command struct {
	StartCommand string `json:"StartCommand"`
	EndCommand   string `json:"EndCommand"`
}

var config Config

func init() {
	configEnv := os.Getenv("GOFOREVER_CONFIG")

	if configEnv == "" {
		configEnv = "./config.json"
	}

	configFile, err := ioutil.ReadFile(string(configEnv))
	if err != nil {
		checkErr(err)
	}
	err2 := json.Unmarshal(configFile, &config)
	if err2 != nil {
		fmt.Println("whoops:", err)
	}
}

func checkErr(err error) {
	if err != nil {
		logOut(string(err.Error()))
	}
}
