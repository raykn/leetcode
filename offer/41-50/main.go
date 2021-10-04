package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring_("aab"))
	// fmt.Println(countDigitOne(30000))
	// fmt.Println(countDigitOne(400000))
	// fmt.Println(countDigitOne(1000000))
	// fmt.Println(countDigitOne(2000000))
	// fmt.Println(countDigitOne(3000000))
}

// TODO 42-连续子数组的最大和
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// TODO 43-1～n 整数中1出现的次数 - https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/solution/mian-shi-ti-43-1n-zheng-shu-zhong-1-chu-xian-de-2/
func countDigitOne(n int) int {
	// cur = 0: res += high * digit
	// cur = 1: res += high * digit + low +1
	// cur > 1: res += high * digit + digit

	res := 0
	digit := 1     // 个位开始
	high := n / 10 // high 位
	cur := n % 10  // cur位
	low := 0       // low 位
	for high != 0 || cur != 0 {
		switch cur {
		case 0:
			res += high * digit
		case 1:
			res += high*digit + low + 1
		default:
			res += (high + 1) * digit
		}
		low += cur * digit // 低位
		cur = high % 10    // 下一个cur == high最后一位
		high /= 10         // 高位
		digit *= 10        // 进位
	}
	return res
}

// TODO 44-数字序列中某一位的数字
func findNthDigit(n int) int {
	// range     nums    size
	// 1  ~ 9     9      9
	// 10 ~ 99    90     180
	// ...
	//  x ~ y     9x     9xy
	digit, digitNum, count := 1, 1, 9
	for n > count {
		n -= count
		digit++
		digitNum *= 10
		count = 9 * digit * digitNum
	}
	num := digitNum + (n-1)/digit // real target num
	index := (n - 1) % digit
	return int(strconv.Itoa(num)[index] - '0')
}

// TODO 45-把数组排成最小的数
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})
	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprintf("%d", nums[i]))
	}
	return res.String()
}

func compare(a, b int) bool {
	s1 := fmt.Sprintf("%d%d", a, b)
	s2 := fmt.Sprintf("%d%d", b, a)
	return s1 < s2
}

// TODO 45-把数组排成最小的数 快排
func minNumber_(nums []int) string {
	n := len(nums)
	if n < 1 {
		return ""
	}

	strs := make([]string, n)
	for i, _ := range nums {
		strs[i] = fmt.Sprintf("%d", nums[i])
	}
	quickSort(strs, 0, n-1)

	return strings.Join(strs, "")
}

func quickSort(strs []string, lo, hi int) {
	if lo >= hi { // 递归出口
		return
	}
	part := partion(strs, lo, hi)
	quickSort(strs, lo, part-1)
	quickSort(strs, part+1, hi)
}

func partion(strs []string, lo, hi int) int {
	temp := strs[lo]
	i, j := lo, hi
	for i < j {
		for ; i < j; j-- {
			if strs[j]+temp < temp+strs[j] {
				break
			}
		}
		strs[i] = strs[j]

		for ; i < j; i++ {
			if strs[i]+temp > temp+strs[i] {
				break
			}
		}
		strs[j] = strs[i]
	}

	strs[i] = temp
	return i
}

// TODO 46-把数字翻译成字符串 dp优化
func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, sum := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, sum = q, sum, 0
		sum = q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			sum += p
		}
	}
	return sum
}

// TODO 46-把数字翻译成字符串 dfs
func translateNum_(num int) int {
	if num == 0 {
		return 1
	}
	var dp func(n int) int
	dp = func(n int) int {
		if n < 10 {
			return 1
		}
		tmp := n % 100
		if tmp < 26 && tmp > 9 {
			return dp(n/10) + dp(n/100)
		}
		return dp(n / 10)
	}
	return dp(num)
}

// TODO 46-把数字翻译成字符串 dp不优化
func translateNum__(num int) int {
	src := strconv.Itoa(num)
	dp := make([]int, len(src)+1)
	dp[0] = 1
	for i := 1; i < len(src); i++ {
		dp[i] = dp[i-1]
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			if i > 1 {
				dp[i] += dp[i-2]
			} else {
				dp[i]++
			}
		}
	}
	return dp[len(src)-1]
}

// TODO 47-礼物的最大价值
func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO 48-最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := []byte(s)
	exist := make(map[byte]int)
	start, max := 0, 0
	for i := 0; i < len(str); i++ {
		if index, ok := exist[str[i]]; ok {
			if i-start > max {
				max = i - start
			}
			start = index + 1
			i = start
			exist = make(map[byte]int)
		}
		exist[str[i]] = i
	}
	if len(str)-start > max {
		max = len(str) - start
	}
	return max
}

// TODO 48-最长不含重复字符的子字符串 优化版
func lengthOfLongestSubstring_(s string) int {
	if len(s) == 0 {
		return 0
	}
	exist := make(map[byte]int)
	max, tmp := 1, 0
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		index, ok := exist[str[i]]
		exist[str[i]] = i
		if ok && tmp >= i-index {
			tmp = i - index
		} else {
			tmp++
		}
		max = maxN(max, tmp)
	}
	return max
}

func maxN(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// TODO 49-丑数 最小堆做法
var factors = []int{2, 3, 5}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}

// TODO 49-丑数 动态规划
func nthUglyNumber_(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(x2, min(x3, x5))
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TODO 50-第一个只出现一次的字符
func firstUniqChar(s string) byte {
	sBytes := []byte(s)
	existMap := make([]int, 30)
	for _, v := range sBytes {
		existMap[v-'a']++
	}
	for _, v := range sBytes {
		if existMap[v-'a'] == 1 {
			return v
		}
	}
	return ' '
}
