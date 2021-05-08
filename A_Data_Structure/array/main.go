package main

import (
	"fmt"
	"sort"
)

// 寻找数组的中心索引
func pivotIndex(nums []int) int {
	for i := range nums {
		var sum1, sum2 int
		for left := 0; left < i; left++ {
			sum1 = sum1 + nums[left]
		}
		for right := len(nums) - 1; right > i; right-- {
			sum2 = sum2 + nums[right]
		}
		if sum1 == sum2 {
			return i
		}
	}
	return -1
}

func pivotIndex2(nums []int) int {
	sumMap := make(map[int]int)
	sum := 0
	for i, val := range nums {
		sum += val
		sumMap[i] = sum
	}
	sumMap[-1] = 0
	sumMap[len(nums)] = 0
	for i := 0; i < len(nums); i++ {
		sum1, _ := sumMap[i-1]
		sum2, _ := sumMap[i]
		sum3 := sum - sum2
		if sum1 == sum3 {
			return i
		}
	}
	return -1
}

func pivotIndex3(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	leftSum := 0
	for i, val := range nums {
		if leftSum == (sum - leftSum - val) {
			return i
		}
		leftSum += val
	}
	return -1
}

// 搜索插入位置
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if start > end {
			return start
		}
		if end < start {
			return end
		}
		mid := start + (end-start)/2
		switch {
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			start = mid + 1
		case target < nums[end]:
			end = mid - 1
		}
	}
}

func searchInsert2(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	index := len(nums)
	for start <= end {
		mid := start + (end-start)/2
		if target <= nums[mid] {
			index = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return index
}

// 合并区间 —— 给出一个区间的集合，请合并所有重叠的区间
func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	if len(intervals) == 0 {
		return result
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	tmp := []int{intervals[0][0], intervals[0][1]}
	for _, val := range intervals {
		if val[0] <= tmp[1] && tmp[1] <= val[1] {
			tmp[1] = val[1]
		}

		if tmp[1] < val[0] {
			result = append(result, tmp)
			tmp = val
		}
	}
	if len(result) == 0 || (len(result) != 0 && result[len(result)-1][1] != intervals[len(intervals)-1][1]) {
		result = append(result, tmp)
	}
	return result
}
func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	for _, val := range intervals {
		if len(result) == 0 || val[0] > result[len(result)-1][1] {
			result = append(result, val)
		} else if val[1] > result[len(result)-1][1] {
			result[len(result)-1][1] = val[1]
		}
	}
	return result
}

// 旋转矩阵
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]         // 左上 <- 左下
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1] // 左下 <- 右下
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1] // 右下 <- 右上
			matrix[j][n-i-1] = tmp                  // 右上 <- 左上
		}
	}
}

// 零矩阵 —— M * N 矩阵中某个元素为 0， 则将其所在的行与列清零
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	iLen := len(matrix)
	jLen := len(matrix[0])
	jMap := make([]int, jLen)
	iMap := make([]int, iLen)
	for i := 0; i < iLen; i++ {
		for j := 0; j < jLen; j++ {
			if matrix[i][j] == 0 {
				iMap[i] = 1
				jMap[j] = 1
			}
		}
	}
	for i := 0; i < iLen; i++ {
		for j := 0; j < jLen; j++ {
			iOk := iMap[i]
			jOk := jMap[j]
			if iOk == 1 || jOk == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroes2(matrix [][]int) {
	var isFirstColHaveZero, isFirstRowHaveZero bool
	iLen, jLen := len(matrix), len(matrix[0])
	for _, col := range matrix {
		if col[0] == 0 {
			isFirstColHaveZero = true
		}
	}
	for _, row := range matrix[0] {
		if row == 0 {
			isFirstRowHaveZero = true
		}
	}
	for i := 0; i < iLen; i++ {
		for j := 0; j < jLen; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// todo: 从 1 开始，第 0 列和行都用作标记了
	for i := 1; i < iLen; i++ {
		for j := 1; j < jLen; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	for _, colVal := range matrix {
		if isFirstColHaveZero {
			colVal[0] = 0
		}
	}
	for j := range matrix[0] {
		if isFirstRowHaveZero {
			matrix[0][j] = 0
		}
	}
}

// 对角线遍历 —— 对角线返回矩阵中的元素  todo: 没做完
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	iLen, jLen := len(matrix), len(matrix[0])
	result := make([]int, 0)
	index, sign := 0, 0
	for l := 0; l < iLen+jLen; l++ {
		sign++
		i, j := 0, 0
		if l < jLen {
			i = 0
		} else {
			i = l - jLen + 1
		}
		if l < jLen {
			j = l
		} else {
			j = jLen - 1
		}

		flagIndex := index
		for i < iLen && j > -1 {
			result = append(result, matrix[i][j])
			index++
			i++
			j--
		}
		if sign%2 == 1 {
			for start := 0; start < (index-flagIndex)>>1; start++ {
				result[flagIndex+start], result[index-start-1] = result[index-start-1], result[flagIndex+start]
			}
		}
	}
	return result
}

func main() {
	fmt.Print(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
