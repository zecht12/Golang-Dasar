package main

import "fmt"

func init() {
	names := [...]string{ "Gilang", "Prima", "Ertansyah", "Nathanoel", "Aroon", "Chow", "Wildan", "Kurnia" }
	slice := names[3:6]

	fmt.Println(slice)
	fmt.Println(slice[1])

	slice[0] = "Gilang"
	newSlice := append(slice,"Haidar")

	fmt.Println(slice)
	fmt.Println(newSlice)

	slice2 := []int{1,2,3,4,5,6}

	fmt.Println(slice2)
}