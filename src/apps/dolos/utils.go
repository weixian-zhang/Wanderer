package main

import (
	"fmt"
	"errors"
	"os"
	"strings"
)

type Utils struct {}

func (u Utils) getEnvarVal(key string) (string, error) {
	var val string = ""
	for _, env := range os.Environ() {

		kv := strings.Split(env, "=")

		if kv[0] == fileEnvVarKey {
			val = kv[1]
		}
	}

	if val == "" {
		return val, errors.New(fmt.Sprintf("Environment varilable not found for %v", key))
	} else {
		return val, nil
	}
}