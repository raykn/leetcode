package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
	fmt.Println(isStraight([]int{0, 2, 3, 4, 5}))
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
}

// TODO 60. n个骰子的点数
func dicesProbability(n int) []float64 {
	dp := make([][]float64, n)
	for i := range dp {
		// 1-6   2-12   3-18   ...
		dp[i] = make([]float64, (i+1)*6-i)
	}
	for i := range dp[0] {
		dp[0][i] = float64(1) / float64(6)
	}
	for i := 1; i < len(dp); i++ {
		for j := range dp[i-1] {
			// 其实就是 k < 6
			for k := range dp[0] {
				dp[i][j+k] += float64(dp[i-1][j]) * float64(dp[0][k])
			}
		}
	}
	return dp[n-1]
}

// TODO 60. n个骰子的点数 ———— 递归解法
func dicesProbability_(n int) []float64 {
	// 所有可能出现的点数之和，他们出现的次数
	count := make([]int, 5*n+1)
	// pre，前面摇到的所有号之和
	// nn 还剩几次，即还要摇几个。
	var next func(pre int, nn int)
	next = func(pre int, nn int) {
		// 还剩一个筛子的时候，直接加上，不递归了，节约时间。
		if nn == 1 {
			for i := 1; i < 7; i++ {
				count[pre-n+i]++
			}
			return
		}
		// 本次可能出现的点数
		for i := 1; i < 7; i++ {
			// 递归求下一个
			next(i+pre, nn-1)
		}
	}
	next(0, n)
	all := 0
	for _, v := range count {
		all += v
	}
	res := make([]float64, 5*n+1)
	for i := 0; i < 5*n+1; i++ {
		res[i] = float64(count[i]) / float64(all)
	}
	return res
}

// TODO 61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sort.Ints(nums)
	i, n := 0, 0
	for i < len(nums) && nums[i] == 0 {
		i++
		n++
	}
	i++
	for ; i < len(nums); i++ {
		sub := nums[i] - nums[i-1]
		switch sub {
		case 0:
			return false
		case 1:
		default:
			n = n - sub + 1
			if n < 0 {
				return false
			}
		}
	}
	return true
}

// TODO 62. 圆圈中最后剩下的数字
func lastRemaining(n int, m int) int {
	var f func(nn, mm int) int
	f = func(nn, mm int) int {
		if nn == 1 {
			return 0
		}
		x := f(nn-1, mm)
		return (mm + x) % nn
	}
	return f(n, m)
}

// TODO 62. 圆圈中最后剩下的数字 ———— 迭代
func lastRemaining_(n int, m int) int {
	var res int
	for i := 2; i != n+1; i++ {
		res = (m + res) % i
	}
	return res
}

// TODO 63. 股票的最大利润
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	r := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > r {
			r = prices[i]
		} else if prices[i] < r {
			max = Max(max, r-prices[i])
		}
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// TODO 64. 求1+2+…+n
func sumNums(n int) int {
	var res = 0
	var sum func(nn int) bool
	sum = func(nn int) bool {
		res += nn
		return nn > 0 && sum(nn-1)
	}
	sum(n)
	return res
}

// TODO 65. 不用加减乘除做加法
func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}

// TODO 66. 构建乘积数组
func constructArr(a []int) []int {
	res := make([]int, len(a))
	if len(a) == 0 {
		return res
	}
	for i := 0; i < len(a); i++ {
		res[i] = 1
	}

	left, right := 1, 1
	for i := 0; i < len(a); i++ {
		res[i] *= left
		left *= a[i]
	}
	for i := len(a) - 1; i >= 0; i-- {
		res[i] *= right
		right *= a[i]
	}
	return res
}

// TODO 66. 构建乘积数组 ———— 一次遍历
func constructArr_(a []int) []int {
	res := make([]int, len(a))
	if len(a) == 0 {
		return res
	}
	for i := 0; i < len(a); i++ {
		res[i] = 1
	}

	left, right := 1, 1
	for i := 0; i < len(a); i++ {
		res[i] *= left
		left *= a[i]

		res[len(a)-i-1] *= right
		right *= a[len(a)-i-1]
	}
	return res
}

// TODO 67. 把字符串转换成整数
func strToInt(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}

		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign * result
}

// TODO 68 - II. 二叉树的最近公共祖先 ———— 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
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

// TODO 68 - II. 二叉树的最近公共祖先 ———— 空间换一次遍历
func lowestCommonAncestor_(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	vis := make(map[int]bool)
	var dfs func(*TreeNode)
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
		vis[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if vis[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}
