package tree

import "math"

// todo 98. 验证二叉搜索树【中等】
func isValidBST(root *TreeNode) bool {
	// 递归解法
	// 要保证的不仅是某一课树是二叉搜索树 而是即便叶节点和根节点也要满足
	var check func(root *TreeNode, minNum, maxNum int) bool
	check = func(root *TreeNode, minNum, maxNum int) bool {
		if root == nil {
			return true
		}
		if root.Val <= minNum || root.Val >= maxNum {
			return false
		}

		return check(root.Left, minNum, root.Val) && check(root.Right, root.Val, maxNum)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}
func isValidBST_(root *TreeNode) bool {
	// 中序遍历解法
	// 中序遍历会形成有升序排序
	tmpRes := math.MinInt64
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= tmpRes {
			return false
		}
		tmpRes = root.Val
		root = root.Right
	}
	return true
}

// todo 94. 二叉树的中序遍历【中等】
func inorderTraversal(root *TreeNode) (resList []int) {
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		resList = append(resList, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return resList
}
func inorderTraversal_(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
