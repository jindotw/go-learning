package main

import (
	"fmt"
	"time"
)

func start(ch chan int) {
	fmt.Println("In go routine")
	ch <- 3
	close(ch)
}

func simpleChan() {
	ch := make(chan int)
	fmt.Println("In main, about to send to channel")
	go start(ch)
	time.Sleep(50 * time.Millisecond)
	val := <-ch
	fmt.Println("Exiting main with chan val", val)
}
