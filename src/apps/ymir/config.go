package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configs struct {
	Port string
	BookApiBaseUrl string
}

func newConfig() (Configs, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	rerr := viper.ReadInConfig()
	if rerr != nil {
		log.Error(rerr)
		return Configs{}, rerr
	}
	
	conf := Configs{}
	merr := viper.Unmarshal(&conf)

	if LogIfError(merr) {
		return Configs{}, merr
	} else {
		return conf, nil
	}
}