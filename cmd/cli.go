package cmd

import (
	"flag"
	"fmt"
)

var (
	help      bool
	url       string
	ufile     string
	group     int
	file      string
	thread    int
	proxy     string
	header    string
	status    string
	recursive bool
	filter    string
	output    string
	debug     bool
)

func Help() bool      { return help }
func Url() string     { return url }
func Ufile() string   { return ufile }
func Group() int      { return group }
func File() string    { return file }
func Thread() int     { return thread }
func Proxy() string   { return proxy }
func Header() string  { return header }
func Status() string  { return status }
func Recursive() bool { return recursive }
func Filter() string  { return filter }
func Output() string  { return output }
func Debug() bool     { return debug }

func init() {
	flag.BoolVar(&help, "h", false, "Help")
	flag.StringVar(&url, "u", "", "URL to Scan")
	flag.StringVar(&ufile, "uf", "", "URL File to Scan")
	flag.IntVar(&group, "group", 0, "Grouping Targets in URL File(run all targets concurrently by default)")
	flag.StringVar(&file, "f", "", "Dictionary File to Load")
	flag.IntVar(&thread, "t", 10, "Thread to Run")
	flag.StringVar(&proxy, "proxy", "", "Proxy (Scheme://Host:Port)")
	flag.StringVar(&output, "o", "", "File to Save the Results")
	flag.StringVar(&header, "header", "", "JSON Data/File of Custom HTTP Headers")
	flag.StringVar(&status, "status", "200,301", "Status Code to Check (concat with ',')")
	flag.StringVar(&filter, "filter", "", "Single Special String to Filter (warning strings、slogan, etc.)")
	flag.BoolVar(&recursive, "r", false, "Recursive Scan")
	flag.BoolVar(&debug, "debug", false, "Debug")

	flag.Usage = Usage
	flag.Parse()
}

func Usage() {
	fmt.Println(`
            _   _____             
    _ _ _ _| |_|   __|___ ___ ___     v1.0
   | | | | |   |__   |  _| .'|   |    
   |_  |___|_|_|_____|___|__,|_|_|    GitHub: https://github.com/hunyaio/yuhScan
   |___|                              

Usage: yuhScan <-u URL>/<-uf URLFile,[-group GroupNum]> <-f File> [options]

Options: [-h] [-t Thread] [-proxy ProxyURL] [-r] [-o Output.txt] [-header JSON] [-status StatusCode] [-filter FilterString] [-debug]

Params:`)
	flag.PrintDefaults()
}
