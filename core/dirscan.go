package core

import (
	"strings"
	"sync"

	"yuhScan/logger"
	"yuhScan/types"
	"yuhScan/utils"
)

func makeTask(pool *types.Pool, cfg types.Config, payloadsCut [][]string, last []string) {
	for _, payloads := range payloadsCut {
		w := types.NewDirScanTask(types.Task{MsgChan: pool.MsgChan, ResultChan: pool.ResultChan}, cfg, payloads, last, nil)
		pool.TaskChan <- w
	}
}

func DirScan(cfg types.Config, wg ...*sync.WaitGroup) {
	var (
		pathExist   []string
		last        []string
		tmp         []string
		pool        = types.NewPool()
		payloadsCut = utils.StringArrayCut(utils.LoadPayloadsFromFile(cfg.File()), cfg.Thread())
	)

	go makeTask(pool, cfg, payloadsCut, last)
	pool.Run()

	for _, i := range pool.Result() {
		last = append(last, i.([]string)...)
	}
	if cfg.Debug() { // DEBUG
		logger.ConsoleLog(logger.DEBUG, "FirstRun Finished", "last:", last)
	}

	// Recursive
	if cfg.Recursive() && len(last) > 0 {
		pathExist = append(pathExist, last...)
	RLOOP:
		go makeTask(pool, cfg, payloadsCut, last)
		pool.Run()
		for _, i := range pool.Result() {
			tmp = append(tmp, i.([]string)...)
		}
		if cfg.Debug() { // DEBUG
			logger.ConsoleLog(logger.DEBUG, "RLOOP", "pathExist:", pathExist, "last:", last, "tmp:", tmp)
		}
		if len(tmp) > 0 {
			pathExist = append(pathExist, tmp...)
			last = tmp
			tmp = []string{}
			goto RLOOP
		}
	} else {
		pathExist = append(pathExist, last...)
	}

	// Output
	result := cfg.Url() + strings.Join(pathExist, "\n"+cfg.Url()) + "\n"
	if len(pathExist) > 0 && len(cfg.OutputFile()) > 0 {
		cfg.Mutex().Lock()
		utils.OutputToFile(cfg.OutputFile(), result)
		cfg.Mutex().Unlock()
	}
	logger.ConsoleLog(logger.RESULT, result)
	if len(wg) == 1 {
		wg[0].Done()
	}
}
