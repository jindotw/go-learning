package btree

import "fmt"

func PreOrderRecur(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.GetValue())
	PreOrderRecur(node.Left)
	PreOrderRecur(node.Right)
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
