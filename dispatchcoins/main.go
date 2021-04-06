package main

import (
	"fmt"
)

/*
Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth

contain 'e'or'E' dispatch 1 gold
contain 'i'or'I' dispatch 2 gold
contain 'o'or'O' dispatch 3 gold
contain 'u'or'U' dospatch 4 gold
*/

var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("left:", left)
	fmt.Println("distribution", distribution)
}

func dispatchCoin() int {
	for _, name := range users {
		distribution[name] = 0
		for _, n := range name {
			switch n {
			case 'e', 'E':
				distribution[name] += 1
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	return coins
}
