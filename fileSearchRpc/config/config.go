package config

import (
	"io/ioutil"

	"github.com/smallnest/rpcx/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Etcd      EtcdConfig  `yaml:"etcd"`
	Rpc       RpcConfig   `yaml:"rpc"`
	Redis     RedisConfig `yaml:"redis"`
	SearchDir string      `yaml:"searchdir"`
}

type EtcdConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	BasePath string `yaml:"basepath"`
}

type RpcConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var GConfig *Config

func init() {
	rpcConfig := Config{
		Etcd: EtcdConfig{
			Host:     "localhost",
			Port:     "2379",
			BasePath: "file_search_rpc",
		},
		Rpc: RpcConfig{
			Host: "localhost",
			Port: "9001",
		},
		Redis: RedisConfig{
			Host: "localhost",
			Port: "6379",
		},
		SearchDir: "E:/Work/searchFiles/",
	}
	cfgData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(cfgData, &rpcConfig)
	if err != nil {
		log.Warn(err)
	}
	GConfig = &rpcConfig
}
