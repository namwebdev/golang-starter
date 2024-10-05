package initialize

import (
	"fmt"
	"time"

	"starter/mod/global"
	"starter/mod/internal/persistent_objects"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	m := global.Config.MySql
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		global.Logger.Error("InitMysql failed, err: %v", zap.Error(err))
	}

	global.Logger.Info("InitMysql success")
	global.Mdb = db
	SetPool()
	migrateTables()
}

func SetPool() {
	sqlDb, err := global.Mdb.DB()
	m := global.Config.MySql
	if err != nil {
		fmt.Printf("SetPool failed, err: %v", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConnections))
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnectionMaxLiftTime))
	sqlDb.SetMaxOpenConns(m.MaxOpenConnections)
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(&persistent_objects.User{}, &persistent_objects.Role{})
	if err != nil {
		fmt.Printf("migrateTables failed, err: %v", err)
	}
}
