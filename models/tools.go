package models

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
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

// 获取int类型的Unix时间戳
func GetUnix() int {
	timeUnix := time.Now().Unix()
	n := int(timeUnix)
	return n
}

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
