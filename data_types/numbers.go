package main

import "fmt"

type MyFloat float64

func (num *MyFloat) Add(addend MyFloat) {
	*num += addend
}

func main() {
	a := MyFloat(2.2)
	a.Add(1.1)
	fmt.Println(a)
}
