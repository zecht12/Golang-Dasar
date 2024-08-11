package main

import "fmt"

func init() {
	type noKtp string
	var ktp noKtp = "11111111"

	fmt.Println(ktp)
	fmt.Println(noKtp("22222222"))
}