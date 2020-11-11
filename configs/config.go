package configs

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
)

type Configuration struct {
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	HttpServerPort   string
}

var loadedConfiguration *Configuration

func GetConfig() *Configuration {
	if loadedConfiguration != nil {
		return loadedConfiguration
	}

	appEnv := os.Getenv("APP_ENVIRONMENT")
	configuration := Configuration{}
	env := "development"
	if len(appEnv) > 0 {
		env = appEnv
	}
	fileName := fmt.Sprintf("./configs/config.%s.json", env)
	if err := gonfig.GetConf(fileName, &configuration); err != nil {
		panic("Problem loading configurations")
	}
	loadedConfiguration = &configuration
	return loadedConfiguration
}
