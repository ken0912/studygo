package main

import (
	"bytes"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var buffer bytes.Buffer
var password string

func main() {
	password = "ken083027"
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("hasePassword:", string(hasePassword))
	sum := 0

}

func test() {
	s := time.Now().Unix()
	or i := 0; i < 100000; i++ {
	buffer.WriteString("hello")
	}

	e := time.Now().Unix()
	fmt.Println("耗时:", e-s)
}
