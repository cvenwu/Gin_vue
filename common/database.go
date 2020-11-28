package common

import (
	"fmt"
	"gin_vue/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 9:36 上午
 * @Desc:
 **/

var DB *gorm.DB

func InitDB() {

	driverName := viper.GetString("datasource.driverName")
	user := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	charset := viper.GetString("datasource.charset")

	//最终需要的格式：root:123456@tcp(localhost:3306)/数据库名字?charset=utf8&parseTime=True&loc=Local
	mysqlDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, database, charset)

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
