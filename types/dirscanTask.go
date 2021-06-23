package types

import (
	"yuhScan/logger"
	"yuhScan/utils"
)

type DirScanTask struct {
	Task
	config   Config
	payloads []string
	page404  []byte
	last     []string
	tmp      []string
}

func (t DirScanTask) Config() Config     { return t.config }
func (t DirScanTask) Payloads() []string { return t.payloads }
func (t DirScanTask) Page404() []byte    { return t.page404 }
func (t DirScanTask) Last() []string     { return t.last }
func (t DirScanTask) Tmp() []string      { return t.tmp }

func NewDirScanTask(task Task, cfg Config, payloads []string, last []string, tmp []string) *DirScanTask {
	return &DirScanTask{
		Task:     task,
		config:   cfg,
		payloads: payloads,
		last:     last,
		tmp:      tmp,
	}
}

func (t *DirScanTask) Working() {
	var (
		pathsExist []string
	)

	// Random 404
	randomPath := utils.RandomStr(10)
	if r, ok := utils.HttpRequest(t.config.url+"/"+randomPath, t.config.header, t.config.proxy); ok {
		t.page404 = utils.ReadHttpBody(r)
		if t.config.debug { // DEBUG
			logger.ConsoleLog(logger.DEBUG, "Random 404 Path", randomPath, "Length:", len(t.page404))
		}
	}

	if len(t.last) == 0 {
		for _, payload := range t.payloads {
			if t.CheckTarget(payload) {
				pathsExist = append(pathsExist, payload)
			}
		}
	} else if len(t.last) > 0 {
		for _, path := range t.last {
			if utils.MatchStr("\\.", path) {
				continue
			} else {
				for _, payload := range t.payloads {
					if t.CheckTarget(path + payload) {
						pathsExist = append(pathsExist, path+payload)
					}
				}
			}
		}
	}
	t.ResultChan <- pathsExist
}

/*
 * 页面校验
 */
func (t DirScanTask) CheckTarget(payload string) bool {
	var (
		url           = t.config.url + "/" + string(payload[1:])
		urlWithEncode = t.config.url + "/" + utils.UrlEncode(payload[1:])
	)

	if t.config.debug { // DEBUG
		logger.ConsoleLog(logger.DEBUG, "Checking URL:", url)
	}

	if r, ok := utils.HttpRequest(urlWithEncode, t.config.header, t.config.proxy); ok {

		// 校验状态码-PASS
		if utils.IsIntInArray(r.StatusCode, t.config.statusList) {

			respBytes := utils.ReadHttpBody(r)

			// 404检测-PASS
			if utils.IsSimilarBytes(t.page404, respBytes, utils.Sum(len(payload[1:]), 10)) {
				if t.config.debug { // DEBUG
					logger.ConsoleLog(logger.DEBUG, payload[1:], "Seems 404", "Length:", len(respBytes))
				}
				return false
			}

			// 过滤-PASS
			if len(t.config.filter) > 0 {
				if utils.MatchStr(t.config.filter, string(respBytes)) {
					return false
				}
			}

			// 30X 特殊显示
			if utils.IsIntInArray(r.StatusCode, []int{301, 302, 303, 307}) {
				location, _ := r.Location()
				logger.ConsoleLog(logger.INFO, "Status:", r.StatusCode, "Access:", url, "==>", location)
			} else {
				logger.ConsoleLog(logger.INFO, "Status:", r.StatusCode, "Length:", len(respBytes), "Access:", url)
			}

			return true
		} else {
			return false
		}
	}
	return false
}
