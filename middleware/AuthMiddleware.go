package middleware

import (
	"gin_vue/common"
	"gin_vue/model"
	"gin_vue/response"
	"github.com/gin-gonic/gin"
	"strings"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 11:53 上午
 * @Desc:
 **/

//gin的中间件就是一个函数，返回值是一个gin.HandlerFunc
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//首先我们在请求的投不中获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {

			response.Fail(ctx, nil, "权限不足")
			//将这次请求抛弃掉
			ctx.Abort()
			return
		}

		//因为前面7位被"Bearer "占据，因此我们解析后面的
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		//如果解析失败或者解析后的token无效，我们就返回权限不足
		if err != nil || !token.Valid {
			response.Fail(ctx, nil, "权限不足")

			//将这次请求抛弃掉
			ctx.Abort()
			return
		}

		//说明token通过验证
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		//TODO:为什么这个可以写成如下代码而不写成这种形式：	db.Where("telephone = ?", telephone).First(&user)
		DB.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			response.Fail(ctx, nil, "权限不足")

			//将这次请求抛弃掉
			ctx.Abort()
			return
		}

		//用户存在，将user的信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
