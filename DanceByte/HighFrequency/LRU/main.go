package main

type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	cap        int
	cache      map[int]*LRUNode
	head, tail *LRUNode
}

func Constructor(capacity int) LRUCache {
	head, tail := new(LRUNode), new(LRUNode)
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*LRUNode, capacity),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToFront(node)
	} else {
		node := &LRUNode{
			key:   key,
			value: value,
		}
		this.pushFront(node)
		this.cache[key] = node
		if len(this.cache) > this.cap {
			tailData := this.tail.prev
			this.remove(tailData)
			delete(this.cache, tailData.key)
		}
	}
}

// 移到链表头部
func (this *LRUCache) moveToFront(node *LRUNode) {
	this.remove(node)
	this.pushFront(node)
}

// 移除节点
func (this *LRUCache) remove(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 新节点添加
func (this *LRUCache) pushFront(node *LRUNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
