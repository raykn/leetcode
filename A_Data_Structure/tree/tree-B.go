package main

import (
	"math"
)

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

// todo 94. 二叉树的后序遍历【中等】 ———— 递归
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

// todo 94. 二叉树的后序遍历【中等】 ———— 迭代
func postorderTraversal_(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}

// todo 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		queue []*TreeNode
		res   [][]int
	)
	queue = append(queue, root)
	for len(queue) != 0 {
		var floor []int
		var next []*TreeNode
		for _, v := range queue {
			if v != nil {
				floor = append(floor, v.Val)
				if v.Left != nil {
					next = append(next, v.Left)
				}
				if v.Right != nil {
					next = append(next, v.Right)
				}
			}
		}
		queue = next
		res = append(res, floor)
	}
	return res
}

// todo 236. 二叉树的最近公共祖先 ———— 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// todo 236. 二叉树的最近公共祖先 ———— DFS ！！！！妙呀
func lowestCommonAncestor_(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

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
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

// todo 116. 填充每个节点的下一个右侧节点指针 ———— 也可以作为 II 的答案，以为没有用到完美二叉树的性质
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		var qFloor []*Node
		floor := q
		pre := floor[0]
		if pre.Left != nil {
			qFloor = append(qFloor, pre.Left)
		}
		if pre.Right != nil {
			qFloor = append(qFloor, pre.Right)
		}

		for i := 1; i < len(floor); i++ {
			node := floor[i]
			pre.Next = node
			pre = node
			if node.Left != nil {
				qFloor = append(qFloor, node.Left)
			}
			if node.Right != nil {
				qFloor = append(qFloor, node.Right)
			}
		}
		q = qFloor
	}
	return root
}

// todo 116. 填充每个节点的下一个右侧节点指针 ———— 最优解
func connect_(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}
func connectTwoNode(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)
	connectTwoNode(node1.Right, node2.Left)
}

// todo 117. 填充每个节点的下一个右侧节点指针 II ———— 与 1 一样，只是 root 不是一个完美二叉树
func connect__(root *Node) *Node {
	if root == nil {
		return root
	}

	cur := root
	var start, end *Node = nil, nil
	for cur != nil || start != nil {
		if cur == nil {
			cur = start
			start, end = nil, nil
		}

		if cur.Left != nil {
			if start == nil {
				start = cur.Left
				end = start

			} else {
				end.Next = cur.Left
				end = end.Next
			}
		}

		if cur.Right != nil {
			if start == nil {
				start = cur.Right
				end = start
			} else {
				end.Next = cur.Right
				end = end.Next
			}
		}
		cur = cur.Next
	}
	return root
}
