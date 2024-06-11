package main

import (
	"advance/Tree/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

// 后序遍历
func (my *myTreeNode) postOrder() {
	if my == nil || my.node == nil {
		return
	}
	left := myTreeNode{my.node.Left}
	right := myTreeNode{my.node.Right}
	left.postOrder()
	right.postOrder()
	my.node.Print()
}

func main() {
	//    3
	//   /  \
	//  0     5
	//  \     /
	//   2  4
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetVaue(4)
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Printf("node count: %d\n", nodeCount)

}
