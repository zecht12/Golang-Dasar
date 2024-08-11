package main

import "fmt"

func init() {
	a := 10
	b := 20
	c := a+b
	d := c-b
	e := c/d
	f := e*d
	g := f%a
	
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	a--
	if a == 10 {
		fmt.Println("nih ketemu a =",a)
	} else {
		fmt.Printf("wah coba lagi")
	}
}