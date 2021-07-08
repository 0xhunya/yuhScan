package main

import (
	"flag"
	"sync"

	"github.com/hunyaio/yuhScan/cmd"
	"github.com/hunyaio/yuhScan/core"
	"github.com/hunyaio/yuhScan/logger"
	"github.com/hunyaio/yuhScan/types"
	"github.com/hunyaio/yuhScan/utils"
)

var (
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	// Check
	if cmd.Help() || (len(cmd.Url()) == 0 && len(cmd.Ufile()) == 0) || len(cmd.File()) == 0 {
		flag.Usage()
		return
	}
	if len(cmd.Url())^len(cmd.Ufile()) == 0 {
		logger.ConsoleLog(logger.ERROR, "URL or URLFile")
	}

	if len(cmd.Url()) > 0 {
		cfg := types.NewConfigFromParams(cmd.Url(), cmd.File(), cmd.Thread(), cmd.Status(), cmd.Header(), cmd.Proxy(), cmd.Filter(), cmd.Output(), cmd.Recursive(), cmd.Debug(), &mutex)
		core.DirScan(cfg)
	} else if len(cmd.Ufile()) > 0 {
		urls := utils.LoadDictOne(cmd.Ufile())
		for n, url := range urls {
			cfg := types.NewConfigFromParams(url, cmd.File(), cmd.Thread(), cmd.Status(), cmd.Header(), cmd.Proxy(), cmd.Filter(), cmd.Output(), cmd.Recursive(), cmd.Debug(), &mutex)
			wg.Add(1)
			go core.DirScan(cfg, &wg)
			if cmd.Group() > 0 && (n+1)%cmd.Group() == 0 {
				wg.Wait()
			}
		}
		wg.Wait()
	} else {
		logger.ConsoleLog(logger.ERROR, "Unknown Error")
	}
}
