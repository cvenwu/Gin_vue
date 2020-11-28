package main

import (
	"gin_vue/common"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 8:44 上午
 * @Desc: 实现用户注册
 **/


func main() {
	common.InitDB()
	r := gin.Default()

	r = CollectRoute(r)

	r.Run()
}



