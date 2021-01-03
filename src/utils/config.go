package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
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
	LogConsole bool   `json:"logConsole"`
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

// ValidateConfig validates configuration settings
func ValidateConfig(config *BotConfig) error {
	if config.Channel == "" {
		return errors.New("channel field is empty")
	}

	if config.Prefix == "" {
		return errors.New("bot command prefix field is empty")
	}

	if config.Token == "" {
		return errors.New("oauth token field is empty")
	}

	if !strings.HasPrefix(config.Token, "oauth:") {
		config.Token = "oauth:" + config.Token
	}

	if config.Username == "" {
		return errors.New("username field is empty")
	}

	return nil
}
