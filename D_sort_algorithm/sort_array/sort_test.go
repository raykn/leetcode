package main

import (
	"math/rand"
	"testing"
)

const testNum = 1000000

// 冒泡排序
func BenchmarkMaoPaoSort(b *testing.B) {
	resList := randList(testNum)
	for i := 0; i < b.N; i++ {
		MaopaoSort(resList)
	}
}

// 选择排序
func BenchmarkSelectSort(b *testing.B) {
	resList := randList(testNum)
	for i := 0; i < b.N; i++ {
		SelectSort(resList)
	}
}

// 插入排序
func BenchmarkInsertSort(b *testing.B) {
	resList := randList(testNum)
	for i := 0; i < b.N; i++ {
		InsertSort(resList)
	}
}

// 归并排序
func BenchmarkMergeSort(b *testing.B) {
	resList := randList(testNum)
	for i := 0; i < b.N; i++ {
		MergeSort(resList)
	}
}

// 快排
func BenchmarkQuickSort(b *testing.B) {
	resList := randList(testNum)
	for i := 0; i < b.N; i++ {
		QuickSort__(resList)
	}
}

// 堆排
func BenchmarkHeadSort(b *testing.B) {
	resList := randList(testNum)
	for i := 0; i < b.N; i++ {
		HeapSort(resList)
	}
}

func randList(count int) (ret []int) {
	for i := 0; i < count; i++ {
		r := rand.Int()
		ret = append(ret, r)
	}
	return
}
