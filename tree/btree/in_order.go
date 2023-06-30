package btree

import "fmt"

func InOrderRecur(root *Node) {
	if root == nil {
		return
	}
	InOrderRecur(root.Left)
	fmt.Println(root.GetValue())
	InOrderRecur(root.Right)
}

func InOrderIter(root *Node) {
	st := make([]*Node, 0, 10)
	curr := root

	for curr != nil || len(st) > 0 {
		if curr != nil {
			st = append(st, curr)
			curr = curr.Left
		} else {
			last := len(st) - 1
			curr = st[last]
			st = st[:last]
			fmt.Println(curr.Value)
			curr = curr.Right
		}
	}
}
