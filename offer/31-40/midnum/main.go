package midnum

import "container/heap"

// 小跟堆大根堆分别储存
type MedianFinder struct {
	maxH *maxHeap
	minH *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	maxH := &maxHeap{}
	minH := &minHeap{}
	heap.Init(maxH)
	heap.Init(minH)
	return MedianFinder{maxH, minH}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxH.Len() == 0 {
		heap.Push(this.minH, num)
	} else {
		maxHPeek := this.maxH.Peek()
		if num >= maxHPeek {
			heap.Push(this.minH, num)
		} else {
			heap.Push(this.maxH, num)
		}
	}

	// 看高度差是否超过2,超过2就平衡下
	if this.minH.Len()-this.maxH.Len() >= 2 {
		temp := heap.Pop(this.minH)
		heap.Push(this.maxH, temp)
	}

	if this.minH.Len()-this.maxH.Len() <= -2 {
		temp := heap.Pop(this.maxH)
		heap.Push(this.minH, temp)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// 高度不相等，多的那个堆顶就是中位数
	if this.minH.Len() > this.maxH.Len() {
		return float64(this.minH.Peek())
	} else if this.minH.Len() < this.maxH.Len() {
		return float64(this.maxH.Peek())
	}
	// 高度相等，中位数是两个堆顶取出除二
	return (float64(this.minH.Peek()) + float64(this.maxH.Peek())) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type maxHeap []int

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *maxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 查看堆顶元素
func (h *maxHeap) Peek() int {
	return (*h)[0]
}

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 查看堆顶元素
func (h *minHeap) Peek() int {
	return (*h)[0]
}
