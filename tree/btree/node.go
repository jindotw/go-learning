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
	return node.Left == nil
}

func (node *Node) isRightEmpty() bool {
	return node.Right == nil
}

// ConstructTree
//
//	  4
//	2   6
//
// 1   3    8
// /*
func ConstructTree() *Node {
	one := &Node{nil, nil, 1}
	three := &Node{nil, nil, 3}
	eight := &Node{nil, nil, 8}

	two := &Node{one, three, 2}
	six := &Node{nil, eight, 6}

	return &Node{two, six, 4}
}
