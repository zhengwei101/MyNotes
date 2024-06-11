package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetVaue(value int) {
	node.Value = value
}

func CreateNode(v int) *Node {
	return &Node{Value: v}
}
