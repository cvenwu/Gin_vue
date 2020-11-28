package common

import (
	"fmt"
	"gin_vue/model"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 9:36 上午
 * @Desc:
 **/

var DB *gorm.DB

func InitDB() {

	driverName := "mysql"
	user := "root"
	password := "1018222wxw"
	database := "gin_test"
	host := "localhost"
	port := "3306"

	//最终需要的格式：root:123456@tcp(localhost:3306)/数据库名字?charset=utf8&parseTime=True&loc=Local
	mysqlDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	var err error
	DB, err = gorm.Open(driverName, mysqlDSN)
	if err != nil {
		panic(err)
	}


	//让gorm自动创建数据表
	DB.AutoMigrate(&model.User{})
}


func GetDB() *gorm.DB {
	return DB
}