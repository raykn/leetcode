package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(printNumbers_(3))
}

// TODO 11-旋转数组的最小数字
func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		mid := l + (r-l)/2
		if numbers[mid] < numbers[r] {
			r = mid
		} else if numbers[mid] > numbers[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return numbers[l]
}

// TODO 12-矩阵中的路径 我的
type dir struct {
	d1, d2 int
}

var dirList = []dir{
	{d1: 0, d2: 1},
	{d1: 0, d2: -1},
	{d1: 1, d2: 0},
	{d1: -1, d2: 0},
}

type vv struct {
	i, j int
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	target := []byte(word)
	start := target[0]
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == start {
				vMap := make(map[vv]bool, 0)
				vMap[vv{i: i, j: j}] = true
				if visit(board, target[1:], i, j, vMap) {
					return true
				}
			}
		}
	}
	return false
}

func visit(board [][]byte, target []byte, i, j int, vMap map[vv]bool) bool {
	if len(target) == 0 {
		return true
	}
	start := target[0]
	iLimit, jLimit := len(board), len(board[0])
	for _, v := range dirList {
		iIndex, jIndex := i+v.d1, j+v.d2
		if iIndex < 0 || jIndex < 0 || iIndex == iLimit || jIndex == jLimit {
			continue
		}
		if vMap[vv{i: iIndex, j: jIndex}] || board[iIndex][jIndex] != start {
			continue
		}
		vMap[vv{i: iIndex, j: jIndex}] = true
		if visit(board, target[1:], iIndex, jIndex, vMap) {
			return true
		}
		vMap[vv{i: iIndex, j: jIndex}] = false
	}
	return false
}

// TODO 12-矩阵中的路径 100%
func exist_(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(board, i, j, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, index int) bool {
	if index == len(word)-1 {
		return word[index] == board[i][j]
	}

	char := board[i][j]
	board[i][j] = ' '
	if i != 0 {
		if board[i-1][j] == word[index+1] {
			if dfs(board, i-1, j, word, index+1) {
				return true
			}
		}
	}
	if i != len(board)-1 {
		if board[i+1][j] == word[index+1] {
			if dfs(board, i+1, j, word, index+1) {
				return true
			}
		}
	}
	if j != 0 {
		if board[i][j-1] == word[index+1] {
			if dfs(board, i, j-1, word, index+1) {
				return true
			}
		}
	}
	if j != len(board[0])-1 {
		if board[i][j+1] == word[index+1] {
			if dfs(board, i, j+1, word, index+1) {
				return true
			}
		}
	}
	board[i][j] = char
	return false
}

// TODO 13-机器人的运动范围
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	var queue [][]int

	var dirList [2][]int
	dirList[0] = append(dirList[0], []int{0, 1}...)
	dirList[1] = append(dirList[1], []int{1, 0}...)
	type vv struct {
		i, j int
	}
	vMap := make(map[vv]bool)
	vMap[vv{}] = true
	queue = append(queue, []int{0, 0})
	var sum = 1
	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		x, y := tmp[0], tmp[1]
		for _, v := range dirList {
			tx, ty := x+v[0], y+v[1]
			if tx < 0 || tx >= m || ty < 0 || ty >= n ||
				vMap[vv{i: tx, j: ty}] ||
				getNum(tx)+getNum(ty) > k {
				continue
			}
			queue = append(queue, []int{tx, ty})
			vMap[vv{i: tx, j: ty}] = true
			sum++
		}
	}
	return sum
}

func getNum(i int) int {
	var res int
	for i != 0 {
		res += i % 10
		i /= 10
	}
	return res
}

