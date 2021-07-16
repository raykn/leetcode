package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type data struct {
	sum, count int
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
