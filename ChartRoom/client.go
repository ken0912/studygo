package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()
	//get client input and send it to server
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			conn.Write(str[:n])
		}
	}()

	//get msg from server
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)

		//server closed
		if n == 0 {
			fmt.Println("exit!")
			return
		}

		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("read msg from server:\r\n", string(buf[:n]))
	}
}
