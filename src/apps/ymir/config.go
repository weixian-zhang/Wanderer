package main

import (
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
	LogIfError(rerr)
	
	conf := Configs{
		BookApiBaseUrl:  viper.GetString("BookApiBaseUrl"),
		Port:  viper.GetString("Port"),
	}
	merr := viper.Unmarshal(&conf)

	if LogIfError(merr) {
		return Configs{}, merr
	} else {
		return conf, nil
	}
}