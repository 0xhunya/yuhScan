# yuhScan
基于Go语言实现的web目录快速扫描工具

```
            _   _____             
    _ _ _ _| |_|   __|___ ___ ___     v1.0
   | | | | |   |__   |  _| .'|   |    
   |_  |___|_|_|_____|___|__,|_|_|    GitHub: https://github.com/hunyaio/yuhScan
   |___|                              

Usage: yuhScan <-u URL>/<-uf URLFile,[-group GroupNum]> <-f File> [options]

Options: [-h] [-t Thread] [-proxy ProxyURL] [-r] [-o Output.txt] [-header JSON] [-status StatusCode] [-filter FilterString] [-debug]

Params:
  -debug
    	Debug
  -f string
    	Dictionary File to Load
  -filter string
    	Single Special String to Filter (warning strings、slogan, etc.)
  -group int
    	Grouping Targets in URL File(run all targets concurrently by default)
  -h	Help
  -header string
    	JSON Data/File of Custom HTTP Headers
  -o string
    	File to Save the Results
  -proxy string
    	Proxy (Scheme://Host:Port)
  -r	Recursive Scan
  -status string
    	Status Code to Check (concat with ',') (default "200,301")
  -t int
    	Thread to Run (default 10)
  -u string
    	URL to Scan
  -uf string
    	URL File to Scan
```

## 参数
#### -h
帮助
#### -u/-uf
指定扫描目标URL（单个）/目标URL文件（批量）
#### -f
指定使用的字典文件
#### -group
当使用 `-uf` 参数批量扫描时,指定同时扫描的目标数.默认同时扫描文件中所有URL.
#### -t
指定单目标扫描的并发数,默认为10
#### -proxy
指定代理,格式 `Scheme://Host:Port`
#### -r
启用递归扫描模式
#### -o
指定扫描结果的输出文件
#### -header
指定JSON数据格式的自定义HTTP header.用于特殊情境（cookie、token）的扫描.
#### -status
指定捕获的HTTP状态码.默认 `200,301`（用,连接）
#### -filter
指定过滤字符串.用于过滤一些带有固定特征字符串的页面,仅可指定单个简单字符串
#### -debug
开启debug模式,运行过程会输出debug信息

## 下载
```
go get github.com/hunyaio/yuhScan
```

## 编译
```
sh build.sh
```
需要Go环境在`$GOPATH`下编译，MacOS/Linux可直接执行编译脚本，Windows可通过Git Bash执行编译脚本

## 最后
纪念曾经在「玉衡」的一众一起学习和摸鱼的时光

感谢[d4m1ts](https://github.com/damit5)在这期间一直帮我测试工具，还提出了很多建议