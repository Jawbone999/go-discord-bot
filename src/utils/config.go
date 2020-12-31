package utils

import (
	"encoding/json"
	"io/ioutil"
)

const configFile = "config.json"

// BotConfig contains JSON data from config.json
type BotConfig struct {
	Server     string `json:"twitchServer"`
	Port       int    `json:"twitchPort"`
	Token      string `json:"twitchToken"`
	Channel    string `json:"twitchChannel"`
	Username   string `json:"botUsername"`
	Prefix     string `json:"botPrefix"`
	LogLevel   string `json:"logLevel"`
	LogConsole string `json:"logConsole"`
	LogFile    string `json:"logFile"`
}

// GetConfig returns a BotConfig struct containing data from ./config.json
func GetConfig() (BotConfig, error) {
	config := BotConfig{}

	file, err := ioutil.ReadFile(configFile)

	if err == nil {
		err = json.Unmarshal([]byte(file), &config)
	}

	return config, err
}
