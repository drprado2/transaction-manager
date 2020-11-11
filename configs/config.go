package configs

import (
	"github.com/tkanos/gonfig"
	"os"
	"strconv"
)

type Configuration struct {
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	HttpServerPort   string
	MigrationsPath   string
}

var loadedConfiguration *Configuration

func GetConfig() *Configuration {
	if loadedConfiguration != nil {
		return loadedConfiguration
	}

	configuration := Configuration{}
	fileName := "configs/config.json"
	gonfig.GetConf(fileName, &configuration)
	overrideWithEnvVars(&configuration)
	loadedConfiguration = &configuration
	return loadedConfiguration
}

func overrideWithEnvVars(config *Configuration) {
	if value := os.Getenv("DatabaseHost"); len(value) > 0 {
		config.DatabaseHost = value
	}
	if value := os.Getenv("DatabasePort"); len(value) > 0 {
		iValue, _ := strconv.Atoi(value)
		config.DatabasePort = iValue
	}
	if value := os.Getenv("DatabaseUser"); len(value) > 0 {
		config.DatabaseUser = value
	}
	if value := os.Getenv("DatabasePassword"); len(value) > 0 {
		config.DatabasePassword = value
	}
	if value := os.Getenv("DatabaseName"); len(value) > 0 {
		config.DatabaseName = value
	}
	if value := os.Getenv("HttpServerPort"); len(value) > 0 {
		config.HttpServerPort = value
	}
	if value := os.Getenv("MigrationsPath"); len(value) > 0 {
		config.MigrationsPath = value
	}
}
