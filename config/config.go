package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s.io/klog/v2"
)

var keyMap map[KeyName]string

type Config struct {
	Server Server	`json:"server"`
}

type Server struct {
	Name string `json:"name"`
	Host string	`json:"host"`
	Port string `json:"port"`
}

func init() {
	var config Config
	f, err := ioutil.ReadFile("./.gin-client-go.yaml")
	if err != nil {
		klog.Fatal(err)
		return
	}
	if err := yaml.Unmarshal(f, &config); err != nil {
		klog.Fatal(err)
		return
	}
	keyMap = make(map[KeyName]string)
	keyMap[ServerName] = config.Server.Name
	keyMap[ServerHost] = config.Server.Host
	keyMap[ServerPort] = config.Server.Port
	fmt.Printf("keyConfig: %v", keyMap)
}

func GetString(keyName KeyName) string {
	return keyMap[keyName]
}
