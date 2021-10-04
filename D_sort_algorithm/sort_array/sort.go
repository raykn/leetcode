package main

import (
	"fmt"
)

// MaopaoSort 冒泡排序
func MaopaoSort(nums []int) {
	var swap bool
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

// SelectSort 选择排序
func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}

// InsertSort 插入排序
func InsertSort(nums []int) {
	for i, j := 1, 0; i < len(nums); i++ {
		tmp := nums[i]
		for j = i; j > 0 && nums[j-1] > tmp; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = tmp
	}
}

// InsertSort_ 插入排序 ———— offical
func InsertSort_(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// QuickSort 快速排序 ———— 单路快排
func QuickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	flagNum := nums[0]
	index := 1
	left, right := 0, len(nums)-1
	for left < right {
		if nums[index] > flagNum {
			nums[index], nums[right] = nums[right], nums[index]
			right--
		} else {
			nums[index], nums[left] = nums[left], nums[index]
			left++
			index++
		}
	}
	nums[left] = flagNum
	QuickSort(nums[:left])
	QuickSort(nums[left+1:])
}

// QuickSort_ 快速排序 ———— 双路快排
func QuickSort_(nums []int) {
	if len(nums) <= 1 {
		return
	}

	flagNum := nums[0]
	left, right := 1, len(nums)-1
	for {
		for left <= right && nums[left] < flagNum {
			left++
		}
		for right > left && nums[right] > flagNum {
			right--
		}
		if left > right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	nums[right], nums[0] = nums[0], nums[right]
	QuickSort_(nums[:right])
	QuickSort_(nums[right+1:])
}

// QuickSort__ 快速排序 ———— 三路快排
func QuickSort__(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 直接选第一个或最后一个位置可能导致无限循环
	// 比如选了第一个，本身就是第一个，那么后续所有 QuickSort(nums[right+1:])
	// 都是多余的排序 且没有改变顺序 不断重试
	flagNum := nums[len(nums)/2]
	index, left, right := 0, 0, len(nums)
	for index < right {
		switch {
		case nums[index] < flagNum:
			nums[index], nums[left] = nums[left], nums[index]
			index++
			left++

		case nums[index] > flagNum:
			right--
			nums[index], nums[right] = nums[right], nums[index]

		default:
			index++
		}
	}
	QuickSort__(nums[:left])
	QuickSort__(nums[right:])
}

// QuickSortFast ———— leetcode 100%
func QuickSortFast(nums []int) {
	var fastSort func(start, end int)
	fastSort = func(l, h int) {
		if l >= h {
			return
		}
		midP := l + (h-l)/2
		mid := nums[midP]
		nums[midP], nums[l] = nums[l], nums[midP]

		start, end := l, h
		for start < end {
			for start < end && nums[end] > mid {
				end--
			}
			if start < end && nums[end] <= mid {
				nums[start], nums[end] = nums[end], nums[start]
			}
			for start < end && nums[start] <= mid {
				start++
			}
			if start < end && nums[start] > mid {
				nums[start], nums[end] = nums[end], nums[start]
			}
		}
		nums[start] = mid
		fastSort(l, start-1)
		fastSort(start+1, h)
	}
	fastSort(0, len(nums)-1)
}

// MergeSort ———— 归并排序
func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := len(nums) / 2
	MergeSort(nums[:mid])
	MergeSort(nums[mid:])

	arrLeft, arrRight := make([]int, mid), make([]int, len(nums)-mid)
	leftRange, rightRange := mid-1, len(arrRight)-1
	copy(arrLeft, nums[:mid])
	copy(arrRight, nums[mid:])
	var indexL, indexR int
	for i := 0; i < len(nums); i++ {
		switch {
		case indexL > leftRange:
			nums[i] = arrRight[indexR]
			indexR++

		case indexR > rightRange:
			nums[i] = arrLeft[indexL]
			indexL++

		case arrLeft[indexL] < arrRight[indexR]:
			nums[i] = arrLeft[indexL]
			indexL++

		default:
			nums[i] = arrRight[indexR]
			indexR++
		}
	}
}

// HeapSort ———— 堆排序
func HeapSort(nums []int) {
	// 新建一个最大堆
	adjustHead(nums)
	for i := len(nums) - 1; i >= 1; i-- {
		// 弹出最大堆得堆顶放在最后
		nums[0], nums[i] = nums[i], nums[0]
		// 重建最大堆
		rebuildHead(nums, 0, i-1)
	}
}

func rebuildHead(nums []int, par, last int) {
	left := 2*par + 1  // 左子节点
	right := 2*par + 2 // 右子节点
	maxIndex := left
	if right <= last && nums[right] > nums[left] {
		// 找到最大子节点
		maxIndex = right
	}
	if left <= last && nums[par] < nums[maxIndex] {
		// 与最大节点对比
		nums[par], nums[maxIndex] = nums[maxIndex], nums[par]
		rebuildHead(nums, maxIndex, last)
	}
}

func adjustHead(nums []int) {
	for i := 0; i < len(nums); i++ {
		par := (i - 1) >> 1 // 找到父节点
		child := i          // 定义子节点
		// 从子节点道根节点构建最大堆
		for child > 0 && nums[par] < nums[child] {
			nums[child], nums[par] = nums[par], nums[child]
			child = par
			par = (par - 1) >> 1
		}
	}
}

func out(nums []int) {
	for _, v := range nums {
		fmt.Print(" ", v)
	}
	fmt.Println()
}
