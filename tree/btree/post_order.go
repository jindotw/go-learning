package btree

import "fmt"

func PostOrderRecur(root *Node) {
	if root != nil {
		PostOrderRecur(root.Left)
		PostOrderRecur(root.Right)
		fmt.Println(root.GetValue())
	}
}

func PostOrderIter(root *Node) {
	st := make([]*Node, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]

		if node.isLeftEmpty() && node.isRightEmpty() {
			fmt.Println(node.GetValue())
			st = st[:last]
		} else {
			if !node.isRightEmpty() {
				st = append(st, node.Right)
				node.Right = nil
			}
			if !node.isLeftEmpty() {
				st = append(st, node.Left)
				node.Left = nil
			}
		}
	}
}

func PostOrderCommon(root *Node) {
	st := make([]*Node, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		node := pop(&st)

		if node != nil {
			st = append(st, node, nil)
			if !node.isRightEmpty() {
				st = append(st, node.Right)
			}
			if !node.isLeftEmpty() {
				st = append(st, node.Left)
			}
		} else {
			fmt.Println(pop(&st).GetValue())
		}
	}
}
