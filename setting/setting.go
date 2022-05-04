/*
@Description：读取配置文件数据
@Author : gilbert
@Date : 2022/5/2 11:16 PM
*/

package setting

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// MySql 数据库配置
type MySql struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// AppConfig 应用配置
type AppConfig struct {
	Env     string `yaml:"env"`
	Port    string `yaml:"port"`
	Release bool   `yaml:"release"`
	*MySql  `yaml:"mysql"`
}

var app AppConfig

func InitConfig() AppConfig {
	yamlFile, err := ioutil.ReadFile("/Users/gilbert/go/src/inventory/config/app.yaml")
	if err != nil {
		log.Println("read file failed, err:", err)
	}

	err = yaml.Unmarshal(yamlFile, app)
	if err != nil {
		log.Println("unmarshal failed, err:", err)
	}
	return app
}
