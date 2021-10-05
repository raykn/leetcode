package main

import (
	"fmt"
	"sort"
)

func main() {
	var res = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(res))
}

// TODO 25. K 个一组翻转链表 ———— 困难 自己
func reverseKGroup(head *ListNode, k int) *ListNode {
	var resHead, pre *ListNode
	var flag = true
	for {
		cur, tmpHead := head, head
		for i := 0; i < k; i++ {
			if cur == nil {
				return resHead
			}
			cur = cur.Next
		}

		// head 进去后被安在链表最尾
		reverseHead := reverse(tmpHead, cur, k)
		if flag {
			resHead = reverseHead
			flag = false
		} else {
			pre.Next = reverseHead
		}

		pre = head
		head = cur
	}
}
func reverse(cur, pre *ListNode, num int) *ListNode {
	for i := 0; i < num; i++ {
		if cur == nil {
			return pre
		}
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

// TODO 25. K 个一组翻转链表 ———— 困难 100% nb
func reverseKGroup_(head *ListNode, k int) *ListNode {
	fake := &ListNode{Next: head}
	pre := fake
	now, next := head, head
	for {
		for i := 0; i < k; i++ {
			if next == nil {
				return fake.Next
			}
			next = next.Next
		}

		// 之所以是 i = 1，因为每次将后面插入到前面来，动 n - 1 个node
		// 12345 -> 21345 -> 32145 -> 43215 -> 54321
		for i := 1; i < k; i++ {
			// 因为 pre 不作为真实节点，所以全部后移 .Next
			// 存 now.Next
			// now.Next = now.Next.Next 逆转
			// 把之前 pre 接上
			pre.Next, now.Next, now.Next.Next = now.Next, now.Next.Next, pre.Next
			print(pre.Next)
			// 正常反转链表是拿 		pre -> now -> now.Next 交换
			// 而这里虚拟头节点是拿 	pre.Next -> now.Next -> now.Next.Next 交换
			// tmp := now.Next.Next
			// now.Next.Next = pre.Next
			// pre.Next = now.Next
			// now.Next = tmp

		}
		pre, now = now, next
	}
}

// TODO 206. 反转链表 ———— 骚
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}

// TODO 206. 反转链表 ———— 正常写法
func reverseList_(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

// TODO 206. 反转链表 ———— 递归
func reverseList__(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// TODO 3. 无重复字符的最长子串 ———— 自己
func lengthOfLongestSubstring(s string) int {
	exist := make(map[byte]int)
	var right, max int
	n := len(s)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(exist, s[i-1])
		}
		for right < n && exist[s[right]] == 0 {
			exist[s[right]]++
			right++
		}
		max = maxNum(max, right-i)
	}
	return max
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// TODO 3. 无重复字符的最长子串 ———— 100%
func lengthOfLongestSubstring_(s string) int {
	ans := 0
	exist := [128]bool{}
	for start, j := 0, 0; j < len(s); j++ {
		// 发现重复 清空并更新 start
		for exist[s[j]] {
			exist[s[start]] = false
			start++
		}
		exist[s[j]] = true
		if j-start+1 > ans {
			ans = j - start + 1
		}
	}
	return ans
}

// TODO 215. 数组中的第K个最大元素 ———— 高频常考 快排优化
func findKthLargest(nums []int, k int) int {
	if k == 0 || len(nums) == 0 {
		return 0
	}
	// 快排分左右
	partition := func(nums []int, left, right int) int {
		point := nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			if nums[i] < point {
				j++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		nums[j], nums[left] = nums[left], nums[j]
		return j
	}
	n := len(nums)
	left, right, target := 0, n-1, n-k
	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

// TODO 215. 数组中的第K个最大元素 ———— 高频常考 堆排
func findKthLargest_(nums []int, k int) int {
	sLen := len(nums)
	// 建堆
	buildHeap(nums)
	for i := sLen - 1; i > sLen-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		// 注意这里不用调整整个堆 i之前就行
		moveHeap(nums, 0, i)
	}
	return nums[0]
}

func buildHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		moveHeap(nums, i, len(nums))
	}
}

// 调整堆节点
func moveHeap(nums []int, k, sLen int) {
	if k >= sLen {
		return
	}
	idx := k
	child1, child2 := k*2+1, k*2+2
	if child1 < sLen && nums[child1] > nums[idx] {
		idx = child1
	}
	if child2 < sLen && nums[child2] > nums[idx] {
		idx = child2
	}
	if idx != k {
		nums[idx], nums[k] = nums[k], nums[idx]
		moveHeap(nums, idx, sLen)
	}
}

// TODO 103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	flag := true
	res := make([][]int, 0)
	var queue = []*TreeNode{root}
	for len(queue) != 0 {
		flag = !flag
		var tmpQue []*TreeNode
		var tmpRes []int
		for i := 0; i < len(queue); i++ {
			tmpRes = append(tmpRes, queue[i].Val)
			if queue[i].Left != nil {
				tmpQue = append(tmpQue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmpQue = append(tmpQue, queue[i].Right)
			}
		}
		if flag {
			for i := 0; i < len(tmpRes)/2; i++ {
				tmpRes[i], tmpRes[len(tmpRes)-i-1] = tmpRes[len(tmpRes)-i-1], tmpRes[i]
			}
		}
		res = append(res, tmpRes)
		queue = tmpQue
	}
	return res
}

// TODO 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		p1 := nums[i]
		if p1 > 0 {
			break
		}
		if i > 0 && p1 == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			p2, p3 := nums[left], nums[right]
			if p1+p2+p3 == 0 {
				res = append(res, []int{p1, p2, p3})
				for left < right && nums[left] == p2 {
					left++
				}
				for left < right && nums[right] == p3 {
					right--
				}

			} else if p1+p2+p3 > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

// TODO 买卖股票一系列

// TODO 33. 搜索旋转排序数组 ———— 自己
func search(nums []int, target int) int {
	index := len(nums) - 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			index = i - 1
			break
		}
	}

	if target > nums[index] {
		return -1
	}

	f := func(nums []int, l, r int) int {
		if l < 0 || r >= len(nums) {
			return -1
		}
		for l <= r {
			mid := l + (r-l)/2
			if target == nums[mid] {
				return mid
			} else if target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return -1
	}

	if res := f(nums, 0, index); res != -1 {
		return res
	}
	return f(nums, index+1, len(nums)-1)
}

