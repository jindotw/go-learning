package main

import "fmt"

type Point struct {
	x, y float64
}

func (pt Point) AddX(dx float64) {
	pt.x += dx
}

func main() {
	var p1 = &Point{3, 4}
	(*p1).AddX(2.5)
	fmt.Println(p1)
}
