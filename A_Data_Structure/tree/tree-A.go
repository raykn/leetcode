package tree

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// todo 637. 二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	var res []float64
	res = append(res, float64(root.Val))
	q := []*TreeNode{root}
	for len(q) > 0 {
		var tmp []*TreeNode
		var num int
		var count int
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				num += q[i].Left.Val
				tmp = append(tmp, q[i].Left)
				count++
			}
			if q[i].Right != nil {
				num += q[i].Right.Val
				tmp = append(tmp, q[i].Right)
				count++
			}
		}
		if count > 0 {
			res = append(res, float64(num)/float64(count))
		}
		q = tmp
	}
	return res
}
func averageOfLevels_(root *TreeNode) []float64 {
	var res []float64
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := q
		q = nil
		var num int
		for _, v := range nextLevel {
			num += v.Val
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		res = append(res, float64(num)/float64(len(nextLevel)))
	}
	return res
}
func averageOfLevels__(root *TreeNode) []float64 {
	levelData := []data{}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(levelData) {
			levelData[level].sum += node.Val
			levelData[level].count++
		} else {
			levelData = append(levelData, data{sum: node.Val, count: 1})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	res := make([]float64, len(levelData))
	for i, d := range levelData {
		res[i] = float64(d.sum) / float64(d.count)
	}
	return res
}

// todo 617. 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil || root2 == nil {
		if root1 == nil {
			return &TreeNode{
				Val:   root2.Val,
				Left:  mergeTrees(nil, root2.Left),
				Right: mergeTrees(nil, root2.Right),
			}
		}
		if root2 == nil {
			return &TreeNode{
				Val:   root1.Val,
				Left:  mergeTrees(root1.Left, nil),
				Right: mergeTrees(root1.Right, nil),
			}
		}

	} else {
		return &TreeNode{
			Val:   root1.Val + root2.Val,
			Left:  mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
		}
	}
	return nil
}
func mergeTrees_(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees_(root1.Left, root2.Left),
		Right: mergeTrees_(root1.Right, root2.Right),
	}
}
func mergeTrees__(t1, t2 *TreeNode) *TreeNode {
	// 迭代做法
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	merged := &TreeNode{Val: t1.Val + t2.Val}
	queue := []*TreeNode{merged}
	queue1 := []*TreeNode{t1}
	queue2 := []*TreeNode{t2}
	for len(queue1) > 0 && len(queue2) > 0 {
		node := queue[0]
		queue = queue[1:]
		node1 := queue1[0]
		queue1 = queue1[1:]
		node2 := queue2[0]
		queue2 = queue2[1:]
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if left1 != nil || left2 != nil {
			if left1 != nil && left2 != nil {
				left := &TreeNode{Val: left1.Val + left2.Val}
				node.Left = left
				queue = append(queue, left)
				queue1 = append(queue1, left1)
				queue2 = append(queue2, left2)
			} else if left1 != nil {
				node.Left = left1
			} else { // left2 != nil
				node.Left = left2
			}
		}
		if right1 != nil || right2 != nil {
			if right1 != nil && right2 != nil {
				right := &TreeNode{Val: right1.Val + right2.Val}
				node.Right = right
				queue = append(queue, right)
				queue1 = append(queue1, right1)
				queue2 = append(queue2, right2)
			} else if right1 != nil {
				node.Right = right1
			} else { // right2 != nil
				node.Right = right2
			}
		}
	}
	return merged
}

// todo 606. 根据二叉树创建字符串
func tree2str(root *TreeNode) string {
	var res string
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		if r == nil {
			return
		}
		res += "("
		res += strconv.Itoa(r.Val)
		if r.Left == nil || r.Right == nil {
			if r.Left != nil {
				f(r.Left)
				res += ")"
				return
			}
			if r.Right != nil {
				res += "()"
				f(r.Right)
				res += ")"
				return
			}
		}
		f(r.Left)
		f(r.Right)
		res += ")"
	}
	f(root)
	return res[1 : len(res)-1]
}
func tree2str_(root *TreeNode) string {
	var f func(r *TreeNode) string
	f = func(r *TreeNode) string {
		if r == nil {
			return ""
		}
		if r.Left == nil && r.Right == nil {
			return strconv.Itoa(r.Val)
		}
		if r.Right == nil {
			return strconv.Itoa(r.Val) + "(" + f(r.Left) + ")"
		}
		return strconv.Itoa(r.Val) + "(" + f(r.Left) + ")(" + f(r.Right) + ")"
	}
	return f(root)
}
func tree2str__(root *TreeNode) string {
	// 栈解决
	var str strings.Builder
	if root == nil {
		return ""
	}
	stack := []interface{}{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n, ok := node.(*TreeNode); ok {
			str.WriteString(fmt.Sprintf("%d", n.Val))
			// 先进右 因为先进后出
			if n.Right != nil {
				stack = append(stack, ")")
				stack = append(stack, n.Right)
				stack = append(stack, "(")
			}
			if n.Left != nil {
				stack = append(stack, ")")
				stack = append(stack, n.Left)
				stack = append(stack, "(")
			}
			if n.Right != nil && n.Left == nil {
				stack = append(stack, "()")
			}

		} else {
			str.WriteString(node.(string))
		}
	}
	return str.String()
}

