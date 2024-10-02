package initialize

import (
	"fmt"
	"starter/mod/global"
)

func Run() {
	LoadConfig()
	fmt.Println("---- MySql configuration: ", global.Config.MySql)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":5000")

	fmt.Println("Server is running on http://localhost:5000")
}
