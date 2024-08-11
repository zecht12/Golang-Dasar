package main

import "fmt"

func init() {
	var name [6]string
	name[0] = "Gilang"
	name[1] = "Prima"
	name[2] = "Ertansyah"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var tahun [3]uint16 = [3]uint16{
		28,
		10,
		2004,
	}

	fmt.Println(tahun)
	fmt.Println(len(tahun))
}