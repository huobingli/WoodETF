package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var conf_file_path string = "D:\\gitProject\\WoodETF\\woodServer\\config\\setting.yaml"

type redis_conf struct {
	redis struct {
		host     string `yaml:"host"`
		port     int    `yaml:"port"`
		user     string `yaml:"user"`
		password string `yaml:"password"`
	}
}

func load_config() {

	conf := new(redis_conf)

	yamlFile, err := ioutil.ReadFile(conf_file_path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(yamlFile)

	err = yaml.Unmarshal(yamlFile, &conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
}

func main() {
	load_config()
}
