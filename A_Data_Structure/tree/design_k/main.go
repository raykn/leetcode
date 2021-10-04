package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type KthLargest struct {
	root   *TreeNode
	num    int
	k      int
	target int
}

func Constructor(k int, nums []int) KthLargest {
	this := new(KthLargest)
	this.buildTree(nums)
	this.k = k
	return *this
}

func (this *KthLargest) Add(val int) int {
	this.insertNode(val)
	return this.getK()
}

func (this *KthLargest) getK() int {
	return this.target
}

func (k *KthLargest) buildTree(nums []int) {
	for _, val := range nums {
		k.insertNode(val)
	}
}

func (k *KthLargest) insertNode(val int) {
	k.num++
	if k.root == nil {
		k.root = &TreeNode{Val: val}
		return
	}
	var (
		node  = k.root
		count = k.num - k.k + 1
	)
	for {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				break
			} else {
				node = node.Left
			}

		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				break
			} else {
				node = node.Right
			}

		}
	}

	k.searchK(k.root, count)
}

func (k *KthLargest) searchK(root *TreeNode, count int) {
	if root == nil || count <= 0 {
		return
	}
	fmt.Println(count, " ", root.Val)
	k.searchK(root.Left, count)
	count--
	if count == 0 {
		k.target = root.Val
	}
	k.searchK(root.Right, count)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
func main() {
	k := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(k.Add(3))
	fmt.Println(k.Add(5))
	fmt.Println(k.Add(10))
	fmt.Println(k.Add(9))
	fmt.Println(k.Add(4))
}
