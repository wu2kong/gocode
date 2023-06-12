package main

import (
	"fmt"
	"sync"
)

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

// 构建先进先出队列链表
type LinkNode struct {
	Value *TreeNode
	Next  *LinkNode
}

// 链表队列
type LinkQueue struct {
	head *LinkNode
	size int
	lock sync.Mutex
}

// 入队
func (queue *LinkQueue) Insert(node *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果队列为空，那么增加节点
	if queue.head == nil {
		queue.head = &LinkNode{Value: node}
	} else {
		newNode := &LinkNode{Value: node}
		// 找到尾部指针
		tailNode := queue.head
		for tailNode.Next != nil {
			tailNode = tailNode.Next
		}

		tailNode.Next = newNode
	}

	queue.size++
}

// 出队
func (queue *LinkQueue) Pop() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		return nil
	}

	topNode := queue.head

	queue.head = topNode.Next
	queue.size--

	return topNode.Value
}

// 队列大小
func (queue *LinkQueue) Size() int {
	return queue.size
}

// 层次遍历
func LayerOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := new(LinkQueue)

	// 每个节点：节点出队时遍历、节点的左孩子入队、节点的右孩子入队
	queue.Insert(root)

	// 开始遍历
	for queue.size > 0 {
		element := queue.Pop()
		fmt.Printf("%d ", element.Data)

		if element.Left != nil {
			queue.Insert(element.Left)
		}

		if element.Right != nil {
			queue.Insert(element.Right)
		}
	}
}

func main() {
	treeRootNode := BuildTree()
	LayerOrder(treeRootNode)
}
