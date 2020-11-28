package controller

import (
	"gin_vue/common"
	"gin_vue/model"
	"gin_vue/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 9:34 上午
 * @Desc:
 **/

//实现用户的注册
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

	//TODO:对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 2003,
			"msg":  "加密错误",
		})
		return
	}

	newUser := model.User{
		Username:  username,
		Password:  string(hashedPassword),
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
	//TODO:为什么ID!=0就存在
	//因为我们使用gorm插入第一条记录将会让id为1，后面的依次递增，所以id为0的记录一直都不存在
	if user.ID != 0 {
		return true
	}
	return false
}

//实现用户的登录:只需要手机号以及密码
func Login(ctx *gin.Context) {
	//获取参数
	//username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")

	//实现验证
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

	//判断手机号是否存在
	db := common.GetDB()

	//如果手机号存在，我们就直接返回true
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2004,
			"msg":  "手机号不存在，请稍后再试",
		})
		return
	}

	//如果存在就继续判断密码是否正确
	//如果密码验证成功err就是nil，否则就会返回一个错误
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2005,
			"msg":  "密码不正确！！！",
		})
		return
	}



	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2006,
			"msg": "生成token失败",
		})
		return
	}


	//如果正确就返回token
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1000,
		"msg":  "登录成功！！！",
		"data": gin.H{
			"token": token,
		},
	})
}


//获取用户信息的时候用户一定是经过了认证的，并且我们可以从上下文中获取到用户的信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1000,
		"data": gin.H{
			"user": user,
		},
	})

}