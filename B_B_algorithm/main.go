package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Get(3)
	cache.Get(4)
}

// 146. LRU 缓存机制
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

// 198. 打家劫舍
// 213. 打家劫舍 II 首尾相连
func rob(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		return max(_rob(nums[:n-1]), _rob(nums[1:]))
		// 198 return _rob(nums)
	}
}
func _rob(nums []int) int {
	index, maxNum := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		index, maxNum = maxNum, max(index+v, maxNum)
	}
	return maxNum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1019. 链表中的下一个更大节点
func nextLargerNodes(head *ListNode) []int {
	if head.Next == nil {
		return []int{0}
	}
	res := make([]int, 0)
	var left, right int
	var same bool
	sameNext := &ListNode{}
	point := head.Val
	curr := head
	for head.Next != nil {
		same = false
		for curr != nil && curr.Val <= point {
			if curr.Val == point {
				same = true
				sameNext.Next = curr.Next
			}
			curr = curr.Next
			right++
		}

		if curr == nil {
			head = head.Next
			curr = head
			point = head.Val
			res = append(res, 0)

		} else {
			// 有相等 需要一个个走
			if same && sameNext.Next != curr.Next {
				res = append(res, curr.Val)
				head = head.Next
				curr = head
			} else {
				head = curr
				for i := 0; i < right-left; i++ {
					res = append(res, curr.Val)
				}
			}

			point = curr.Val
		}

		left = right
	}
	res = append(res, 0)
	return res
}

// 面试题 04.03. 特定深度节点链表
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return make([]*ListNode, 0)
	}
	resList := make([]*ListNode, 0)
	preList := make([]*TreeNode, 0)
	preList = append(preList, tree)
	for len(preList) > 0 {
		node := preList[0]
		currNode := &ListNode{
			Val: node.Val,
		}
		head := currNode
		nextList := make([]*TreeNode, 0)

		if node.Left != nil {
			nextList = append(nextList, node.Left)
		}
		if node.Right != nil {
			nextList = append(nextList, node.Right)
		}
		for i := 1; i < len(preList); i++ {
			if preList[i].Left != nil {
				nextList = append(nextList, preList[i].Left)
			}
			if preList[i].Right != nil {
				nextList = append(nextList, preList[i].Right)
			}
			currNode.Next = &ListNode{
				Val: preList[i].Val,
			}
			currNode = currNode.Next
		}
		resList = append(resList, head)
		preList = nextList
	}
	return resList
}

// 39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	tmp := make([]int, 0)
	tmp = append(tmp, candidates[0])
	for i := 1; i < len(candidates); i++ {
		if candidates[i] == candidates[i-1] {
			continue
		}
		tmp = append(tmp, candidates[i])
	}
	sums := make([]int, 0)
	res := make([][]int, 0)

	var f func(index, sum int, sums, nums []int)
	f = func(index, sum int, sums, nums []int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, sums)
			return
		}
		for i := index; i < len(nums); i++ {
			sums = append(sums, nums[i])
			sum += nums[i]
			if sum == target {
				one := make([]int, 0)
				for _, val := range sums {
					one = append(one, val)
				}
				res = append(res, one)
				return
			}
			f(i, sum, sums, nums)
			sums = sums[0 : len(sums)-1]
			sum -= nums[i]
		}
	}
	f(0, 0, sums, tmp)
	return res
}

// 779. 第K个语法符号
// 在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。
// 给定行数 N 和序数 K，返回第 N 行中第 K个字符。（K从1开始）
func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}
	if K == 2 {
		return K - 1
	}
	// 0 0 1 0 1 0 1
	var flag = 1
	res := []int{0, 0, 1}
	for K >= 2 {
		high := getHighBit(K)
		K -= high
		flag = -flag
	}
	if flag < 0 {
		return 1 ^ res[K]
	}
	return res[K]
}
func getHighBit(x int) int {
	tmp := x
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x | (x >> 32)
	x = x | (x >> 16)
	x = (x + 1) >> 1
	if tmp > x {
		return x
	}
	return x >> 1
}

// 779 递归写法
func kthGrammar_1(N int, K int) int {
	if N == 1 {
		return 0
	}
	// 对比中间位置
	if K <= (1 << (N - 2)) {
		return kthGrammar_1(N-1, K)
	}
	return kthGrammar_1(N-1, K-(1<<(N-2))) ^ 1
}
func ConvertToIntSlice(s string) []int {
	res := make([]int, 0)
	s = strings.ReplaceAll(s, " ", "")
	s = s[1 : len(s)-1]
	list := strings.Split(s, ",")
	for _, elem := range list {
		num, _ := strconv.Atoi(elem)
		res = append(res, num)
	}
	return res
}
func ConvertToIntSlice2(s string) [][]int {
	res := make([][]int, 0)
	s = strings.ReplaceAll(s, " ", "")
	s = s[2 : len(s)-2]
	list := strings.Split(s, "],[")
	for _, elem := range list {
		slice1 := ConvertToIntSlice("[" + elem + "]")
		res = append(res, slice1)
	}
	return res
}
func ConvertToStrSlice(s string) []string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, " \"", "")
	s = s[1 : len(s)-1]
	list := strings.Split(s, ",")
	return list
}

// 82. 删除排序链表中的重复元素 II
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		next    *ListNode
		pre     = &ListNode{Val: head.Val - 1}
		tmp     = pre
		resHead = pre
	)
	for head.Next != nil {
		next = head.Next
		if tmp.Val != head.Val && head.Val != next.Val {
			pre.Next = &ListNode{Val: head.Val}
			pre = pre.Next
		}
		tmp = head
		head = head.Next
	}
	if tmp.Val != head.Val {
		pre.Next = &ListNode{Val: head.Val}
		pre = pre.Next
	}
	return resHead.Next
}

// 用拼接的方式
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		next    *ListNode
		pre     = &ListNode{Val: head.Val - 1}
		tmp     = pre
		resHead = pre
	)
	for head.Next != nil {
		next = head.Next
		if tmp.Val != head.Val && head.Val != next.Val {
			pre.Next = head
			pre = pre.Next
		}
		tmp = head
		head = head.Next
	}
	if tmp.Val != head.Val {
		pre.Next = head
	} else {
		pre.Next = nil
	}
	return resHead.Next
}

// 官方
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 好设计
func deleteDuplicates3(head *ListNode) *ListNode {
	flag := false
	if head == nil {
		return nil
	}
	res := &ListNode{}
	p := res
	// 依次遍历
	// 如果当前节点与后面节点的值相同，那么当前就点就是重复节点，flag为true
	// 不过不同分情况讨论
	// 1.如果此时flag为true，那么表明当前节点与它前一个节点的值一样，同样为重复节点
	// 2.如果此时flag为false，那么表明当前节点只出现了一次，是要保留的节点
	// 不管是情况1还是情况2，做完都需要把flag变为false，从下一个节点重新开始比较
	// 因为我们用的是 head.Next!=nil 所以要判断最后一个节点是否是重复节点，判断条件同上
	// 最后因为我们是直接把不重复的节点接到答案中，所以避免最后接入的节点带有其他值，加一个 p.Next=nil 条件
	for head.Next != nil {
		if head.Val == head.Next.Val {
			flag = true
		} else {
			if flag {
				flag = !flag
			} else {
				p.Next = head
				p = p.Next
			}
		}
		head = head.Next
	}
	if !flag {
		p.Next = head
		p = p.Next
	}
	p.Next = nil
	return res.Next
}
