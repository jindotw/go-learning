package btree

import "fmt"

func PostOrderRecur(root *Node) {
	if root == nil {
		return
	}
	PostOrderRecur(root.Left)
	PostOrderRecur(root.Right)
	fmt.Println(root.GetValue())
}

func PostOrderIter(root *Node) {
	st := make([]*Node, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]

		// root node
		if node.isLeftEmpty() && node.isRightEmpty() {
			st = st[:last]
			fmt.Println(node.Value)
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