// TODO 33. 搜索旋转排序数组 ———— 100%
func search_(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			// 没到旋转
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// TODO 236. 二叉树的最近公共祖先 ———— 自己
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	v := make(map[int]bool)
	parent := make(map[int]*TreeNode)

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)
	for p != nil {
		v[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if v[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

// TODO 236. 二叉树的最近公共祖先 ———— 100%
func lowestCommonAncestor_(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor_(root.Left, p, q)
	right := lowestCommonAncestor_(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		} else {
			return left
		}
	}
	return right
}

// TODO 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	cur1, cur2 := headA, headB
	for {
		if cur1 == cur2 {
			return cur1
		}
		if cur1 == nil {
			cur1 = headB
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = headA
		} else {
			cur2 = cur2.Next
		}
	}
}

// TODO 143. 重排链表 ———— 自己 100%
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur1, cur2 := head, head
	for cur2 != nil && cur2.Next != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next.Next
	}
	headB := cur1.Next
	cur1.Next = nil

	var pre *ListNode
	for headB != nil {
		pre, headB, headB.Next = headB, headB.Next, pre
	}
	headB = pre

	for head != nil && headB != nil {
		tmpA, tmpB := head.Next, headB.Next
		head.Next = headB
		headB.Next = tmpA
		head, headB = tmpA, tmpB
	}
}

// TODO 31. 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	sort.Ints(nums[i+1:])
	// reverseNums(nums[i+1:])
}

func reverseNums(nums []int) {
	for i, end := 0, len(nums); i < end/2; i++ {
		nums[i], nums[end-1-i] = nums[end-1-i], nums[i]
	}
}

// TODO 23. 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	head := lists[0]
	for i := 1; i < len(lists); i++ {
		head = mergeList(head, lists[i])
	}
	return head
}

func mergeList(p, q *ListNode) *ListNode {
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}

	head := &ListNode{}
	cur := head
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}
	return head.Next
}

// TODO 23. 合并K个升序链表 ———— 100%
func mergeKLists_(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	l := mergeKLists(lists[:len(lists)/2])
	r := mergeKLists(lists[len(lists)/2:])
	res := &ListNode{}
	temp := res
	for l != nil && r != nil {
		if l.Val < r.Val {
			temp.Next = l
			l = l.Next
		} else {
			temp.Next = r
			r = r.Next
		}
		temp = temp.Next
	}
	if l != nil {
		temp.Next = l
	}
	if r != nil {
		temp.Next = r
	}
	return res.Next
}

// TODO 42. 接雨水

