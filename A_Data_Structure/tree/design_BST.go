package main

// todo  173. 二叉搜索树迭代器
// type BSTIterator struct {
// 	stack []*TreeNode
// 	cur   *TreeNode
// }
//
// func Constructor(root *TreeNode) BSTIterator {
// 	return BSTIterator{cur: root}
// }
//
// func (this *BSTIterator) Next() int {
// todo 迭代遍历
// 	for node := this.cur; node != nil; node = node.Left {
// 		this.stack = append(this.stack, node)
// 	}
// 	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
// 	val := this.cur.Val
// 	this.cur = this.cur.Right
// 	return val
// }
//
// func (this *BSTIterator) HasNext() bool {
// 	return this.cur != nil || len(this.stack) > 0
// }
