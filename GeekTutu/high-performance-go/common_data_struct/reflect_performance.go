package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Name    string `json:"server-name"` // CONFIG_SERVER_NAME
	IP      string `json:"server-ip"`   // CONFIG_SERVER_IP
	URL     string `json:"server-url"`  // CONFIG_SERVER_URL
	Timeout string `json:"timeout"`     // CONFIG_TIMEOUT
}

func readConfig() *Config {
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

//func main() {
//	err := os.Setenv("CONFIG_SERVER_NAME", "global_server")
//	if err != nil {
//		return
//	}
//	if err = os.Setenv("CONFIG_SERVER_IP", "10.0.0.1"); err != nil {
//		return
//	}
//	if err = os.Setenv("CONFIG_SERVER_URL", "geektutu.com"); err != nil {
//		return
//	}
//	if err = os.Setenv("CONFIG_SERVER_TIMEOUT", "1"); err != nil {
//		return
//	}
//	c := readConfig()
//	fmt.Printf("%+v", c)
//}
