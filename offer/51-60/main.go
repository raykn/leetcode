package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// runtime.GOMAXPROCS(1)
	// for i := 0; i < 260; i++ {
	// 	i := i
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }
	// select {}
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}

// TODO 51-数组中的逆序对
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	var tmp []int
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			cnt += mid - i + 1
			tmp = append(tmp, nums[j])
			j++
		} else {
			tmp = append(tmp, nums[i])
			i++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}

// TODO 52-两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	exist := make(map[*ListNode]struct{})
	for headA != nil {
		exist[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := exist[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode_(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

// TODO 53-I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	var index = -1
	// <= 边界也涉及

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			index = mid
			break
		}
		// 两边都要-
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if index < 0 {
		return 0
	}
	l, r := index, index
	for l >= 0 && nums[l] == target {
		l--
	}
	l++

	for r < len(nums) && nums[r] == target {
		r++
	}
	r--

	return r - l + 1
}

// TODO 53 - II. 0～n-1中缺失的数字
func missingNumber(nums []int) int {
	i := 0
	for _, v := range nums {
		if v != i {
			return i
		}
		i++
	}
	if i == len(nums) {
		return i
	}
	i++
	return i
}

// TODO 53 - II. 0～n-1中缺失的数字 —— 二分法
func missingNumber_(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		// 左边的排序没问题的
		// 说明缺失的数字在右边
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// TODO 54. 二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var list []int
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		list = append(list, r.Val)
		dfs(r.Right)
	}
	dfs(root)
	return list[len(list)-k]
}

// TODO 54. 二叉搜索树的第k大节点 —— 倒数 直接用后续遍历 k--
func kthLargest_(root *TreeNode, k int) int {
	var res int
	var f func(r *TreeNode) int
	f = func(r *TreeNode) int {
		if r == nil {
			return res
		}
		f(r.Right)
		if k == 0 {
			return res
		}
		k--
		if k == 0 {
			res = r.Val
		}
		f(r.Left)
		return res
	}
	f(root)
	return res
}

// TODO 55 - I. 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO 55 - II. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	return !(deep(root) == -1)
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := deep(root.Left)
	r := deep(root.Right)
	// 是-1提前返回 fast fail
	if l == -1 || r == -1 {
		return -1
	}
	// 下面可以封装 math.Abs
	if l-r > 1 || r-l > 1 {
		return -1
	}
	// 下面可以封装 max(x, y int)
	if l > r {
		return l + 1
	}
	return r + 1
}

// TODO 56 - I. 数组中数字出现的次数
func singleNumbers(nums []int) []int {
	// 总的异或结果
	res := 0
	for _, v := range nums {
		res ^= v
	}

	// 用于隔开分组异或
	// 某个位置为1 一定是结果集其中一个的特殊一位
	// 因为如果两个都有这一位 异或后就应该是 0
	pos := 1
	for {
		if res&pos != 0 {
			break
		}
		pos <<= 1
	}

	var x, y int
	for _, v := range nums {
		if v&pos != 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	return []int{x, y}
}

// TODO 57. 和为s的两个数字
func twoSum(nums []int, target int) []int {
	var res []int
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{nums[l], nums[r]}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return res
}

// TODO 57. 和为s的两个数字 —— 100%
func twoSum_(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for nums[l]+nums[r] != target {
		if nums[l]+nums[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{nums[l], nums[r]}
}

// TODO 57 - II. 和为s的连续正数序列
func findContinuousSequence(target int) [][]int {
	var res [][]int
	if target < 3 {
		return res
	}
	l, r, m := 1, 2, (target+1)/2
	sum := l + r
	for l < m {
		if target == sum {
			var tmp []int
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			sum -= l
			l++
		} else if target > sum {
			r++
			sum += r
		} else {
			sum -= l
			l++
		}
	}
	return res
}

// TODO 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	words := regexp.MustCompile(`\s+`).Split(strings.Trim(s, " "), -1)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseWords_(s string) string {
	// main process
	sb := trimSpaces(s)
	reverse(sb, 0, len(sb)-1)
	reverseEachWord(sb)
	return string(sb)
}

// 1. 双指针去掉首尾多余空格，同时转化为字节数组
func trimSpaces(s string) []byte {
	left, right := 0, len(s)-1
	// 去掉两端空格
	for left <= right && s[left] == ' ' {
		left++
	}
	for left <= right && s[right] == ' ' {
		right--
	}
	// 去掉中间多余空格
	sb := make([]byte, 0, right-left+1)
	for left <= right {
		if s[left] != ' ' {
			// 不为空格则放入sb
			sb = append(sb, s[left])
		} else if sb[len(sb)-1] != ' ' {
			// sb最后一个字符不为空格则放入，此处保证了单词直接只保留1个空格
			sb = append(sb, s[left])
		}
		left++
	}
	return sb
}

// 2. 翻转区间 [left, right] 左闭右闭区间
func reverse(sb []byte, left, right int) {
	for ; left < right; left, right = left+1, right-1 {
		sb[left], sb[right] = sb[right], sb[left]
	}
}

// 3. 翻转每个单词
func reverseEachWord(sb []byte) {
	// start,end：单词的起始位置，n：字符串长度
	start, end, n := 0, 0, len(sb)
	for start < n {
		for end < n && sb[end] != ' ' {
			end++
		}
		// 此时 sb[start, end) 是一个单词，翻转之，注意：此时end指向的是空格的位置
		reverse(sb, start, end-1)
		// 更新 start,end 去找下一个单词的起始位置
		start, end = end+1, end+1
	}
}

// TODO 58 - I. 翻转单词顺序
func reverseLeftWords(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[n:] + s[:n]
}

// TODO 59 - I. 滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var q []int
	for i := 0; i < k; i++ {
		for len(q) != 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, nums[i])
	}

	var res []int
	res = append(res, q[0])
	for i := k; i < len(nums); i++ {
		if nums[i-k] == q[0] {
			q = q[1:]
		}
		for len(q) != 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, nums[i])
		res = append(res, q[0])
	}
	return res
}

// TODO