// TODO 13-机器人的运动范围
func movingCount_(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	type vv struct {
		i, j int
	}
	vMap := make(map[vv]bool)
	vMap[vv{}] = true

	var queue [][]int
	queue = append(queue, []int{0, 0})
	var sum = 1
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		x, y := t[0], t[1]
		if x+1 < m && !vMap[vv{i: x + 1, j: y}] {
			if CalcSum(x+1, y) <= k {
				sum++
				vMap[vv{i: x + 1, j: y}] = true
				queue = append(queue, []int{x + 1, y})
			}
		}
		if y+1 < n && !vMap[vv{i: x, j: y + 1}] {
			if CalcSum(x, y+1) <= k {
				sum++
				vMap[vv{i: x, j: y + 1}] = true
				queue = append(queue, []int{x, y + 1})
			}
		}
	}
	return sum
}

func CalcSum(x, y int) int {
	var sum = 0
	for x >= 10 {
		sum += x % 10
		x = x / 10
	}
	sum += x
	for y >= 10 {
		sum += y % 10
		y = y / 10
	}
	sum += y
	return sum
}

// TODO 13-机器人的运动范围 100%
type pos struct {
	x int
	y int
}

func (p pos) CalcSum() int {
	var x int = p.x
	var y int = p.y
	var sum int = 0
	for x >= 10 {
		sum += x % 10
		x = x / 10
	}
	sum += x
	for y >= 10 {
		sum += y % 10
		y = y / 10
	}
	sum += y
	return sum
}

func movingCount__(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	var fini []bool = make([]bool, m*n)
	var vec []pos = make([]pos, 0)
	var ans int = 0
	vec = append(vec, pos{0, 0})
	fini[0*n+0] = true
	ans++
	for len(vec) > 0 {
		var x int = vec[0].x
		var y int = vec[0].y
		if x+1 < m && !fini[(x+1)*n+y] {
			var pos pos = pos{x + 1, y}
			var sum int = pos.CalcSum()
			if sum <= k {
				ans++
				vec = append(vec, pos)
				fini[(x+1)*n+y] = true
			}
		}
		if y+1 < n && !fini[x*n+y+1] {
			var pos pos = pos{x, y + 1}
			var sum int = pos.CalcSum()
			if sum <= k {
				ans++
				vec = append(vec, pos)
				fini[x*n+y+1] = true
			}
		}
		vec = vec[1:]
	}
	return ans
}

// TODO 13-机器人的运动范围 非递归解法 100% 动态规划方程式
func movingCount___(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	visited := make([]bool, m*n)
	visited[0] = true
	ans := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || get(i)+get(j) > k {
				continue
			}
			if i-1 >= 0 && visited[(i-1)*n+j] {
				visited[i*n+j] = true
			}
			if j-1 >= 0 && visited[i*n+(j-1)] {
				visited[i*n+j] = true
			}
			if visited[i*n+j] {
				ans++
			}
		}
	}
	return ans
}

func get(x int) int {
	res := 0
	for x != 0 {
		res += x % 10
		x /= 10
	}
	return res
}

// TODO 14-I. 剪绳子
func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	x := n % 3
	switch x {
	case 0:
		return int(math.Pow(3, float64(int(n/3))))
	case 1:
		return 4 * int(math.Pow(3, float64(int(n/3-1))))
	default:
		return 2 * int(math.Pow(3, float64(int(n/3))))
	}
}

// TODO 14-I. 剪绳子 动态规划
func cuttingRope_(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			dp[i] = int(
				math.Max(float64(dp[i]),
					math.Max(float64(j*(i-j)), float64(j*dp[i-j]))))
		}
	}
	return dp[n]
}

// TODO 14-II. 剪绳子 取模 1000000007
func cuttingRope__(n int) int {
	if n <= 3 {
		return n - 1
	}
	ret := 1
	for n > 4 {
		ret = ret * 3 % 1000000007
		n -= 3
	}
	return ret * n % 1000000007
}

// TODO 15-二进制中1的个数
func hammingWeight(num uint32) int {
	var res int
	for ; num > 0; num &= num - 1 {
		res++
	}
	return res
}

// TODO 15-二进制中1的个数
func hammingWeight_(num uint32) int {
	var res int
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			res++
		}
	}
	return res
}

