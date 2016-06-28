package config

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	CONFIG_FILE_NAME = "conf.yml"
)

type (
	Configuration struct {
		DATABASE struct {
			DIALECT  string `yaml:"dialect"`
			HOST     string `yaml:"host"`
			PORT     int    `yaml:"port"`
			DBNAME   string `yaml:"dbname"`
			USER     string `yaml:"user"`
			PASSWORD string `yaml:"passwd"`
		} `yaml:"database"`
	}
)

func Configure() Configuration {
	c := Configuration{}

	data, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		color.Red(fmt.Sprintf("Config Error 500 : %v\n", err.Error()))
	}
	color.Green(fmt.Sprintf("Data Loaded: \n%v\n", string(data)))

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		color.Red(fmt.Sprintf("Config Error 501 : %v\n", err.Error()))
	}
	color.Green(fmt.Sprintf("Config Loaded: %v\n", c))

	return c
}
