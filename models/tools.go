package models

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// MD5加密
func MD5maker(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// 表示把string转换成int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// 表示把int转换成string
func String(n int) string {
	str := strconv.Itoa(n)
	return str
}
