package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var CONF_FILE string = "/home/aaron/git/go-txt2img-discord-bot/config.json"
var (
	Token     string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile(CONF_FILE)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
