package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(permutation_("aab"))
}

// TODO 31-栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 || len(popped) == 0 {
		return true
	}
	var stack []int
	i, j := 0, 0
	stack = append(stack, pushed[i])
	i++
	for j < len(popped) {
		if len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		} else {
			if i == len(pushed) {
				break
			}
			stack = append(stack, pushed[i])
			i++
		}
	}
	return len(stack) == 0 && j == len(popped)
}

// TODO 31-栈的压入、弹出序列 100%
func validateStackSequences_(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}

// TODO 32-从上到下打印二叉树
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var q []*TreeNode
	var res []int
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return res
}

// TODO 32-从上到下打印二叉树 II
func levelOrder_(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var tmpQ []*TreeNode
		var tmpRes []int
		for _, v := range q {
			tmpRes = append(tmpRes, v.Val)
			if v.Left != nil {
				tmpQ = append(tmpQ, v.Left)
			}
			if v.Right != nil {
				tmpQ = append(tmpQ, v.Right)
			}
		}
		q = tmpQ
		res = append(res, tmpRes)
	}
	return res
}

// TODO 32-从上到下打印二叉树 III
func levelOrder__(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var q []*TreeNode
	q = append(q, root)
	flag := false
	for len(q) > 0 {
		var tmpQ []*TreeNode
		var tmpRes []int
		for _, v := range q {
			tmpRes = append(tmpRes, v.Val)
			if v.Left != nil {
				tmpQ = append(tmpQ, v.Left)
			}
			if v.Right != nil {
				tmpQ = append(tmpQ, v.Right)
			}
		}
		q = tmpQ
		var tmpArray []int
		if flag {
			for i := len(tmpRes) - 1; i >= 0; i-- {
				tmpArray = append(tmpArray, tmpRes[i])
			}

		} else {
			tmpArray = tmpRes
		}
		res = append(res, tmpArray)
		flag = !flag
	}
	return res
}

// TODO 33-二叉搜索树的后序遍历序列
func verifyPostorder(postorder []int) bool {
	return checkTree(postorder, 0, len(postorder)-1)
}

func checkTree(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	index := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && checkTree(postorder, i, index-1) && checkTree(postorder, index, j-1)
}

// TODO 33-二叉搜索树的后序遍历序列 单调递增栈
func verifyPostorder_(postorder []int) bool {
	var stack []int
	root := math.MaxInt32
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}

// TODO 34-二叉树中和为某一值的路径
func pathSum(root *TreeNode, target int) [][]int {
	var resList [][]int
	if root == nil {
		return resList
	}

	path := []int{root.Val}
	var dfs func(root *TreeNode, sum int)
	dfs = func(r *TreeNode, sum int) {
		if sum == target && r.Left == nil && r.Right == nil {
			res := make([]int, len(path))
			copy(res, path)
			resList = append(resList, res)
		}
		if r.Left != nil {
			path = append(path, r.Left.Val)
			dfs(r.Left, sum+r.Left.Val)
			path = path[:len(path)-1]
		}
		if r.Right != nil {
			path = append(path, r.Right.Val)
			dfs(r.Right, sum+r.Right.Val)
			path = path[:len(path)-1]
		}
	}
	dfs(root, root.Val)
	return resList
}

// TODO 35-复杂链表的复制
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{}
	cur := newHead
	oldCur := head
	for oldCur != nil {
		node := &Node{
			Val:  oldCur.Val,
			Next: oldCur.Next,
		}
		cur.Next = node
		cur = cur.Next
		oldCur = oldCur.Next
	}
	oldCur = head
	newCur := newHead.Next
	for oldCur != nil {
		if oldCur.Random != nil {
			step := 0
			tmp := head
			for tmp != nil {
				if tmp == oldCur.Random {
					break
				}
				step++
				tmp = tmp.Next
			}

			tmp = newHead.Next
			for step > 0 {
				tmp = tmp.Next
				step--
			}
			newCur.Random = tmp
		}
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	return newHead.Next
}

