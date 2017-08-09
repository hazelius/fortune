package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DBInfo struct {
	Host     string `yaml:"host"`
	UserName string `yaml:"username"`
	Password string `yaml:"Password"`
}

func GetConf() (map[string]DBInfo, error) {
	buf, err := ioutil.ReadFile("database.yml")
	if err != nil {
		panic(err)
	}

	dbconf := make(map[string]DBInfo)
	err = yaml.Unmarshal(buf, dbconf)

	return dbconf, err
}
