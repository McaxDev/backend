package utils

import (
	"os"
	"reflect"
)

func LoadConfig(config any) {

	if config == nil {
		return
	}

	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tags := typ.Field(i).Tag

		envKey := tags.Get("env")
		if envKey == "" {
			continue
		}

		env, exists := os.LookupEnv(envKey)
		if exists {
			field.SetString(env)
		} else {
			field.SetString(tags.Get("def"))
		}
	}
}
