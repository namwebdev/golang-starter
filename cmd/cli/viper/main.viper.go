package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

type Server struct {
	Port int
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: \n", err))
	}

	// fmt.Println("Server running on port: ", viper.GetInt("server.port"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Error unmarshalling config: ", err)
	}

	fmt.Println("Databases: ", config.Databases)

	for _, db := range config.Databases {
		fmt.Println("Database: ", db)
	}
}
