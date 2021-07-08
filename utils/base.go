package utils

import (
	"net/url"
	"regexp"
	"strconv"

	"github.com/hunyaio/yuhScan/logger"
)

func Sum(a int, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
func MinInt2(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func MatchStr(substr string, str string) bool {
	ok, _ := regexp.MatchString(substr, str)
	if ok {
		return true
	} else {
		return false
	}
}
func IsIntInArray(n int, arr []int) bool {
	for _, i := range arr {
		if n == i {
			return true
		}
	}
	return false
}
func StringToInt(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		logger.ConsoleLog(logger.ERROR, "Error Converting String to Int")
	}
	return res
}
func StringArrayCut(array []string, num int) [][]string {
	if num > len(array) {
		logger.ConsoleLog(logger.WARN, "Array Cut Limit:", num, "=>", len(array))
		num = len(array)
	}
	res := make([][]string, num)

	for k := range array {
		p := (k + num) % num
		res[p] = append(res[p], array[k])
	}
	return res
}
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}
