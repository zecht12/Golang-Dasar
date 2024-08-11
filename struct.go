package main

import "fmt"

type Customer struct{
	id		int
	name 	string
	age 	uint8
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name,"my name is", customer.name)
}

func main() {
	var gilang Customer

	gilang.id = 1
	gilang.name = "Gilang"
	gilang.age = 19

	fmt.Println(gilang)
	gilang.sayHello("Gilang")
}