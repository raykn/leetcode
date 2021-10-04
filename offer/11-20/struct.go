package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	l := 0
	r := len(matrix[0]) - 1
	u := 0
	b := len(matrix) - 1
	flag := 1
	i, j := 0, 0
	for l <= r && u <= b {
		switch flag {
		case 1:
			for i = l; i <= r; i++ {
				res = append(res, matrix[j][i])
			}
			i--
			u++
		case 2:
			for j = u; j <= b; j++ {
				res = append(res, matrix[j][i])
			}
			j--
			r--
		case 3:
			for i = r; i >= l; i-- {
				res = append(res, matrix[j][i])
			}
			i++
			b--
		case 4:
			for j = b; j >= u; j-- {
				res = append(res, matrix[j][i])
			}
			j++
			l++
		}

		flag++
		if flag == 5 {
			flag = 1
		}
	}
	return res
}
