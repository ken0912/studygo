package main

import (
	"mylogger/logger"
)

func main() {
	// log := logger.NewConsoleLog("Warning")
	log := logger.NewFileLog("debug", "./", "log.log", 1024*10*1024)

	for i := 1; i <= 26; i++ {
		// fmt.Println(i)
		log.Info("成功写入%v/%v", i, 26)
	}
}
