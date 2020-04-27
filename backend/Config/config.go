package Config

import (
	"fmt"
	"github.com/spf13/viper"
)

type databaseConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type logConfig struct {
	LogFilePath string `mapstructure:"logFilePath"`
	LogFileName string `mapstructure:"logFileName"`
}

type config struct {
	Db  databaseConfig `mapstructure:"database"`
	Log logConfig      `mapstructure:"log"`
}

var Conf config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Unable to decode into struct: %s \n", err))
	}
}
