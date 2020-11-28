package common

import (
	"gin_vue/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 10:51 上午
 * @Desc: JWT鉴权
 **/

//定义一个jwt加密的秘钥
var jwtKey = []byte("a_secret_crect")

//定义token的claim
type Claims struct {
	//定义一个标准的jwt的claim
	jwt.StandardClaims
	UserId uint
}

//发放token给用户
func ReleaseToken(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//设置过期时间
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			//设置是谁签发的
			Issuer: "sivan.tech",
			//设置啥时候签发的,
			IssuedAt: time.Now().Unix(),
			//设置主题
			Subject: "user token",
		},
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

//从tokenString中解析出我们的claim并返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
