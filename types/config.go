package types

import (
	"encoding/json"
	"sync"

	"github.com/hunyaio/yuhScan/logger"
	"github.com/hunyaio/yuhScan/utils"
)

type Config struct {
	url        string
	file       string
	thread     int
	statusList []int
	header     map[string]string
	proxy      string
	filter     string
	outputFile string
	recursive  bool
	debug      bool
	mutex      *sync.Mutex
}

func (cfg Config) Url() string               { return cfg.url }
func (cfg Config) File() string              { return cfg.file }
func (cfg Config) Thread() int               { return cfg.thread }
func (cfg Config) StatusList() []int         { return cfg.statusList }
func (cfg Config) Header() map[string]string { return cfg.header }
func (cfg Config) Proxy() string             { return cfg.proxy }
func (cfg Config) Filter() string            { return cfg.filter }
func (cfg Config) OutputFile() string        { return cfg.outputFile }
func (cfg Config) Recursive() bool           { return cfg.recursive }
func (cfg Config) Debug() bool               { return cfg.debug }
func (cfg Config) Mutex() *sync.Mutex        { return cfg.mutex }

/*
 * Constructor
 */
func NewConfig(url string, file string, thread int, statuslist []int, header map[string]string, proxy string, filter string, output string, recursive bool, debug bool, mutex *sync.Mutex) Config {
	return Config{
		url:        url,
		file:       file,
		thread:     thread,
		statusList: statuslist,
		header:     header,
		proxy:      proxy,
		filter:     filter,
		outputFile: output,
		recursive:  recursive,
		debug:      debug,
		mutex:      mutex,
	}
}

/*
 * Constuctor With Check From Params
 */
func NewConfigFromParams(url string, file string, thread int, status string, header string, proxy string, filter string, output string, recursive bool, debug bool, mutex *sync.Mutex) Config {

	// Check URL
	url = utils.UrlParse(url)
	// Check Status
	statuslist := utils.StatusParse(status)
	// Check Proxy
	if len(proxy) > 0 {
		utils.CheckProxy(proxy)
	}
	// Check Header
	headerData := make(map[string]string)
	if len(header) > 0 {
		if utils.FileExits(header) {
			if err := json.Unmarshal(utils.FileRead(header), &headerData); err != nil {
				logger.ConsoleLog(logger.ERROR, "Error Parsing JSON File")
			}
		} else {
			if err := json.Unmarshal([]byte(header), &headerData); err != nil {
				logger.ConsoleLog(logger.ERROR, "Error Parsing JSON Data")
			}
		}
	} else {
		headerData = nil
	}

	return Config{
		url:        url,
		file:       file,
		thread:     thread,
		statusList: statuslist,
		header:     headerData,
		proxy:      proxy,
		filter:     filter,
		outputFile: output,
		recursive:  recursive,
		debug:      debug,
		mutex:      mutex,
	}
}
