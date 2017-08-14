package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConf() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("database")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	// If no config is found, use the default(s)
	viper.SetDefault("host", "localhost")

	// fmt.Printf("\n%s\n\n", theMessage["host"])
}

func GetDbInfo(db_no string) map[string]string{
	return viper.GetStringMapString(db_no)
}
