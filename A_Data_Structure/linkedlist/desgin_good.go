package main

// type Node struct {
// 	Prev *Node
// 	Next *Node
// 	Val  int
// }
//
// type MyLinkedList struct {
// 	Head *Node
// 	Tail *Node
// 	Size int
// }
//
// /** Initialize your data structure here. */
// func Constructor() MyLinkedList {
// 	node := &Node{}
// 	head, tail := node, node
// 	head.Next = tail
// 	tail.Prev = head
// 	return MyLinkedList{
// 		Head: head,
// 		Tail: tail,
// 		Size: 0,
// 	}
// }
//
// /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
// func (this *MyLinkedList) Get(index int) int {
// 	if index < 0 || index > this.Size-1 {
// 		return -1
// 	}
// 	curl := this.Head
// 	if index <= this.Size-index {
// 		for i := 0; i <= index; i++ {
// 			curl = curl.Next
// 		}
// 	} else {
// 		curl = this.Tail
// 		for i := 0; i < this.Size-index; i++ {
// 			curl = curl.Prev
// 		}
// 	}
// 	return curl.Val
// }
//
// /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
//
// func (this *MyLinkedList) AddAtHead(val int) {
// 	this.Size++
// 	node := &Node{Val: val}
// 	firstNode := this.Head.Next
//
// 	this.Head.Next = node
// 	node.Next = firstNode
// 	node.Prev = this.Head
// 	firstNode.Prev = node
// }
//
// /** Append a node of value val to the last element of the linked list. */
// func (this *MyLinkedList) AddAtTail(val int) {
// 	this.Size++
// 	node := &Node{Val: val}
// 	endNode := this.Tail.Prev
//
// 	endNode.Next = node
// 	node.Prev = endNode
// 	node.Next = this.Tail
// 	this.Tail.Prev = node
// }
//
// /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
// func (this *MyLinkedList) AddAtIndex(index int, val int) {
// 	if index > this.Size {
// 		return
// 	} else if index == this.Size {
// 		this.AddAtTail(val)
// 	} else if index <= 0 {
// 		this.AddAtHead(val)
// 	} else {
// 		curl := this.Head
// 		for i := 0; i < index; i++ {
// 			curl = curl.Next
// 		}
// 		node := &Node{Val: val}
// 		curlNext := curl.Next
//
// 		curl.Next = node
// 		node.Prev = curl
// 		node.Next = curlNext
// 		curlNext.Prev = node
// 		this.Size++
// 	}
// 	return
// }
//
// /** Delete the index-th node in the linked list, if the index is valid. */
// func (this *MyLinkedList) DeleteAtIndex(index int) {
// 	if index < 0 || index > this.Size-1 {
// 		return
// 	}
// 	curl := this.Head
// 	for i := 0; i <= index; i++ {
// 		curl = curl.Next
// 	}
// 	prev := curl.Prev
// 	next := curl.Next
// 	prev.Next = next
// 	next.Prev = prev
// 	this.Size--
// }
