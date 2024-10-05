package initialize

import (
	"starter/mod/global"
	"starter/mod/package/logger"
)

func InitLogger() {
	global.Logger = *logger.NewLogger(global.Config.Logger)
}
