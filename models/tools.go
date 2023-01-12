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

// 把string转换成Float64
func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
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

// map[int]int到[]string
func Map2Slice(m map[int]int) []string {
	s := make([]string, 0, len(m))
	for _, v := range m {
		vString := strconv.Itoa(v)
		s = append(s, vString)
	}
	return s
}

// 判断值是否在slice内
func InSliceOK(fruits []string, n string) bool {
	fm := make(map[string]int)
	for i, v := range fruits {
		fm[v] = i
	}
	_, ok := fm[n]
	return ok
}

// map[string]string到[]string
func MapString2Slice(m map[string]string) []string {
	s := make([]string, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// 获取纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}
