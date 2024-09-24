package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Client Client `yaml:"client"`
	Server Server `yaml:"server"`
	Other  Other  `yaml:"other"`
}

type Client struct {
	LocalServerAddr   string `yaml:"LocalServerAddr"`
	RemoteServerAddr  string `yaml:"RemoteServerAddr"`
	RemoteControlAddr string `yaml:"RemoteControlAddr"`
}

type Server struct {
	ControlAddr string `yaml:"ControlAddr"`
	TunnelAddr  string `yaml:"TunnelAddr"`
	VisitAddr   string `yaml:"VisitAddr"`
}

type Other struct {
	Checkip string `yaml:"Checkip"`
	Mode    string `yaml:"Mode"`
}

// 定义一个全局变量来存储配置
var config Config

func LoadConfig() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("无法读取 YAML 文件: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("无法解析 YAML 文件: %v", err)
	}
}

// 提供一个方法来获取配置
func GetConfig() Config {
	return config
}