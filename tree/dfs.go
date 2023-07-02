package main

import (
	"fmt"
	"paul.idv/golang/tree/btree"
)

func preOrder(tree *btree.Node) {
	btree.PreOrderRecur(tree)
	fmt.Println()
	btree.PreOrderIter(tree)
	fmt.Println()
	btree.PreOrderCommon(tree)
}

func inOrder(tree *btree.Node) {
	btree.InOrderRecur(tree)
	fmt.Println()
	btree.InOrderIter(tree)
	fmt.Println()
	btree.InOrderCommon(tree)
}

func postOrder(tree *btree.Node) {
	btree.PostOrderRecur(tree)
	fmt.Println()
	btree.PostOrderIter(tree)
	fmt.Println()
	btree.PostOrderCommon(btree.ConstructTree())
}

func main() {
	postOrder(btree.ConstructTree())
}
