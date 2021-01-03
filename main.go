package main

import "github.com/Jawbone999/go-twitch-bot/src/utils"

func main() {
	configs, err := utils.GetConfig()
	if err != nil {
		utils.HandleError("Failed to read config file", err)
		return
	}
	utils.SetupLogger()

}