// TODO 16-数值的整数次方 递归快速幂
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := myPow(x, n>>1)
	if n%2 == 0 {
		return res * res
	}
	return x * res * res
}

// TODO 16-数值的整数次方 迭代快速幂
func myPow_(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	for n >= 1 {
		if n&1 == 1 {
			res *= x
			n--
		} else {
			x *= x
			n = n >> 1
		}
	}
	return res
}

// TODO 17-打印从1到最大的n位数
func printNumbers(n int) []int {
	index := 1
	for n > 0 {
		n--
		index *= 10
	}
	index--
	res := make([]int, index)
	for i := 0; i < index; i++ {
		res[i] = i + 1
	}
	return res
}

// TODO 17-打印从1到最大的n位数 大数处理！！理论上返回字符串才对
func printNumbers_(n int) []int {
	res := make([]int, int(math.Pow(10, float64(n)))-1)
	count := 0
	var dfs func(index int, num []byte, now int)
	dfs = func(index int, num []byte, now int) {
		if index == now {
			t, _ := strconv.Atoi(string(num))
			res[count] = t
			count++
			return
		}
		for i := '0'; i <= '9'; i++ {
			num[index] = byte(i)
			dfs(index+1, num, now)
		}
	}
	for i := 1; i <= n; i++ {
		for j := '1'; j <= '9'; j++ {
			num := make([]byte, i)
			num[0] = byte(j)
			dfs(1, num, i)
		}
	}
	return res
}

// TODO 18-删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	tail := &ListNode{}
	tail.Next = head

	pre := tail
	curr := pre.Next
	for curr != nil {
		if curr.Val == val {
			pre.Next = curr.Next
			break
		}
		pre = curr
		curr = curr.Next
	}
	return tail.Next
}

// TODO 18-删除链表的节点 双100%
func deleteNode_(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}
	return head
}

// TODO 19-正则表达式匹配【困难】 动态规划
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

// TODO 19-正则表达式匹配【困难】 递归
func isMatch_(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	var first bool
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		first = true
	}
	if len(p) > 1 && p[1] == '*' {
		return (first && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
	return first && isMatch(s[1:], p[1:])
}

// TODO 20-表示数值的字符串
type state int
type charType int

// 起始的空格
// 符号位
// 整数部分
// 左侧有整数的小数点
// 左侧无整数的小数点（根据前面的第二条额外规则，需要对左侧有无整数的两种小数点做区分）
// 小数部分
// 字符 \text{e}e
// 指数部分的符号位
// 指数部分的整数部分
// 末尾的空格
const (
	StateInitial state = iota
	StateIntSign
	StateInteger
	StatePoint
	StatePointWithoutInt
	StateFraction
	StateExp
	StateExpSign
	StateExpNumber
	StateEnd
)

const (
	CharNumber charType = iota
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
)

func toCharType(ch byte) charType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}

func isNumber(s string) bool {
	// 每一个状态可能的下一个状态
	transfer := map[state]map[charType]state{
		StateInitial: {
			CharSpace:  StateInitial,
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
			CharSign:   StateIntSign,
		},
		StateIntSign: {
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateInteger: {
			CharNumber: StateInteger,
			CharExp:    StateExp,
			CharPoint:  StatePoint,
			CharSpace:  StateEnd,
		},
		StatePoint: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StatePointWithoutInt: {
			CharNumber: StateFraction,
		},
		StateFraction: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StateExp: {
			CharNumber: StateExpNumber,
			CharSign:   StateExpSign,
		},
		StateExpSign: {
			CharNumber: StateExpNumber,
		},
		StateExpNumber: {
			CharNumber: StateExpNumber,
			CharSpace:  StateEnd,
		},
		StateEnd: {
			CharSpace: StateEnd,
		},
	}
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == StateInteger ||
		state == StatePoint ||
		state == StateFraction ||
		state == StateExpNumber ||
		state == StateEnd
}
