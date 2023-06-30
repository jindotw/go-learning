package btree

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (node *Node) GetValue() int {
	return node.Value
}

func (node *Node) isLeftEmpty() bool {
	return node.Left != nil
}

func (node *Node) isRightEmpty() bool {
	return node.Right != nil
}

func ConstructTree() *Node {
	lft := &Node{nil, nil, 1}
	rgt := &Node{nil, nil, 3}
	return &Node{lft, rgt, 2}
}
