package main

import (
	"fmt"
	"strconv"
)

type User struct {
	ID   int
	name string
}

func main() {
	var userlist []*User

	for i := 1; i <= 10; i++ {
		var u = &User{}
		u.ID = i
		u.name = "a" + strconv.Itoa(i)
		userlist = append(userlist, u)
	}
	fmt.Println("userlist:", &userlist)
	userlist[0].ID = 10
	for _, v := range userlist {
		fmt.Println("v:", *v)
	}
}
