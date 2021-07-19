package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/hunyaio/yuhScan/logger"
)

// 禁用跳转
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func NewHttpClient(proxyUrl string) *http.Client {
	if len(proxyUrl) > 0 {
		proxy, _ := url.Parse(proxyUrl)
		return &http.Client{
			CheckRedirect: redirectPolicyFunc,
			Timeout:       time.Second * 10,
			Transport: &http.Transport{
				Proxy:           http.ProxyURL(proxy),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	} else {
		return &http.Client{
			CheckRedirect: redirectPolicyFunc,
			Timeout:       time.Second * 10,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}
}

func ReadHttpBody(r *http.Response) (resp []byte) {
	if r != nil {
		defer r.Body.Close()
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.ConsoleLog(logger.WARN, err)
		}
		return resp
	}
	return
}

func HttpRequest(u string, header map[string]string, proxyUrl string) (*http.Response, bool) {
	// Init
	client := NewHttpClient(proxyUrl)
	request, _ := http.NewRequest("GET", u, nil)
	// Headers
	for k, v := range header {
		request.Header.Add(k, v)
	}
	// Request
	r, err := client.Do(request)
	if err != nil {
		logger.ConsoleLog(logger.WARN, err)
		return r, false
	}

	return r, r != nil
}
