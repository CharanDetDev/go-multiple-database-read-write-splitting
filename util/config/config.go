package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/converse"
	"github.com/spf13/viper"
)

var Env Config

type Config struct {
	API_PORT string `mapstructure:"API_PORT"`

	MYSQL_DRIVER_NAME string `mapstructure:"MYSQL_DRIVER_NAME"`
	MYSQL_HOSTNAME_1  string `mapstructure:"MYSQL_HOSTNAME_1"`
	MYSQL_HOSTNAME_2  string `mapstructure:"MYSQL_HOSTNAME_2"`
	MYSQL_USERNAME    string `mapstructure:"MYSQL_USERNAME"`
	MYSQL_PASSWORD    string `mapstructure:"MYSQL_PASSWORD"`
	MYSQL_DB_NAME     string `mapstructure:"MYSQL_DB_NAME"`

	REDIS_ADDRESS  string `mapstructure:"REDIS_ADDRESS"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_DB_NUM   string `mapstructure:"REDIS_DB_NUM"`
	SKIP_UNIT_TEST string `mapstructure:"SKIP_UNIT_TEST"`
}

func ConfigInit() bool {

	viper.AddConfigPath(".")
	viper.AddConfigPath("./util/config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("***** Initail Config With Viper ERROR :: viper.ReadInConfig() ***** | ", err)
		return false
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		fmt.Println("***** Initail Config With Viper ERROR :: viper.Unmarshal(&Env) ***** | ", err)
		return false
	}

	fn := reflect.ValueOf(&Env).Elem()
	for i := 0; i < fn.NumField(); i++ {
		value := converse.ParseToString(fn.Field(i).Interface())
		reflect.ValueOf(&Env).Elem().FieldByName(fn.Type().Field(i).Name).SetString(value)
	}

	return true

}

func ConfigInitForTest() {
	viper.AddConfigPath(".")
	// viper.AddConfigPath("../config") // for Database unit test
	viper.AddConfigPath("../util/config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("***** Initail Config With Viper ERROR :: viper.ReadInConfig() ***** | ", err)
		panic(err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		fmt.Println("***** Initail Config With Viper ERROR :: viper.Unmarshal(&Env) ***** | ", err)
		panic(err)
	}

	fn := reflect.ValueOf(&Env).Elem()
	for i := 0; i < fn.NumField(); i++ {
		value := converse.ParseToString(fn.Field(i).Interface())
		reflect.ValueOf(&Env).Elem().FieldByName(fn.Type().Field(i).Name).SetString(value)
	}

}
