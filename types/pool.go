package types

import (
	"time"

	"yuhScan/logger"
)

type Worker interface {
	Working()
}

type Task struct {
	Args       map[string]interface{}
	MsgChan    chan string
	ResultChan chan interface{}
}

type Pool struct {
	num        int
	result     []interface{}
	TaskChan   chan Worker
	MsgChan    chan string
	ResultChan chan interface{}
}

func NewPool() *Pool {
	return &Pool{
		num:        0,
		TaskChan:   make(chan Worker),
		MsgChan:    make(chan string, 50),
		ResultChan: make(chan interface{}),
	}
}

func (p *Pool) Num() int              { return p.num }
func (p *Pool) Result() []interface{} { return p.result }

func (p *Pool) Run() {
	logger.ConsoleLog(logger.STATUS, "Starting Worker-Pool...")
	p.num = 0
	p.result = []interface{}{}
ReceiveTask:
	for {
		select {
		case worker := <-p.TaskChan:
			go worker.Working()
			p.num += 1
		case <-time.After(time.Second * 1):
			break ReceiveTask
		}
	}

	for len(p.result) != p.num {
		select {
		case res := <-p.ResultChan:
			p.result = append(p.result, res)
		case msg := <-p.MsgChan:
			logger.ConsoleLog(logger.STATUS, msg)
		}
	}
	logger.ConsoleLog(logger.STATUS, "Worker-Pool Finished")
}
