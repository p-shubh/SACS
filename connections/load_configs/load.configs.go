package loadConfig

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfigs(loadFor string) {

	//creating switch statement
	switch loadFor {
	case "env":
		viper.SetConfigName(".env")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logrus.Info("Failed to load : " + loadFor)
				// Config file not found; ignore error if desired
			} else {
				// Config file was found but another error was produced
				logrus.Info("File has an issue while loading")
			}
		}
	// case "prod":
	// 	viper.SetConfigName("prod")
	// 	viper.AddConfigPath("./config")
	default:
		log.Fatal("Invalid loadFor parameter")
	}

}
