package controller

import (
	"gin_vue/common"
	"gin_vue/model"
	"gin_vue/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 9:34 上午
 * @Desc:
 **/

func Register(ctx *gin.Context) {
	//1. 获取参数
	telephone := ctx.PostForm("telephone")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//2. 实现数据认证

	if len(telephone) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "手机号不可以为空",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "密码不可以少于6位",
		})
		return
	}

	//如果用户的名称没有传，我们生成一个10位的随机字符串
	if len(username) == 0 {
		username = util.RandString(10)
	}

	//3. 查询用户是否存在-->判断手机号是否存在
	db := common.GetDB()
	if isTelephoneExists(db, telephone) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "手机号已经存在！！！",
		})
		return
	}

	//3. 进行用户注册
	//说明用户不存在，我们就新建用户
	newUser := model.User{
		Username:  username,
		Password:  password,
		Telephone: telephone,
	}
	//这里要传入地址
	db.Create(&newUser)

	log.Println(username, password, telephone)

	//4. 返回信息

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1000,
		"msg":  "注册成功！！！",
	})
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)

	//说明用户存在
	if user.ID != 0 {
		return true
	}

	return false

}
