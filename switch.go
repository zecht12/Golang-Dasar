package main

import "fmt"

func init() {
	name := "Gilang"

	switch name {
	case "Gilang":
		fmt.Println("halo gilang")
	default:
		fmt.Println(`ohh jadi namamu `+name+" hiii salken")
	}

	switch length := len(name); length>5 {
	case true:
		fmt.Println("jadi namamu lebih dari 5 huruf ya")
	case false:
		fmt.Println("jadi namamu kurang dari 5 huruf ya")
	}
}