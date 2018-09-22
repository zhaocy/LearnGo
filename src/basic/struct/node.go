package main

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	value       int
}

func (node TreeNode) print() {
	fmt.Println(node.value)
}

func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored.")
		return
	}
	node.value = value
}

//工厂函数
func createNode(value int) *TreeNode {
	return &TreeNode{nil, nil, value}
}

func main() {
	//创建treenode
	//var root treeNode
	root := TreeNode{value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{nil, nil, 6}

	root.Right.Left = new(TreeNode)
	fmt.Println(root)

	nodes := []TreeNode{
		{},
		{nil, &root, 7},
		{value: 3},
	}
	fmt.Println(nodes)

	tn := createNode(10)
	fmt.Println(tn)

	root.setValue(10)
	root.print()

	var pRoot TreeNode
	p := &pRoot
	p.setValue(300)
	p.print()
}