// TODO 200. 岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var dfs func(g [][]byte, i, j int)
	dfs = func(g [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(g) || j >= len(g[0]) || g[i][j] == '0' {
			return
		}
		g[i][j] = '0'
		dfs(g, i+1, j)
		dfs(g, i-1, j)
		dfs(g, i, j+1)
		dfs(g, i, j-1)
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

// TODO 54. 螺旋矩阵 ———— 自己
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	row, col, total := len(matrix), len(matrix[0]), len(matrix)*len(matrix[0])
	var res []int
	v := make(map[int]bool)
	i, j := 0, 0
	for {
		for j < col && !v[i*col+j] {
			res = append(res, matrix[i][j])
			v[i*col+j] = true
			j++
		}
		j--
		i++

		for i < row && !v[i*col+j] {
			res = append(res, matrix[i][j])
			v[i*col+j] = true
			i++
		}
		i--
		j--

		for j >= 0 && !v[i*col+j] {
			res = append(res, matrix[i][j])
			v[i*col+j] = true
			j--
		}
		j++
		i--

		for i >= 0 && !v[i*col+j] {
			res = append(res, matrix[i][j])
			v[i*col+j] = true
			i--
		}
		i++
		j++

		if len(res) == total {
			return res
		}
	}
}

// TODO 54. 螺旋矩阵 ———— 100% 更丝滑
func spiralOrder_(matrix [][]int) []int {
	var res []int
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for up <= down && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if up > down || right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--

		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

// TODO 92. 反转链表 II ———— 自己
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Next: head}

	l := right - left + 1
	cur := newHead
	left -= 1
	for left != 0 && cur != nil {
		left--
		cur = cur.Next
	}

	if cur == nil {
		return newHead.Next
	}

	tmpHead, tmpTail := cur, cur.Next
	var pre *ListNode
	cur = cur.Next
	for cur != nil && l > 0 {
		l--
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	tmpHead.Next = pre
	tmpTail.Next = cur
	return newHead.Next
}

// TODO 92. 反转链表 II ———— 100%
func reverseBetween_(head *ListNode, left int, right int) *ListNode {
	if right-left == 0 {
		return head
	}
	newHead := &ListNode{Next: head}
	pre := newHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		pre.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, pre.Next
	}
	return newHead.Next
}

// TODO 199. 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := []int{root.Val}
	for len(queue) != 0 {
		var tmpQ []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				tmpQ = append(tmpQ, node.Left)
			}
			if node.Right != nil {
				tmpQ = append(tmpQ, node.Right)
			}
		}
		if len(tmpQ) > 0 {
			res = append(res, tmpQ[len(tmpQ)-1].Val)
		}
		queue = tmpQ
	}
	return res
}

// TODO 300. 最长递增子序列 ———— 正常动态
func lengthOfLIS(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	var res int
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// TODO 300. 最长递增子序列 ———— 动态 + 贪心
func lengthOfLIS_(nums []int) int {
	tails := make([]int, len(nums))
	var res = 0
	for _, val := range nums {
		left, right := 0, res
		for left < right {
			mid := (left + right) / 2
			if val > tails[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		tails[left] = val
		fmt.Println(tails)
		if res == right {
			res++
		}
	}
	return res
}

// TODO 300. 最长递增子序列 ———— 100%
func lengthOfLIS__(nums []int) int {
	var dp []int
	for _, num := range nums {
		if len(dp) == 0 || num > dp[len(dp)-1] {
			dp = append(dp, num)
		} else {
			// 贪心 尝试重建最长序列
			l, r, index := 0, len(dp)-1, 0
			for l <= r {
				mid := (l + r) >> 1
				// 这个判断保证至少会走一次这里
				// 因为全部 num > dp[mid]
				// 那应该是上面的 num > dp[len(dp) - 1]
				if num <= dp[mid] {
					index = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[index] = num
		}
	}
	return len(dp)
}

// TODO 42. 接雨水　————　动态规划
func trap(height []int) int {
	n := len(height)
	leftDP := make([]int, n)
	rightDP := make([]int, n)
	res := 0

	leftDP[0] = height[0]
	for i := 1; i < n; i++ {
		leftDP[i] = max(leftDP[i-1], height[i])
	}

	rightDP[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightDP[i] = max(rightDP[i+1], height[i])
	}

	fmt.Println(leftDP)
	fmt.Println(rightDP)

	for i, h := range height {
		res += min(leftDP[i], rightDP[i]) - h
	}
	return res
}

// TODO 42. 接雨水　————　单调栈
func trap_(height []int) int {
	var stack []int
	var ans int
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIndex := stack[len(stack)-1]
			curWidth := i - leftIndex - 1
			curHeight := min(height[leftIndex], h) - height[topIndex]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return ans
}
