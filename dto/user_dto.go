package dto

import "gin_vue/model"

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 12:34 下午
 * @Desc: 用来表示我们希望给前端返回的信息，对于一些不需要的信息我们没有必要进行返回
 **/

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json"telephone"`
}

//定义一个转换的函数
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Username,
		Telephone: user.Telephone,
	}
}
