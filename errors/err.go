package main

import (
	"errors"
	"fmt"
)

type MyInt int

func (a MyInt) Error() string {
	return fmt.Sprintf("%s %d", "custom type error", int(a))
}

func doIt() error {
	var a = 5
	var b = MyInt(a)
	return b
}

type MyFloat float64

func (a MyFloat) Error() string {
	return fmt.Sprintf("%s %f", "custom type error", float64(a))
}

func main() {
	if err := doIt(); err != nil {
		var (
			a MyInt
			b MyFloat
		)
		if errors.As(err, &a) {
			fmt.Println("err is of MyInt")
		}
		if errors.As(err, &b) {
			fmt.Println("err is of MyFloat")
		}
	}
}
