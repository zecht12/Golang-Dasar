package main

import "fmt"

func init() {
	mapPerson := map[string]string{
		"firstname":"gilang", 
		"age": "18",
	}

	fmt.Println(mapPerson["firstname"])
}