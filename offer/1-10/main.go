package main

import "strings"

func main() {
	// 测试
}

// TODO 03-数组中重复的数字
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] && nums[i] == nums[nums[i]] {
			return nums[i]
		}
		for nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}

// 更好理解的一种解法
func findRepeatNumber_(nums []int) int {
	i := 0
	for i < len(nums) {
		if i == nums[i] {
			i++
			continue
		}
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}

// TODO 04-二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := len(matrix), len(matrix[0])
	iIndex, jIndex := 0, j-1
	for iIndex < i && jIndex >= 0 {
		if target == matrix[iIndex][jIndex] {
			return true
		}
		if target > matrix[iIndex][jIndex] {
			iIndex++
		} else {
			jIndex--
		}
	}
	return false
}

// TODO 05-替换空格
func replaceSpace(s string) string {
	var str strings.Builder
	var build = []byte("%20")
	for _, v := range s {
		if v == ' ' {
			str.Write(build)
		} else {
			str.WriteByte(byte(v))
		}
	}
	return str.String()
}

// TODO 06-从尾到头打印链表
func reversePrint(head *ListNode) []int {
	var reverse func(root *ListNode)
	var res []int
	reverse = func(root *ListNode) {
		if root == nil {
			return
		}
		reverse(root.Next)
		res = append(res, root.Val)
	}
	reverse(head)
	return res
}

// TODO 07-重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	var index int
	for ; index < len(inorder); index++ {
		if inorder[index] == val {
			break
		}
	}
	index++ // 用来做分片终点位置 左开右闭 ++好判断
	root := &TreeNode{
		Val: val,
		// 其实这里inorder[:index-1] 也可以 因为把当前val包含进来没啥用
		Left:  buildTree(preorder[1:index], inorder[:index]),
		Right: buildTree(preorder[index:], inorder[index:]),
	}
	return root
}

// TODO 10-斐波那契数列
func fib(n int) int {
	i, j := 0, 1
	for k := 0; k < n; k++ {
		i, j = j, (i+j)%1000000007
	}
	return i
}

// TODO 10-青蛙跳台阶问题
func numWays(n int) int {
	i, j := 0, 1
	for k := 0; k < n; k++ {
		i, j = j, (i+j)%1000000007
	}
	return j
}