// TODO 35-复杂链表的复制 人才
func copyRandomList_(head *Node) *Node {
	ptr, nodeHash := head, make(map[*Node]*Node)
	for ; ptr != nil; ptr = ptr.Next {
		nodeHash[ptr] = &Node{ptr.Val, nil, nil}
	}
	for ptr = head; ptr != nil; ptr = ptr.Next {
		nodeHash[ptr].Next, nodeHash[ptr].Random = nodeHash[ptr.Next], nodeHash[ptr.Random]
	}
	return nodeHash[head]
}

// TODO 36-二叉搜索树与双向链表
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var pre *TreeNode
	var head *TreeNode
	var convert func(root *TreeNode)
	convert = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		convert(cur.Left)
		if pre == nil {
			head = cur
		} else {
			pre.Right = cur
			cur.Left = pre
		}
		pre = cur
		convert(cur.Right)
	}
	convert(root)
	pre.Right = head
	head.Left = pre
	return head
}

// TODO 38-字符串的排列
func permutation(s string) []string {
	var (
		res    []string
		sByte  = []byte(s)
		dfs    func(s []byte, str []byte)
		vMap   = make(map[int]bool)
		resMap = make(map[string]bool)
	)
	dfs = func(s []byte, str []byte) {
		if len(s) == len(str) {
			tmp := string(str)
			if !resMap[tmp] {
				res = append(res, tmp)
				resMap[tmp] = true
			}
			return
		}
		for i, v := range s {
			if vMap[i] {
				continue
			}
			vMap[i] = true
			dfs(s, append(str, v))
			vMap[i] = false
		}
	}
	dfs(sByte, []byte{})
	return res
}

// TODO 38-字符串的排列 100%
func permutation_(s string) []string {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	var (
		n    = len(s)
		res  []string
		temp = make([]byte, 0, n)
		vis  = make([]bool, n)
		dfs  func(int)
	)
	dfs = func(x int) {
		if x == n {
			res = append(res, string(temp))
			return
		}
		for i, j := range vis {
			if j || (i > 0 && !vis[i-1] && t[i-1] == t[i]) {
				continue
			}
			temp = append(temp, t[i])
			vis[i] = true
			dfs(x + 1)
			temp = temp[:len(temp)-1]
			vis[i] = false
		}
	}
	dfs(0)
	return res
}

// TODO 38-字符串的排列 还没看 应该比100%还快
func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(nums []byte) bool {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for j >= 0 && nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	return true
}

func permutation__(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for {
		ans = append(ans, string(t))
		if !nextPermutation(t) {
			break
		}
	}
	return
}

// TODO 39-数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	p, num := 0, 0
	for _, v := range nums {
		if num > 0 && p != v {
			num--
		} else {
			num++
			p = v
		}
	}
	return p
}

// TODO 40－最小的k个数 堆重要学习下
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	// 方法一 基于快排的快速选择
	// return quicksearch(arr, 0, len(arr)-1, k)

	// 方法二 快排 然后取前k个
	// quicksort(arr, 0, len(arr)-1)
	// return arr[:k]

	// 方法三  建堆，大根堆
	d := &heapInt{}
	for _, v := range arr {
		if d.Len() < k {
			heap.Push(d, v)
		} else {
			if d.Peek() > v {
				heap.Pop(d)
				heap.Push(d, v)
			}
		}
	}
	return *d

}

func partition(nums []int, i, j int) int {
	l, m, r := i, i, j
	for l < r {
		for l < r && nums[r] >= nums[m] {
			r--
		}
		for l < r && nums[l] <= nums[m] {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[m], nums[l] = nums[l], nums[m]
	return l
}

func quicksearch(nums []int, i, j, k int) []int {
	t := partition(nums, i, j)
	if t == k-1 {
		return nums[:k]
	}
	if t < k-1 {
		return quicksearch(nums, t+1, j, k)
	}
	return quicksearch(nums, i, t-1, k)
}

func quicksort(nums []int, i, j int) {
	if i >= j {
		return
	}
	m := partition(nums, i, j)
	quicksort(nums, i, m-1)
	quicksort(nums, m+1, j)
}

type heapInt []int

// Less 小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0]
}
