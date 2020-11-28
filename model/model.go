package model

import "github.com/jinzhu/gorm"

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 9:32 上午
 * @Desc:
 **/


type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
}