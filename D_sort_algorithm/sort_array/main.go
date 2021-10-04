package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	n       = 10
	numRand = 100
)

func main() {
	rand.Seed(time.Now().Unix())
	maopaoSort(randNums())
	selectSort(randNums())
	insertSort(randNums())
	insertSort_(randNums())
	quickSort(randNums())
	quickSort_(randNums())
	quickSort__(randNums())
	quickSort___(randNums())
	mergeSort(randNums())
	heapSort(randNums())
}

func randNums() []int {
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(numRand))
	}
	return nums
}

// 冒泡
func maopaoSort(nums []int) {
	fmt.Println("———— maopaoSort ————")
	fmt.Println(nums)
	MaopaoSort(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 选择
func selectSort(nums []int) {
	fmt.Println("———— selectSort ————")
	fmt.Println(nums)
	SelectSort(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 插入
func insertSort(nums []int) {
	fmt.Println("———— insertSort ————")
	fmt.Println(nums)
	InsertSort(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 插入2
func insertSort_(nums []int) {
	fmt.Println("———— insertSort_ ————")
	fmt.Println(nums)
	InsertSort_(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 快排
func quickSort(nums []int) {
	fmt.Println("———— quickSort ————")
	fmt.Println(nums)
	QuickSort(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 快排2
func quickSort_(nums []int) {
	fmt.Println("———— quickSort_ ————")
	fmt.Println(nums)
	QuickSort_(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 快排3
func quickSort__(nums []int) {
	fmt.Println("———— quickSort__ ————")
	fmt.Println(nums)
	QuickSort__(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 快排4
func quickSort___(nums []int) {
	fmt.Println("———— quickSort___ leetcode ————")
	fmt.Println(nums)
	QuickSortFast(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 归并
func mergeSort(nums []int) {
	fmt.Println("———— mergeSort ————")
	fmt.Println(nums)
	MergeSort(nums)
	fmt.Println(nums)
	fmt.Println()
}

// 堆排
func heapSort(nums []int) {
	fmt.Println("———— heapSort ————")
	fmt.Println(nums)
	HeapSort(nums)
	fmt.Println(nums)
	fmt.Println()
}
