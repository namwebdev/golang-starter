package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"starter/mod/package/logger"
	"starter/mod/package/setting"
)

var (
	Config setting.Config
	Logger logger.LoggerZap
	Mdb    *gorm.DB
	Redis  *redis.Client
)
