package main

import (
	"fmt"
	"paul.idv/golang/tree/btree"
)

func preOrder(tree *btree.Node) {
	btree.PreOrderRecur(tree)
	fmt.Println()
	btree.PreOrderIter(tree)
}

func postOrder(tree *btree.Node) {
	btree.PostOrderRecur(tree)
}

func main() {
	postOrder(btree.ConstructTree())
}
