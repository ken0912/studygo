package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	if err != nil {
		fmt.Println("err:", err)
	}
	c.Start()

	select {}
}
