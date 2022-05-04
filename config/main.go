/*
@Description：测试读取yaml配置文件
@Author : gilbert
@Date : 2022/5/4 7:15 PM
*/

package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type User struct {
	Name string `yaml:"name"`
	Age  string `yaml:"age"`
}

func (u *User) getConf() *User {
	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, u)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return u
}

func main() {
	var user User
	user.getConf()
	fmt.Printf("name:%s age:%s\n", user.Name, user.Age)
}
