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

func inOrder(tree *btree.Node) {
	btree.InOrderRecur(tree)
	fmt.Println()
	btree.InOrderIter(tree)
}

func postOrder(tree *btree.Node) {
	btree.PostOrderRecur(tree)
	fmt.Println()
	btree.PostOrderIter(tree)
}

func main() {
	postOrder(btree.ConstructTree())
}
