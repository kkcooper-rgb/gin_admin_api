package core

import (
	"fmt"
	"go_admin_api/config"
	"go_admin_api/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func MysqlInit() error {
	var err error
	var dbconfig = config.Config.Mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbconfig.Username,
		dbconfig.Password,
		dbconfig.Host,
		dbconfig.Post,
		dbconfig.Db,
		dbconfig.Charset)

	// 错误点：这里使用了 := 会创建局部变量Db，覆盖全局变量
	// 改为直接赋值给全局Db
	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{ // 去掉 := 前的Db，使用全局变量
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}

	// 多余的判断，可删除（gorm.Open失败时已返回错误）
	// if Db.Error != nil {
	//     return err
	// }

	sqlDb, err := Db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDb.SetMaxIdleConns(dbconfig.MaxIdle)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(dbconfig.MaxOpen)
	global.Log.Infof("mysql连接成功")
	return nil
}
