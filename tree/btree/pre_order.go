package btree

import "fmt"

func PreOrderRecur(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.GetValue())
	PreOrderRecur(root.Left)
	PreOrderRecur(root.Right)
}

func PreOrderIter(root *Node) {
	// stack if FILO
	st := make([]*Node, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]
		st = st[:last]
		fmt.Println(node.Value)
		if !node.isRightEmpty() {
			st = append(st, node.Right)
		}
		if !node.isLeftEmpty() {
			st = append(st, node.Left)
		}
	}
}

func PreOrderCommon(root *Node) {
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
			if !node.isLeftEmpty() {
				st = append(st, node.Left)
			}

			st = append(st, node)
			st = append(st, nil)
		} else {
			node = pop(&st)
			fmt.Println(node.GetValue())
		}
	}
}
