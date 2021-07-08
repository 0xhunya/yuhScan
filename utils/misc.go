package utils

import (
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/hunyaio/yuhScan/logger"
)

func CheckURL(url string) bool {
	// Init
	urlRules := map[string]string{
		"domain": `^(https?://)?([A-Za-z0-9\-]+\.)+[A-Za-z\-]+(:[0-9]{1,5})?(\/.*)?$`,
		"ip":     `^(https?://)?([0-9]{1,3}\.){3}[0-9]{1,3}(:[0-9]{1,5})?(\/.*)?$`,
	}
	// Check
	if MatchStr(urlRules["domain"], url) || MatchStr(urlRules["ip"], url) {
		return true
	} else {
		return false
	}
}

func CheckProxy(proxyUrl string) {
	_, err := url.Parse(proxyUrl)
	if err != nil {
		logger.ConsoleLog(logger.ERROR, "Error Parsing ProxyURL")
	}
}

func UrlParse(s string) (url string) {
	if CheckURL(s) {
		if string(s[len(s)-1]) == "/" {
			s = s[:len(s)-1]
		}
		if s[0:4] != "http" {
			url = "http://" + s
		} else {
			url = s
		}
	} else {
		logger.ConsoleLog(logger.ERROR, "URL ERROR")
	}
	return
}

func StatusParse(s string) (statuslist []int) {
	for _, v := range strings.Split(s, ",") {
		statuslist = append(statuslist, StringToInt(v))
	}
	return
}

func RandomStr(length int) string {
	var (
		r      = rand.New(rand.NewSource(time.Now().UnixNano()))
		letter = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		res    = make([]byte, length)
	)

	for i := 0; i < length; i++ {
		res[i] = letter[r.Intn(len(letter))]
	}
	return string(res)
}

func IsSimilarBytes(arr1 []byte, arr2 []byte, diff int) bool {

	// 长度校验
	if len(arr1) == len(arr2) || Sum(len(arr1), len(arr2)) == diff {
		// 内容校验
		k := MinInt2(len(arr1), len(arr2)) / 5
		return string(arr1[:k]) == string(arr2[:k])
	}
	return false
}
