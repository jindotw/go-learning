package main

import "paul.idv/golang/tree/btree"

func preOrder(tree *btree.Node) {
	btree.PreOrderRecur(tree)
}

func main() {
	preOrder(btree.ConstructTree())
}
