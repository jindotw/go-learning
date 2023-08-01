package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
)

/*
Demonstrates struct comparison
*/

type Foo struct {
	x   int
	y   int
	z   interface{}
	arr []int
	f   func(int)
}

// struct can be compared using == operator as long as the struct does not contain any slices, maps or function
// struct contains slices and maps can be compared using reflect.DeepEqual()
// comparison for slice of bytes can be done using bytes.Equal(), bytes.Compare() and bytes.EqualFold()

func cmpStruct() {
	var fx = func(a int) {
		fmt.Println(a, "was passed to me")
	}
	a := Foo{
		x:   10,
		y:   15,
		z:   0,
		arr: []int{1, 2, 3},
		f:   fx,
	}
	b := Foo{
		x:   10,
		y:   15,
		z:   0,
		arr: []int{1, 2, 3},
		f:   fx,
	}

	fmt.Printf("a == b ? %t\n", reflect.DeepEqual(a, b))
	fx(355)
}

func scanText() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var n int
	var sb strings.Builder
	// var text string
	for scanner.Scan() {
		n++
		// text = append(text, fmt.Sprintf("%d. %s\n", n, scanner.Text())...)
		// text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
		sb.WriteString(fmt.Sprintf("%d. %s\n", n, scanner.Text()))
	}
	fmt.Println(sb.String())
}

func main() {
	scanText()
}
