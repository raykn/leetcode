package queue

// TODO 59 - II. 队列的最大值
type MaxQueue struct {
	queue   []int
	maxList []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:   make([]int, 0),
		maxList: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxList) == 0 {
		return -1
	}
	return this.maxList[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.maxList) != 0 && value > this.maxList[len(this.maxList)-1] {
		this.maxList = this.maxList[:len(this.maxList)-1]
	}
	this.maxList = append(this.maxList, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	if this.maxList[0] == val {
		this.maxList = this.maxList[1:]
	}
	return val
}
