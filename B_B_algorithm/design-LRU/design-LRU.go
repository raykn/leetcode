package design_LRU

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode // 附带队头和队尾 使所有节点必有前后
}

// TODO 146. LRU 缓存机制
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

// Get
// 找到要更新位置 + 删除原来位置 + 更新到队列头
func (this *LRUCache) Get(key int) int {
	cache, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveHead(cache)
	return cache.value
}

// Put
// 找不到要插入 + 插入后要判断size + 清cache后要 size--
func (this *LRUCache) Put(key int, value int) {
	cache, ok := this.cache[key]
	if !ok {
		cache = initDLinkedNode(key, value)
		this.moveHead(cache)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}
	cache.value = value
	this.cache[key] = cache
	return
}
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
func (this *LRUCache) moveHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
