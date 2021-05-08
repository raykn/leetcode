package main

import (
	"math"
	"strconv"
)

// todo 564. 寻找最近的回文数
func nearestPalindromic(n string) string {
	target, _ := strconv.Atoi(n)
	var num10, numTmp, res = true, target, target
	if numTmp%10 == 1 && target > 10 {
		numTmp -= 1
		res = numTmp
	}
	for i := 0; i < len(n)-1; i++ {
		t := numTmp % 10
		numTmp /= 10
		if t > 0 {
			num10 = false
			break
		}
	}
	switch {
	case target <= 10, num10:
		return strconv.Itoa(res - 1)
	case target == 11:
		return "9"
	}

	var left, right string
	mid := getByteToRes(n)
	switch compare(mid, n) {
	case -1:
		left = mid
		right = getByteToRes(getTargetStr(n, 1))

	case 0:
		left = getByteToRes(getTargetStr(n, -1))
		right = getByteToRes(getTargetStr(n, 1))

	case 1:
		left = mid
		right = getByteToRes(getTargetStr(n, -1))
	}

	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)
	one, two := math.Abs(float64(leftNum-target)), math.Abs(float64(rightNum-target))
	if one < two {
		return left

	} else if int(one) == int(two) {
		if compare(left, right) < 0 {
			return left
		}
		return right

	} else {
		return right
	}
}
func getByteToRes(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes)/2; i++ {
		bytes[len(bytes)-1-i] = bytes[i]
	}
	return string(bytes)
}
func compare(one, two string) int {
	oneInt, _ := strconv.Atoi(one)
	twoInt, _ := strconv.Atoi(two)
	if oneInt > twoInt {
		return 1
	} else if oneInt == twoInt {
		return 0
	}
	return -1
}
func getTargetStr(target string, offset int) string {
	bytes := []byte(target)
	i, j := 0, len(bytes)-1
	for i <= j {
		i++
		j--
	}
	lInt, _ := strconv.Atoi(string(bytes[0:i]))
	lStr := strconv.Itoa(lInt + offset)
	if i == len(bytes) {
		return lStr // 1位
	}
	rStr := string(bytes[i:])
	return lStr + rStr
}

// todo 460 LFU 缓存
// type LFUData struct {
// 	key   int
// 	value int
// 	count int
// }
// type LFUCache struct {
// 	cap      int
// 	dataMap  map[int]*list.Element
// 	countMap map[int]*list.List
// 	minCount int
// }
//
// func Constructor(capacity int) LFUCache {
// 	return LFUCache{
// 		cap:      capacity,
// 		dataMap:  make(map[int]*list.Element, capacity),
// 		countMap: make(map[int]*list.List, capacity),
// 	}
// }
// func (this *LFUCache) Get(key int) int {
// 	v, ok := this.dataMap[key]
// 	if !ok {
// 		return -1
// 	}
// 	return this.getAndUpdate(v)
// }
// func (this *LFUCache) Put(key int, value int) {
// 	data, ok := this.dataMap[key]
// 	if ok {
// 		d := data.Value.(*LFUData)
// 		d.value = value
// 		this.getAndUpdate(data)
// 		return
// 	}
// 	if len(this.dataMap) == this.cap {
// 		minList, ok := this.countMap[this.minCount]
// 		if !ok {
// 			return
// 		}
// 		delData := minList.Front()
// 		minList.Remove(delData)
// 		d := delData.Value.(*LFUData)
// 		delete(this.dataMap, d.key)
// 	}
// 	newData := &LFUData{
// 		key:   key,
// 		value: value,
// 		count: 1,
// 	}
// 	newList, ok := this.countMap[newData.count]
// 	if !ok {
// 		newList = list.New()
// 		this.countMap[newData.count] = newList
// 	}
// 	newValue := newList.PushBack(newData)
// 	this.dataMap[key] = newValue
// 	this.minCount = newData.count
// }
// func (this *LFUCache) getAndUpdate(v *list.Element) int {
// 	data := v.Value.(*LFUData)
// 	oldList, ok := this.countMap[data.count]
// 	if !ok {
// 		return -1
// 	}
// 	oldList.Remove(v)
// 	if oldList.Len() == 0 && data.count == this.minCount {
// 		this.minCount++
// 	}
// 	data.count++
// 	newList, ok := this.countMap[data.count]
// 	if !ok {
// 		newList = list.New()
// 		this.countMap[data.count] = newList
// 	}
// 	newData := newList.PushBack(data)
// 	this.dataMap[data.key] = newData
// 	return data.value
// }

