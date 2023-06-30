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
