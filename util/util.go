package util

import (
	"math/rand"
	"time"
)

/**
 * @Author: yirufeng
 * @Date: 2020/11/28 9:36 上午
 * @Desc:
 **/

//生成一个n位的随机字符串
func RandString(n int) string {
	letters := []byte("qwekqholfsdlgnruikhiurjojpgperwgklfdnkjgbdfnjkgdfgd")
	result := make([]byte, n)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		result[i] = letters[rand.Intn(n)]
	}
	return string(result)
}
