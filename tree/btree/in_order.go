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
