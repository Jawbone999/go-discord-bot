package main

import (
	"github.com/Jawbone999/go-twitch-bot/src/utils"
)

func main() {
	config, err := utils.GetConfig()
	if err != nil {
		utils.HandleError("Failed to read config file", err)
		return
	}

	err = utils.validateConfig(config)
	if err != nil {
		utils.HandleError("Failed to validate config file", err)
	}

	utils.SetupLogger(config)

}
