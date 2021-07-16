package main

import "fmt"

//
// type MyLinkedList struct {
// 	head *ListNode
// 	tail *ListNode
// 	len  int
// }
//
// /** Initialize your data structure here. */
// func Constructor() MyLinkedList {
// 	return MyLinkedList{}
// }
//
// /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
// func (this *MyLinkedList) Get(index int) int {
// 	if index < 0 {
// 		return -1
// 	}
// 	i := 0
// 	cur := this.head
// 	for cur != nil {
// 		if i == index {
// 			return cur.Val
// 		}
// 		i++
// 		cur = cur.Next
// 	}
// 	return -1
// }
//
// /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
// func (this *MyLinkedList) AddAtHead(val int) {
// 	newHead := &ListNode{
// 		Val:  val,
// 		Next: this.head,
// 	}
// 	this.head = newHead
// 	if this.len == 0 {
// 		this.tail = newHead
// 	}
// 	this.len++
// }
//
// /** Append a node of value val to the last element of the linked list. */
// func (this *MyLinkedList) AddAtTail(val int) {
// 	newTail := &ListNode{Val: val}
// 	this.len++
// 	if this.tail == nil {
// 		this.head = newTail
// 		this.tail = newTail
// 		return
// 	}
// 	this.tail.Next = newTail
// 	this.tail = newTail
// }
//
// /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
// func (this *MyLinkedList) AddAtIndex(index int, val int) {
// 	if index < 0 || index > this.len {
// 		return
// 	}
// 	insertNode := &ListNode{Val: val}
//
// 	if index == 0 {
// 		this.AddAtHead(val)
// 		return
// 	}
// 	if index == this.len {
// 		this.AddAtTail(val)
// 		return
// 	}
// 	this.len++
// 	cur := this.head
// 	i := 0
// 	index--
// 	for cur != nil {
// 		if index == i {
// 			insertNode.Next = cur.Next
// 			cur.Next = insertNode
// 			return
// 		}
// 		i++
// 		cur = cur.Next
// 	}
// }
//
// /** Delete the index-th node in the linked list, if the index is valid. */
// func (this *MyLinkedList) DeleteAtIndex(index int) {
// 	if index < 0 || this.len == 0 || index > this.len {
// 		return
// 	}
//
// 	if index == 0 {
// 		this.head = this.head.Next
// 		this.len--
// 		return
// 	}
//
// 	i := 0
// 	index--
// 	cur := this.head
// 	for cur.Next != nil {
// 		if i == index {
// 			this.len--
// 			cur.Next = cur.Next.Next
// 			if i == this.len-1 {
// 				this.tail = cur
// 			}
// 			return
// 		}
// 		i++
// 		cur = cur.Next
// 	}
// }

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	newNode := &ListNode{}
	head, tail := newNode, newNode
	head.Next = tail
	tail.Prev = head
	return MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size-1 {
		return -1
	}
	cur := this.head
	if index <= this.size-index {
		for i := 0; i <= index; i++ {
			cur = cur.Next
		}

	} else {
		cur = this.tail
		for i := 0; i < this.size-index; i++ {
			cur = cur.Prev
		}
	}
	return cur.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.size++
	newHead := &ListNode{Val: val}
	firstNode := this.head.Next

	this.head.Next = newHead
	newHead.Next = firstNode
	newHead.Prev = this.head
	firstNode.Prev = newHead
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.size++
	node := &ListNode{Val: val}
	endNode := this.tail.Prev

	endNode.Next = node
	node.Prev = endNode
	node.Next = this.tail
	this.tail.Prev = node
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index > this.size:
		return

	case index == this.size:
		this.AddAtTail(val)

	case index <= 0:
		this.AddAtHead(val)

	default:
		cur := this.head
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		node := &ListNode{Val: val}
		curNext := cur.Next

		cur.Next = node
		node.Prev = cur
		node.Next = curNext
		curNext.Prev = node
		this.size++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size-1 {
		return
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	prev := cur.Prev
	next := cur.Next
	prev.Next = next
	next.Prev = prev
	this.size--
}

func (this *MyLinkedList) p() {
	cur := this.head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("---------  %d ", this.size)
	if this.tail != nil {
		fmt.Print(" %d", this.tail.Val)
	}
	fmt.Println()
}
