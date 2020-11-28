package main

import (
	"gin_vue/controller"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 9:39 上午
 * @Desc: 负责收集路由
 **/

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}
