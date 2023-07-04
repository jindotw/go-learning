package main

import "fmt"

func copySlice() {
	arr1 := []int{1, 2, 3}
	arr2 := make([]int, len(arr1))
	arr3 := arr1
	num := copy(arr2, arr1)
	fmt.Printf("copied %d elements\n", num)
	arr2[0] = 100
	fmt.Printf("Modify arr2: %#v\n", arr1)
	arr3[0] = 100
	fmt.Printf("Modify arr3: %#v\n", arr1)
}

func copyMap() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := make(map[string]int, len(m1))

	for k, v := range m1 {
		m2[k] = v
	}

	fmt.Printf("%#v\n", m1)
	fmt.Printf("%#v\n", m2)
}

func copyInterface() {
	// interface has different concrete implementations and thus not easily copiable
}

func main() {
	copySlice()
	copyMap()
}
