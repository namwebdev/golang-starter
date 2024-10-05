package initialize

import (
	"fmt"
	"starter/mod/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: \n", err))
	}

	// fmt.Println("Server running on port: ", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Error unmarshalling config: ", err)
	}
}
