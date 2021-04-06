package main

import (
	"fmt"
	cfg "testconfig/config"
)

func main() {
	var path = "./config.ini"
	config := new(cfg.Config)
	config.InitConfig(path)
	log_print := config.Read("file", "log_print")
	fmt.Printf("log_print:%#v", log_print)

}