// type Node struct {
// 	key, val, freq int
// 	prev, next     *Node
// }
// type DoubleList struct {
// 	head, tail *Node
// }
// type LFUCache struct {
// 	cache              map[int]*Node
// 	freq               map[int]*DoubleList
// 	cap, size, minFreq int
// }
//
// func Constructor(capacity int) LFUCache {
// 	return LFUCache{
// 		cache: make(map[int]*Node),
// 		freq:  make(map[int]*DoubleList),
// 		cap:   capacity,
// 	}
// }
//
// func (this *LFUCache) Get(key int) int {
// 	if node, ok := this.cache[key]; ok {
// 		this.IncFreq(node)
// 		return node.val
// 	}
// 	return -1
// }
//
// func (this *LFUCache) Put(key int, value int) {
// 	if this.cap == 0 {
// 		return
// 	}
// 	if node, ok := this.cache[key]; ok {
// 		node.val = value
// 		this.IncFreq(node)
// 	} else {
// 		if this.size >= this.cap {
// 			node := this.freq[this.minFreq].RemoveLast()
// 			delete(this.cache, node.key)
// 			this.size++
// 		}
// 		x := &Node{key: key, val: value, freq: 1}
// 		this.cache[key] = x
// 		if this.freq[1] == nil {
// 			this.freq[1] = CreateDL()
// 		}
// 		this.freq[1].AddFirst(x)
// 		this.minFreq = 1
// 		this.size++
// 	}
// }
//
// func (this *LFUCache) IncFreq(node *Node) {
// 	currFreq := node.freq
// 	this.freq[currFreq].Remove(node)
// 	if this.minFreq == currFreq && this.freq[currFreq].IsEmpty() {
// 		this.minFreq++
// 		delete(this.freq, currFreq)
// 	}
// 	node.freq++
// 	if this.freq[node.freq] == nil {
// 		this.freq[node.freq] = CreateDL()
// 	}
// 	this.freq[node.freq].AddFirst(node)
// }
//
// func CreateDL() *DoubleList {
// 	head, tail := &Node{}, &Node{}
// 	head.next, tail.prev = tail, head
// 	return &DoubleList{
// 		head: head,
// 		tail: tail,
// 	}
// }
//
// func (this *DoubleList) AddFirst(node *Node) {
// 	node.next = this.head.next
// 	node.prev = this.head
// 	this.head.next.prev = node
// 	this.head.next = node
// }
//
// func (this *DoubleList) Remove(node *Node) {
// 	node.prev.next = node.next
// 	node.next.prev = node.prev
// 	node.next = nil
// 	node.prev = nil
// }
//
// func (this *DoubleList) RemoveLast() *Node {
// 	if this.IsEmpty() {
// 		return nil
// 	}
// 	last := this.tail.prev
// 	this.Remove(last)
// 	return last
// }
//
// func (this *DoubleList) IsEmpty() bool {
// 	return this.head.next == this.tail
// }

// type innerValue struct {
// 	value int
// 	t     int64
// 	freq  int
// }
//
// type LFUCache struct {
// 	sync.RWMutex
// 	cap   int
// 	cache map[int]innerValue
// }
//
// func Constructor(capacity int) LFUCache {
// 	return LFUCache{cap: capacity, cache: make(map[int]innerValue, capacity)}
// }
//
// func (this *LFUCache) Get(key int) int {
// 	this.Lock()
// 	defer this.Unlock()
// 	if this.cap == 0 {
// 		return -1
// 	}
// 	in, ok := this.cache[key]
// 	if ok {
// 		in.t = time.Now().UnixNano()
// 		in.freq++
// 		this.cache[key] = in
// 		return in.value
// 	}
// 	return -1
// }
//
// func (this *LFUCache) Put(key int, value int) {
// 	this.Lock()
// 	defer this.Unlock()
// 	inner, ok := this.cache[key]
//
// 	if ok {
// 		this.cache[key] = innerValue{value: value, t: time.Now().UnixNano(), freq: inner.freq + 1}
// 		return
// 	} else {
// 		usedLen := len(this.cache)
// 		if usedLen < this.cap {
// 			this.cache[key] = innerValue{value, time.Now().UnixNano(), 1}
// 			return
// 		}
// 		var (
// 			oldestKey  int
// 			oldestTime int64
// 			oldestFreq int
// 		)
// 		for k, v := range this.cache {
// 			if (oldestTime == 0 || v.freq < oldestFreq) || (v.freq == oldestFreq && v.t < oldestTime) {
// 				oldestKey = k
// 				oldestTime = v.t
// 				oldestFreq = v.freq
// 			}
// 		}
// 		delete(this.cache, oldestKey)
// 		this.cache[key] = innerValue{value, time.Now().UnixNano(), 1}
// 	}
// }
