package main

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configs struct {
	Port string
	BookApiBaseUrl string
}

func newConfig() (Configs) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	rerr := viper.ReadInConfig()
	LogIfError(rerr)
	
	conf := Configs{
		BookApiBaseUrl:  viper.GetString("BOOKAPIBASEURL"),
		Port:  viper.GetString("PORT"),
	}

	s, _ := json.Marshal(conf)
	log.Info(fmt.Sprintf("loaded envs: %v", string(s)))
	
	logEnvvars()

	return conf
	// merr := viper.Unmarshal(&conf)

	// if LogIfError(merr) {
	// 	return Configs{}, merr
	// } else {
	// 	return conf, nil
	// }
}

func logEnvvars() {
	log.Info(os.Environ())
}