package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type mysql struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type log struct {
	LogFilePath string `mapstructure:"logFilePath"`
	LogFileName string `mapstructure:"logFileName"`
}

type jwt struct {
	Secret string `mapstructure:"secret"`
}

type conf struct {
	Mysql mysql `mapstructure:"mysql"`
	Log   log   `mapstructure:"log"`
	Jwt   jwt   `mapstructure:"jwt"`
}

var Conf conf

func Setup() {
	viper.SetConfigFile("config.toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Unable to decode into struct: %s \n", err))
	}
}
