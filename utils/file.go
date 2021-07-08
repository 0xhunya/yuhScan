package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/hunyaio/yuhScan/logger"
)

func FileExits(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func FileRead(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.ConsoleLog(logger.ERROR, "Error Reading File")
	}
	return content
}

func FileWrite(filename string, s string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		logger.ConsoleLog(logger.ERROR, err)
	}
	defer f.Close()
	f.Write([]byte(s))
}

func FileCreateAndWrite(filename string, s string) {
	f, err := os.Create(filename)
	if err != nil {
		logger.ConsoleLog(logger.ERROR, err)
		return
	}
	defer f.Close()
	f.WriteString(s)
}

func LoadDictOne(filename string) []string {
	var payloads []string

	// Open Dictionary File
	file, err := os.Open(filename)
	if err != nil {
		logger.ConsoleLog(logger.ERROR, "Open File Failed:", err)
		return nil
	}
	defer file.Close()

	// Read Data
	reader := bufio.NewReader(file)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				logger.ConsoleLog(logger.ERROR, err)
			}
		}
		if len(string(data)) > 0 {
			payloads = append(payloads, string(data))
		}
	}

	if len(payloads) == 0 {
		logger.ConsoleLog(logger.ERROR, "Empty File")
	}

	return payloads
}

func LoadPayloadsFromFile(filename string) []string {
	payloads := LoadDictOne(filename)
	for key, payload := range payloads {
		if payload[0] != '/' {
			payloads[key] = "/" + payload
		}
	}
	return payloads
}

func OutputToFile(filename string, s string) {
	if FileExits(filename) {
		FileWrite(filename, s)
	} else {
		FileCreateAndWrite(filename, s)
	}
}
