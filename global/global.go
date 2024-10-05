package global

import (
	"gorm.io/gorm"

	"starter/mod/package/logger"
	"starter/mod/package/setting"
)

var (
	Config setting.Config
	Logger logger.LoggerZap
	Mdb    *gorm.DB
)
