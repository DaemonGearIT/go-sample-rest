package config

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	CONFIG_FILE_NAME = "./config.yml"
)

type Config struct {
	Database struct {
		dialect string
		host    string
		port    int
		dbname  string
		user    string
		passwd  string
	}
}

func (c *Config) Configure() {
	dat, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		color.Red(fmt.Sprintf("Config Errro : %v", err.Error()))
	}
}
