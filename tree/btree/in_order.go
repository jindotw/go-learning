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
			curr = pop(&st)
			fmt.Println(curr.GetValue())
			curr = curr.Right
		}
	}
}

func InOrderCommon(root *Node) {
	st := make([]*Node, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		node := pop(&st)
		if node != nil {
			if !node.isRightEmpty() {
				st = append(st, node.Right)
			}
			st = append(st, node, nil)
			if !node.isLeftEmpty() {
				st = append(st, node.Left)
			}
		} else {
			node = pop(&st)
			fmt.Println(node.GetValue())
		}
	}
}
