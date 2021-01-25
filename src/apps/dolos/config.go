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
	MongoConnstring  string
	AzDiskMountFilePath string
}

func newConfig() (Configs) {

	// log.Info(fmt.Sprintf("get envvar by os Port: %v", os.Getenv("Port")))
	// log.Info(fmt.Sprintf("get envvar by os AzDiskMountFilePath: %v", os.Getenv("AzDiskMountFilePath"))) 
	
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	rerr := viper.ReadInConfig()
	LogIfError(rerr)

	viper.AutomaticEnv()
	
	conf := Configs{
		AzDiskMountFilePath:  viper.GetString("AZDISKMOUNTFILEPATH"),
		Port:  viper.GetString("PORT"),
		MongoConnstring: viper.GetString("MongoConnstring"),
	}
	// merr := viper.Unmarshal(&conf)

	logEnvVar(conf)

	return conf

	// if merr != nil {
	// 	log.Error(merr)
	// 	return Configs{}, merr
	// } else {
	// 	return conf, nil
	// }
}

func logEnvVar(c Configs) {
	s, _ := json.Marshal(c)
	log.Info(fmt.Sprintf("Loaded envvars: %v", string(s)))

	log.Info(os.Environ())
}