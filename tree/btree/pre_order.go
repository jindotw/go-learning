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
