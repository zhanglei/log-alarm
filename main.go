// logAlarm project main.go
package main

import (
	"fmt"
	"os"
	"time"
)

var (
	//log 文件列表
	loglist [4]*Log
	//最后日期
	lastDateStr string
)

//日志文件名格式
///home/wwwroot/api.bq.cn/data/log/sql-20141226.php

//const (
//	apiPath    string = "/home/wwwroot/api.bq.cn/data/log/"
//	swoolePath string = "/home/www/xservice/data/log/"
//)

const (
	apiPath    string = "/Users/yky/Documents/wwwroot/new-beijing/data/log/"
	swoolePath string = "/Users/yky/Documents/wwwroot/xservice/data/log/"
)

const (
	emYky = "568089266@qq.com"
)

const logExt = ".php"

func init() {
	loglist[0] = Newlog("sql", apiPath, []string{emYky})
	loglist[1] = Newlog("order", apiPath, []string{emYky})
	loglist[2] = Newlog("xclient", apiPath, []string{emYky})

	loglist[3] = Newlog("log_error", swoolePath, []string{emYky})
}

func getCurDateStr() string {
	curtime := time.Now()

	curDateStr := curtime.Format("20060102")

	if lastDateStr == "" {
		lastDateStr = curDateStr
	}

	if lastDateStr != curDateStr {

		for _, log := range loglist {
			fmt.Println(log)
			log.LastPos = 0
		}

		lastDateStr = curDateStr
	}

	return lastDateStr
}

func listenRun() {

	dateStr := getCurDateStr()

	for _, log := range loglist {

		file := log.FileName

		log_order := file + "_" + dateStr + logExt
		curLastPos := log.LastPos

		fp, err := os.Open(log.Dir + log_order)

		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}

		defer fp.Close()

		buf := make([]byte, 1024)
		content := "线上服务器(Api)发现错误, 扫描时间" + time.Now().Format("2006-01-02 15:04:05") + "\r\n错误内容如下: \r\n\r\n"

		n, _ := fp.ReadAt(buf, curLastPos)

		if n != 0 {
			content = content + string(buf)
			curLastPos += int64(n)
			for {
				n, _ = fp.ReadAt(buf, curLastPos)
				if n == 0 {
					break
				}
				content = content + string(buf[0:n])
				curLastPos += int64(n)
			}

			//fmt.Println("warning...", curLastPos)
			err = SendMail(log.Email[:], file, content)
			if err != nil {
				//fmt.Println(file, "发送失败", time.Now().Format("2006-01-02 15:04:05"))
			} else {
				//.Println(file, "发送成功", time.Now().Format("2006-01-02 15:04:05"))
			}
		}

		log.LastPos = curLastPos

		time.Sleep(3 * time.Second)

	}

	time.Sleep(10 * time.Second)

}

func main() {
	fmt.Println("listening online log, If they go wrong, we will send a email")
	fmt.Println("start on", time.Now().Format("2006-01-02 15:04:05"))
	for {
		listenRun()
	}
}
