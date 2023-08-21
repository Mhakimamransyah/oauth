package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/Mhakimamransyah/oauth/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MysqlInstance *Mysql

type Mysql struct {
	dB *gorm.DB
}

func (obj *Mysql) GetDB() (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	return obj.dB.WithContext(ctx), cancel
}

func ConnectMysql() {

	conf := config.Config.MySqlDB

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if e != nil {
		panic(e)
	}

	MysqlInstance = &Mysql{
		dB: db,
	}
}

func DisconnectMysql() {
	dbInstance, _ := MysqlInstance.dB.DB()
	dbInstance.Close()
}
