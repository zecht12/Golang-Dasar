package main

import "fmt"

func init() {
	count := 10

	for i := 1; i <= count; i++ {
		fmt.Println(i)
	}

	name := []string{"gilang", "prima", "ertansyah"}

	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	for index, names := range name{
		fmt.Println("index",index,"name",names)
	}
}