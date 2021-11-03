package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var conf_file_path string = "D:\\gitProject\\WoodETF\\woodServer\\config\\setting.yaml"

type redis_conf1 struct {
	redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

type Record struct {
	// redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	// }
}

type redis_conf struct {
	Record Record `yaml:"redis"`
}

func load_config() {

	//conf := new(map[interface{}]interface{})
	conf := new(redis_conf)

	yamlFile, err := ioutil.ReadFile(conf_file_path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(yamlFile)

	err = yaml.Unmarshal(yamlFile, conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
}

func main() {
	load_config()
}
