package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder__(matrix))
}

// TODO 21-调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	index := 0
	for index < len(nums) {
		if nums[index]&1 == 0 {
			break
		}
		index++
	}
	if index == len(nums) {
		return nums
	}

	p := index
	for index < len(nums) {
		if nums[index]&1 == 1 {
			nums[p], nums[index] = nums[index], nums[p]
			p++
		}
		index++
	}
	return nums
}

// TODO 21-调整数组顺序使奇数位于偶数前面 双指针
func exchange_(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]&1 == 1 {
			l++
		}
		for l < r && nums[r]&1 == 0 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]

	}
	return nums
}

// TODO 22-链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	cap := 0
	cur := head
	for cur != nil {
		cap++
		cur = cur.Next
	}

	index := cap - k
	if index <= 0 {
		return head
	}
	cur = head
	for cur != nil {
		index--
		if index < 0 {
			return cur
		}
		cur = cur.Next
	}
	return &ListNode{}
}

// TODO 22-链表中倒数第k个节点 快慢指针
func getKthFromEnd_(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for i := 1; i <= k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// TODO 24-反转链表 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		// next == nil 就返回，因为最后一个节点是新头节点
		return head
	}
	root := reverseList(head.Next)
	if head.Next != nil {
		head.Next.Next = head // 插入尾节点
		head.Next = nil       // 尾节点后面是 nil
	}
	return root
}

// TODO 24-反转链表 递归
func reverseList_(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, tmp *ListNode
	cur := head
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// TODO 24-反转链表 骚操作
func reverseList__(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

// TODO 25-合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	root := &ListNode{}
	cur := root
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			cur.Next = p2
			p2 = p2.Next
		} else {
			cur.Next = p1
			p1 = p1.Next
		}
		cur = cur.Next
	}
	if p2 != nil {
		cur.Next = p2
	}
	if p1 != nil {
		cur.Next = p1
	}
	return root.Next
}

// TODO 26-树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	var yes bool
	var visit func(r *TreeNode)
	var isSame func(a, b *TreeNode) bool

	isSame = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if b == nil {
			return true
		}
		if a == nil && b != nil {
			return false
		}
		return a.Val == b.Val && isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
	}

	visit = func(root *TreeNode) {
		if root == nil {
			return
		}
		if isSame(root, B) {
			yes = true
			return
		}
		visit(root.Left)
		visit(root.Right)
	}

	visit(A)
	return yes
}

// TODO 26-树的子结构
func isSubStructure_(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (visit(A, B) || isSubStructure_(A.Left, B) || isSubStructure_(A.Right, B))
}

func visit(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return visit(A.Left, B.Left) && visit(A.Right, B.Right)
}

// TODO 27-二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

// TODO 28-对称的二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q []*TreeNode
	q = append(q, root.Left, root.Right)
	for len(q) > 0 {
		n1 := q[0]
		n2 := q[1]
		q = q[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		q = append(q, n1.Left)
		q = append(q, n2.Right)
		q = append(q, n1.Right)
		q = append(q, n2.Left)
	}
	return true
}

// TODO 28-对称的二叉树 递归解法
func isSymmetric_(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, root.Right)
}

func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return isSameTree(a.Left, b.Right) && isSameTree(a.Right, b.Left)
}

// TODO 29-顺时针打印矩阵 自己的low写法
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	i, j := len(matrix), len(matrix[0])
	var res []int
	var time = i / 2
	if i&1 == 1 {
		time++
	}
	sum := i * j
	for n := 0; n < time; n++ {
		iIndex, jIndex := n, n
		for jIndex < j-n {
			res = append(res, matrix[iIndex][jIndex])
			jIndex++
		}
		if len(res) == sum {
			return res
		}

		jIndex--
		iIndex++
		for iIndex < i-n {
			res = append(res, matrix[iIndex][jIndex])
			iIndex++
		}
		if len(res) == sum {
			return res
		}

		iIndex--
		jIndex--
		for jIndex >= n {
			res = append(res, matrix[iIndex][jIndex])
			jIndex--
		}
		if len(res) == sum {
			return res
		}

		jIndex++
		iIndex--
		for iIndex > n {
			res = append(res, matrix[iIndex][jIndex])
			iIndex--
		}
	}
	return res
}

// TODO 29-顺时针打印矩阵 100% 状态机
func spiralOrder_(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		res         = make([]int, 0)
		left, right = 0, len(matrix[0]) - 1 // 左右
		up, down    = 0, len(matrix) - 1    // 上下
		flag        = 1                     // 标记走哪里
		i, j        = 0, 0
	)
	for left <= right && up <= down {
		switch flag {
		case 1:
			// 左到右 再往下
			for i = left; i <= right; i++ {
				res = append(res, matrix[j][i])
			}
			i--
			up++

		case 2:
			// 上到下 再往左
			for j = up; j <= down; j++ {
				res = append(res, matrix[j][i])
			}
			j--
			right--

		case 3:
			// 右到做 再往上
			for i = right; i >= left; i-- {
				res = append(res, matrix[j][i])
			}
			i++
			down--

		case 4:
			// 下到上 再往右
			for j = down; j >= up; j-- {
				res = append(res, matrix[j][i])
			}
			j++
			left++
		}

		flag++
		if flag == 5 {
			flag = 1
		}
	}
	return res
}

// TODO 29-顺时针打印矩阵 官方 状态机
func spiralOrder__(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	left, right := len(matrix), len(matrix[0])
	vMap := make([]bool, left*right)
	var (
		total       = left * right
		res         = make([]int, total)
		row, column = 0, 0
		dirs        = [][]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}
		dir = 0
	)
	for i := 0; i < total; i++ {
		res[i] = matrix[row][column]
		vMap[row*right+column] = true
		nextRow := row + dirs[dir][0]
		nextColumn := column + dirs[dir][1]
		if nextRow < 0 || nextRow >= left ||
			nextColumn < 0 || nextColumn >= right ||
			vMap[nextRow*right+nextColumn] {
			dir = (dir + 1) % 4
		}
		row += dirs[dir][0]
		column += dirs[dir][1]
	}
	return res
}
