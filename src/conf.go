package main

import (
	"github.com/spf13/viper"
)

var v_db = InitConf("database")
var v_aws = InitConf("aws")

func InitConf(file_name string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(file_name)
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return v
}

func GetDbInfo(db_no string) map[string]string{
	return v_db.GetStringMapString(db_no)
}

func GetAwsInfo(key string) string{
	return v_aws.GetString(key)
}
