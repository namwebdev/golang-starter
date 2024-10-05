package initialize

import (
	"fmt"
	"starter/mod/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("---- MySql configuration: ", global.Config.MySql)
	InitLogger()
	global.Logger.Info("Logger is initialized", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":5000")

	fmt.Println("Server is running on http://localhost:5000")
}
