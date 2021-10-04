package minStack

// TODO 30-包含min函数的栈

type MinStack struct {
	stack    []int
	minStack []int
	index    int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
		index:    -1,
	}
}

func (s *MinStack) Push(x int) {
	if s.index == -1 {
		s.minStack = append(s.minStack, x)
	} else {
		s.minStack = append(s.minStack, getMinValue(s.Min(), x))
	}
	s.stack = append(s.stack, x)
	s.index++
}

func (s *MinStack) Pop() {
	s.minStack = s.minStack[:len(s.minStack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.index--
}

func (s *MinStack) Top() int {
	return s.stack[s.index]
}

func (s *MinStack) Min() int {
	return s.minStack[s.index]
}

func getMinValue(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