// todo 563. 二叉树的坡度
func findTilt(root *TreeNode) int {
	var res int
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := f(root.Left)
		right := f(root.Right)
		res += int(math.Abs(float64(left - right)))
		return left + right + root.Val
	}
	f(root)
	return res
}

// todo 572. 另一个树的子树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return sameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func sameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// if (p != nil && q == nil) || (p == nil && q != nil) {
	// 	return false
	// }
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// todo 面试 04.02 最小高度树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// todo 100. 相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// if (p != nil && q == nil) || (p == nil && q != nil) {
	// 	return false
	// }
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// todo 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	// 自顶向下
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}
func isBalanced_(root *TreeNode) bool {
	// 从底至顶
	return help(root) >= 0
}
func help(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := help(root.Left)
	rHeight := help(root.Right)
	if lHeight == -1 || rHeight == -1 || abs(lHeight-rHeight) > 1 {
		return -1
	}
	return max(lHeight, rHeight) + 1
}
func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// todo 11. 二叉树的最小深度
func minDepth(root *TreeNode) int {
	// 最小深度状态保存 下传
	// 深度优先
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func minDepth_(root *TreeNode) int {
	// 广度优先
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return 1
	}
	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}

// todo 112. 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
func hasPathSum_(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queNode := []*TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0 {
		now := queNode[0]
		queNode = queNode[1:]
		valTmp := queVal[0]
		queVal = queVal[1:]
		// 到叶子
		if now.Left == nil && now.Right == nil {
			if valTmp == targetSum {
				return true
			}
			continue
		}
		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+valTmp)
		}
		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+valTmp)
		}
	}
	return false
}

// todo 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// todo 257. 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	var f func(root *TreeNode, path []int)
	f = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		// 叶子
		if root.Left == nil && root.Right == nil {
			var p string
			var length = len(path) - 1
			for i, v := range path {
				p += strconv.Itoa(v)
				if i != length {
					p += "->"
				}
			}
			paths = append(paths, p)
			return
		}
		f(root.Left, path)
		f(root.Right, path)
	}
	f(root, []int{})
	return paths
}
func binaryTreePaths_(root *TreeNode) (paths []string) {
	// 官方
	var f func(root *TreeNode, path string)
	f = func(root *TreeNode, path string) {
		if root != nil {
			pathSub := path
			pathSub += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				paths = append(paths, pathSub)
			} else {
				pathSub += "->"
				f(root.Left, pathSub)
				f(root.Right, pathSub)
			}
		}
	}
	f(root, "")
	return paths
}
func binaryTreePaths__(root *TreeNode) (paths []string) {
	// 官方 广度
	if root == nil {
		return paths
	}
	nodeQueue := []*TreeNode{}
	pathQueue := []string{}
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))
	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}

// todo 404. 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	var f func(root *TreeNode, lOrR bool)
	f = func(root *TreeNode, lOrR bool) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if lOrR {
				sum += root.Val
			}
			return
		}
		f(root.Left, true)
		f(root.Right, false)
	}
	f(root, false)
	return sum
}
func sumOfLeftLeaves_(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			if isLeafNode(node.Left) {
				ans += node.Left.Val
			} else {
				q = append(q, node.Left)
			}
		}
		if node.Right != nil && !isLeafNode(node.Right) {
			q = append(q, node.Right)
		}
	}
	return ans
}
func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func findMode(root *TreeNode) (answer []int) {
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		update(root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return
}
func findMode_(root *TreeNode) (answer []int) {
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return
}
func findMode__(root *TreeNode) []int {
	res := make(map[int]int)
	res = traveans(res, root)
	max := 0
	max = Findmax(res)
	ans := []int{}
	for k, v := range res {
		if v == max {
			ans = append(ans, k)
		}
	}
	return ans
}

func traveans(res map[int]int, root *TreeNode) map[int]int {
	if root == nil {
		return res
	}
	res[root.Val]++
	res = traveans(res, root.Left)
	res = traveans(res, root.Right)
	return res
}

func Findmax(res map[int]int) int {
	max := 0
	for _, v := range res {
		if v > max {
			max = v
		}
	}
	return max
}
