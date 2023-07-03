package main

import "fmt"

func breakSelect() {
	ch := make(chan string)
	// ch <- "Before break"

	select {
	case msg := <-ch:
		fmt.Println(msg)
		// break
		fmt.Println("After break")
	default:
		fmt.Println("Default case")
	}
}

func simpleSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go go1(ch1)
	go go2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case val1 := <-ch1:
			fmt.Println("received", val1, "from chan1")
		case val2 := <-ch2:
			fmt.Println("received", val2, "from chan1")
		}
	}
}

func go1(ch chan string) {
	ch <- "from go1"
	close(ch)
}

func go2(ch chan string) {
	ch <- "from go2"
	close(ch)
}
