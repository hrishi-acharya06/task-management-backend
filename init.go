package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	services "task-management/services"
)

func init() {
	file, err := os.Open("go_config.json")
	if err != nil {
		panic(err)
	}
	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	configMap := make(map[string]string)
	er := json.Unmarshal(byteData, &configMap)
	if er != nil {
		panic(er)
	}

	for k, v := range configMap {
		os.Setenv(k, v)
	}
	services.InitializeDB()
	services.MigrateDatabase()

}
