package main

import "fmt"

func sayHello(nama string, umur uint8)  {
	fmt.Println("Hallo nama saya",nama,"umur",umur)
}

func newHello(nama string) string {
	return "halo "+nama
}

func init() {
	sayHello("Gilang",18)

	result := newHello("Gilang")

	fmt.Println(result)
}

