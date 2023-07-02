package main

import (
	"fmt"
)

func start(obj chan int) {
	fmt.Println("In go routine")
	obj <- 3
}

func main() {
	obj := make(chan int)
	fmt.Println("In main, about to send to channel")
	go start(obj)
	val := <-obj
	fmt.Println("Exiting main with chan val", val)
}
