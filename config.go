package main

import (
	"encoding/json"
	"os"
)

type config struct {
	Description string
	Tags        []string
	Bindings    []binding
	Handlers    []handler
}

func configfile(configFilePath string) config {
	dat, err := os.ReadFile(configFilePath)
	var Config config
	if err == nil {
		jsonError := json.Unmarshal(dat, &Config)
		if jsonError != nil {
			panic("JSON в конфігфайлі попердолено")
		}
		return Config
	} else {
		panic("err")
	}
}
