package main

import "fmt"

type Point struct {
	x, y float64
}

func (pt *Point) AddX(dx float64) {
	pt.x += dx
}

func main() {
	p1 := &Point{3, 4}
	p1.AddX(2.5)
	if p1.x > p1.y {
		fmt.Printf("p1.x (%f) > p1.y (%f)\n", p1.x, p1.y)
	} else {
		fmt.Printf("p1.y (%f) > p1.x (%f)\n", p1.y, p1.x)
	}
	fmt.Println(p1)
}
