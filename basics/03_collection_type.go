package main

import (
	"fmt"
)

type user struct {
	name    string
	age     int
	address string
}

func main() {

	// struct
	var person1 user
	person1.name = "osama"
	fmt.Print(person1)

	p1 := user{name: "ali"}

	fmt.Println(p1)

}
