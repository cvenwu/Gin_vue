package main

import (
	"gin_vue/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 8:44 上午
 * @Desc: 实现用户注册
 **/


func main() {

	//在项目启动的时候就要读取我们的配置文件
	InitConfig()

	common.InitDB()
	r := gin.Default()

	r = CollectRoute(r)


	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())

}

func InitConfig() {
	//获取当前的工作目录
	workDir, _ := os.Getwd()

	//设置要读取的文件名
	viper.SetConfigName("application")
	//设置要读取的文件类型
	viper.SetConfigType("yml")
	//设置要读取的文件路径
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}



