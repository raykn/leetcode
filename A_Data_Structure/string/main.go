package main

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

// 题目：最长公共前缀
// 横向比较
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	tmp := strs[0]
	for i := 1; i < len(strs); i++ {
		tmp = lcp(tmp, strs[i])
	}
	return tmp
}

func lcp(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	lenMin := 0
	if l1 > l2 {
		lenMin = l2
	} else {
		lenMin = l1
	}

	index := 0
	for index < lenMin && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

// 纵向比较
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 考虑使用字典树
func longestCommonPrefix2(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	// 实现字典树
	t := Construct()
	// 将strs里边的单词插入字典树中
	for _, str := range strs {
		t.Insert(str)
	}
	// 最长公共前缀的长度
	length := math.MaxInt32
	for _, str := range strs {
		// 最长公共前缀的就是找所有字符串中的最短的前缀
		length = int(math.Min(float64(length), float64(t.LargestCommonPrefix(str))))
	}
	return strs[0][:length]
}

// 实现字典树
type Trie struct {
	child map[byte]*Trie
}

func Construct() *Trie {
	return &Trie{
		child: make(map[byte]*Trie, 0),
	}
}

func (t *Trie) Insert(word string) {
	cur := t
	for i := 0; i < len(word); i++ {
		// 当word不在字典树中
		if cur.child[word[i]] == nil {
			cur.child[word[i]] = &Trie{
				child: make(map[byte]*Trie, 0),
			}
		}
		cur = cur.child[word[i]]
	}
}

// 最长公共前缀
func (t *Trie) LargestCommonPrefix(word string) int {
	cur := t
	count := 0 // 最长公共前缀的长度
	for i := 0; i < len(word); i++ {
		// 当前字符在trie树中，以及不分叉时（表示当前字符属于最长公共前缀）
		if cur.child[word[i]] != nil && len(cur.child) == 1 {
			cur = cur.child[word[i]]
			count++
		} else {
			// 遇到分岔口时，就结束
			break
		}
	}
	return count
}

// 题目：最长回文子串
func longestPalindrome(s string) string {
	res := ""
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(res) {
				res = s[i : i+l+1]
			}

		}

	}
	return res
}

// 516. 最长回文子序列
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[0][n-1]
}

// 647. 回文子串
var num = 0

func countSubstrings(s string) int {
	for i := 0; i < len(s); i++ {
		count(s, i, i)
		count(s, i, i+1)
	}
	return num
}
func count(s string, start, end int) {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		num++
		start--
		end++
	}
}

// 最长无重复字符串
func lengthOfLongestSubstring(s string) int {
	keyMap := make(map[int32]int)
	tmp, res := 0, 0
	for j, str := range s {
		index, ok := keyMap[str]
		if !ok {
			index = -1
		}
		keyMap[str] = j
		if tmp < j-index {
			tmp = tmp + 1
		} else {
			tmp = j - index
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

// 151. 翻转字符串里的单词
// 过滤多空格 反转整个字符串 再反转单词
// golang 的 string 是不可变的 需要强转才能使用
func str2bytes(s string) (res []byte) {
	x := (*reflect.StringHeader)(unsafe.Pointer(&s))
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&res))
	slice.Data = x.Data
	slice.Len = x.Len
	slice.Cap = x.Len
	return
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func reverseWords(s string) string {
	sArr := str2bytes(s)
	sLen := len(sArr)
	index := 0 // 实际长度指针
	currIndex := 0
	for sArr[currIndex] == ' ' {
		currIndex++ // 找到第一个空格
	}
	// 先缩减
	fillSpace := false
	for ; currIndex < sLen; currIndex++ {
		if sArr[currIndex] != ' ' {
			if fillSpace {
				sArr[index] = ' '
				index++
			}
			fillSpace = true

			for currIndex < sLen && sArr[currIndex] != ' ' {
				sArr[index] = sArr[currIndex]
				index++
				currIndex++
			}
		}
	}

	// 整个反转
	for i := 0; i < index/2; i++ {
		sArr[i], sArr[index-i-1] = sArr[index-i-1], sArr[i]
	}

	// 单词反转
	left := 0
	for i := 1; i < index; i++ {
		if sArr[i] == ' ' {
			// 反转单词
			for x := 0; x < (i-left)/2; x++ {
				sArr[left+x], sArr[i-x-1] = sArr[i-x-1], sArr[left+x]
			}
			left = i + 1
		}
	}
	for x := 0; x < (index-left)/2; x++ {
		sArr[left+x], sArr[index-x-1] = sArr[index-x-1], sArr[left+x]
	}
	return string(sArr[0:index])
}

func PreProcess(s string) string {
	l := len(s)
	var res []byte
	flag := 1                     // 用于处理多个连续空格
	for l != 0 && s[l-1] == ' ' { // 处理字符串后面的空格
		l--
	}
	for i := 0; i < l; i++ {
		if s[i] != ' ' {
			res = append(res, s[i])
			flag = 0
		}
		if s[i] == ' ' && flag == 0 {
			res = append(res, s[i])
			flag = 1
		}
	}
	return string(res)
}

func main() {

	fmt.Println(reverseWords(bytes.NewBufferString("the sky is blue").String()))
	fmt.Println(reverseWords(bytes.NewBufferString("the sky is blue").String()))
	fmt.Println(reverseWords(bytes.NewBufferString(" the sky is blue  ").String()))
}

func showInfo(array [][]int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}
