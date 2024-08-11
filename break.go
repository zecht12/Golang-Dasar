package main

import "fmt"

func init() {
	count := 50

	for i := 1; i <= count; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}