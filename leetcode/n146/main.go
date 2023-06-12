package main

type LRUCache struct {
	Cap        int
	Keys       map[int]*Node
	head, tail *Node
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Cap: capacity, Keys: make(map[int]*Node)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.MoreToFront(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.MoreToFront(node)
		return
	} else {
		node = &Node{Val: value, Key: key}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) MoreToFront(n *Node) {
	this.Remove(n)
	this.Add(n)
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
