package main

import "fmt"

type person struct {
	name   string
	age    int16
	gender string
}

func main() {
	var p person
	p.name = "ken"
	p.age = 18
	p.gender = "man"
	fmt.Println("p:", p)
	fmt.Println("********************************")
	var p1 = new(person)
	p1.name = "aaa"
	fmt.Printf("%v\n", p1)
	fmt.Printf("%p", p1)
}
