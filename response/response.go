package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 12:37 下午
 * @Desc: 封装响应返回
 **/

//一般我们希望将我们返回的格式进行统一
//{
//	code: 20001,
//	data: xxx,
//	msg: xx,
//}

//httpStatus是http的状态码
//code 是我们自己定义的业务code
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
