package main

import (
	"fmt"
	"errors"
	"os"
	"strings"
	"reflect"
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

func sliceToByteArray(t interface{}) ([]byte, error) {
	reflectedType :=  reflect.TypeOf(t)

	if reflectedType.Kind() != reflect.Slice {
		return nil, errors.New("argument is not a slice")
	}

	// for ele := range reflectedType.Elem() {

	// }

	return nil, nil
}