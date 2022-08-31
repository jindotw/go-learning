package main

import "fmt"

func main() {
	s := "BMW"

	switch s {
	case "Lexus":
		fmt.Println("This is a Lexus")
		break
	case "BMW":
		fmt.Println("This is a BMW")
		break
	case "Porsche":
		fmt.Println("This is a Porsche")
		break
	}
}
