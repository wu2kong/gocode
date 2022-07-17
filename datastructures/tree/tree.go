package main

import "fmt"

// 二叉树的节点
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree() *TreeNode {
	rootNode := TreeNode{Data: 1}

	rootNode.Left = &TreeNode{Data: 2}
	rootNode.Left.Left = &TreeNode{Data: 4}
	rootNode.Left.Right = &TreeNode{Data: 5}

	rootNode.Right = &TreeNode{Data: 3}
	rootNode.Right.Left = &TreeNode{Data: 6}
	rootNode.Right.Right = &TreeNode{Data: 7}

	return &rootNode
}

// 访问元素
func ProcessNode(t *TreeNode) {
	fmt.Printf("%d ", t.Data)
}

// 先序遍历
func PreOrderTraverse(t *TreeNode) {
	if t != nil {
		ProcessNode(t)
		PreOrderTraverse(t.Left)
		PreOrderTraverse(t.Right)
	}
	return
}

// 中序遍历
func InOrderTraverse(t *TreeNode) {
	if t != nil {
		InOrderTraverse(t.Left)
		ProcessNode(t)
		InOrderTraverse(t.Right)
	}
	return
}

// 后序遍历
func PostOrderTraverse(t *TreeNode) {
	if t != nil {
		PostOrderTraverse(t.Left)
		PostOrderTraverse(t.Right)
		ProcessNode(t)
	}
	return
}

func main() {
	treeRootNode := BuildTree()

	PreOrderTraverse(treeRootNode)
	println()
	// 1 2 4 5 3 6 7

	InOrderTraverse(treeRootNode)
	println()
	// 4 2 5 1 6 3 7

	PostOrderTraverse(treeRootNode)
	println()
	// 4 5 2 6 7 3 1
}
