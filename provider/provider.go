package provider

import (
	"FitnessSpace/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	db := ConnectDatabase(
		"root",
		"123456",
		"192.168.0.104:3309",
		"fitness_space",
	)
	err := db.AutoMigrate(&model.Coach{}, &model.Course{}, &model.Member{})
	if err != nil {
		println("创建数据表出错")
		return nil
	}
	return db
}

func ConnectDatabase(user string, pass string, dbHost string, dbName string) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, dbHost, dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败 : %s", err.Error()))
	}
	println("数据库连接成功")
	return db
}
