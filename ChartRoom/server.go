package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//user struct
type Client struct {
	C    chan string
	Name string
	Addr string
}

//Global map to store online client
var onlinMap map[string]Client

//Global channel send user message
var message = make(chan string)

func WriteMsgToClient(client Client, conn net.Conn) {
	//listen user's own channel
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func FmtMsg(client Client, msg string) (buf string) {
	buf = "[" + client.Addr + "]<" + client.Name + ">: " + msg
	return
}

func HandleConnect(conn net.Conn) {
	defer conn.Close()

	active := make(chan bool)

	//get client info
	netAddr := conn.RemoteAddr().String()
	//creat client
	client := Client{C: make(chan string), Name: "", Addr: netAddr}
	//store client to global map
	onlinMap[netAddr] = client

	//create a goroutine to send msg to client
	go WriteMsgToClient(client, conn)
	//send user login msg to global channel
	// loginMsg := "[" + netAddr + "]<" + client.Name + ">login!"
	loginMsg := FmtMsg(client, "login")
	message <- loginMsg

	isQuit := make(chan bool)
	//process msg user sent
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Printf("client %s out!\r\n", client.Name)
				isQuit <- true
				return
			}
			if err != nil {
				fmt.Println("conn.Read err:", err)
			}
			msg := string(buf[:n])

			//get online user list
			switch {
			case msg == "who\r\n" && len(msg) == 5:
				conn.Write([]byte("The online user list as below:\r\n"))
				for _, client := range onlinMap {
					userInfo := client.Addr + ":" + client.Name + "\r\n"
					conn.Write([]byte(userInfo))
				}
			case len(msg) >= 8 && msg[:6] == "rename":
				newName := strings.Split(msg, "|")[1]
				client.Name = newName[:len(newName)-2]
				onlinMap[netAddr] = client
				conn.Write([]byte("rename successful!\r\n"))
			default:
				if client.Name == "" {
					conn.Write([]byte("please give a name for youself"))
				} else {
					message <- FmtMsg(client, msg)
				}
			}

			active <- true
		}
	}()
	for {
		select {
		case <-isQuit:
			close(client.C)
			delete(onlinMap, client.Addr)
			message <- FmtMsg(client, "logout")
			return
		case <-active:
			//do not anything
		case <-time.After(time.Minute * 1):
			delete(onlinMap, client.Addr)
			message <- FmtMsg(client, "logout")
			return
		}
	}
}

func Manager() {
	//init global map
	onlinMap = make(map[string]Client)

	for {
		//listen global channel
		msg := <-message

		//send msg to online client
		for _, client := range onlinMap {
			client.C <- msg
		}
	}

}

func main() {
	//create listen socket
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	//close socket
	defer ln.Close()

	//goroutine to manager global map and channel
	go Manager()

	//loop listen client connection request
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ln.Accept error:", err)
			continue
		}
		go HandleConnect(conn)
	}
}
