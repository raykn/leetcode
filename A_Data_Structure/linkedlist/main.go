package main

// todo 206. 单链表反转
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		curr := head.Next
		head.Next = pre
		pre = head
		head = curr
	}
	return pre
}

// todo 206. 单链表反转 —— 递归
func reverseList_(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList_(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// todo 206. 单链表反转 —— 最骚写法
func reverseList__(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		// 任意修改赋值顺序都可以
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

// todo 203. 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	preHead := &ListNode{
		Next: head,
	}
	curr := preHead
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return preHead.Next
}

// todo 203. 移除链表元素 —— 递归
func removeElements_(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements_(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

// todo 203. 移除链表元素 —— 最优
func removeElements__(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	ptr := head
	for ptr.Next != nil {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	return head
}

// todo 328. 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	j, o, curr := head, head.Next, head.Next.Next
	for curr != nil {
		tmp := curr.Next
		o.Next = curr.Next
		curr.Next = j.Next
		j.Next = curr

		j = j.Next
		o = o.Next
		if tmp == nil {
			break
		} else {
			curr = tmp.Next
		}
	}
	return head
}

// todo 328. 奇偶链表 —— 优化
func oddEvenList_(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	j := head
	o := head.Next
	OHead := o
	for o != nil && o.Next != nil {
		j.Next = o.Next
		j = j.Next
		o.Next = j.Next
		o = o.Next
	}
	j.Next = OHead
	return head
}

// todo 328. 奇偶链表 —— 最优
func oddEvenList__(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var (
		curr  = head
		jLast *ListNode
		o     = head.Next
		flag  = false
	)
	for curr != nil {
		flag = !flag
		if flag {
			jLast = curr
		}
		if curr.Next == nil {
			break
		}
		next := curr.Next
		curr.Next = curr.Next.Next
		curr = next
	}
	if jLast != nil {
		jLast.Next = o
	}
	return head
}

// todo 234. 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	last := 0
	tmp := head
	for tmp != nil {
		last++
		tmp = tmp.Next
	}
	ji := last%2 > 0
	mid := last/2 + 1
	if ji {
		mid++
	}

	i := 0
	curr := head
	for curr != nil {
		i++
		if mid == i {
			break
		}
		curr = curr.Next
	}
	curr = revert(curr)
	for curr != nil && head != nil {
		if curr.Val != head.Val {
			return false
		}
		curr = curr.Next
		head = head.Next
	}
	return true
}
func revert(head *ListNode) *ListNode {
	var (
		pre  *ListNode
		curr = head
	)
	for curr != nil {
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre
}

// todo 234. 回文链表 —— 优化 步长找mid
func isPalindrome_(head *ListNode) bool {
	one, two := head, head
	for two != nil && two.Next != nil {
		two = two.Next.Next
		one = one.Next
	}
	if two != nil {
		one = one.Next
	}

	var tmp *ListNode
	var oneReverHead *ListNode
	curr := one
	for curr != nil {
		tmp = curr.Next
		curr.Next = oneReverHead
		oneReverHead = curr
		curr = tmp
	}

	otherHead := head
	for otherHead != nil && oneReverHead != nil {
		if otherHead.Val != oneReverHead.Val {
			return false
		}
		otherHead = otherHead.Next
		oneReverHead = oneReverHead.Next
	}
	return true
}

// todo 234. 回文链表 —— 优化 简化+步长找mid
func isPalindrome__(head *ListNode) bool {
	one, two := head, head
	for two.Next != nil && two.Next.Next != nil {
		one = one.Next
		two = two.Next.Next
	}

	var pre *ListNode
	cur := one.Next
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

// todo 21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l3 := ListNode{}
	l4 := &l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l4.Next = l1
			l1 = l1.Next
		} else {
			l4.Next = l2
			l2 = l2.Next
		}
		l4 = l4.Next
	}
	if l1 != nil {
		l4.Next = l1
	}
	if l2 != nil {
		l4.Next = l2
	}
	return l3.Next
}

// todo 2. 两数相加 —— 优化
func addTwoNumbers_(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	var flag int
	for h := newHead; l1 != nil || l2 != nil || flag != 0; {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + flag
		h.Val = sum % 10
		flag = sum / 10
		if l1 != nil || l2 != nil || flag != 0 {
			h.Next = new(ListNode)
			h = h.Next
		}
	}
	return newHead
}

// todo 141. 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	one := head
	two := head.Next
	for two != one {
		if two == nil || two.Next == nil {
			return false
		} else {
			one = one.Next
			two = two.Next.Next
		}
	}
	return true
}

// todo 142. 环形链表 II
// 环外为 a, slow进环后走 b 遇到 fast
// 而此时 fast 走了 a + n(b + c) + b = a + (n + 1)b + nc
// fast 走的路程比 slow 多一倍
// a + (n + 1)b + nc = 2(a + b) ====> a = c + (n - 1)(b + c)
// 得出从 【head到入环点】 == 【n - 1圈的环长】
// 【当 n = 1，则a = c, head到入环 = 相遇点走到入环】
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pQucik := head
	pSlow := head
	for pQucik.Next != nil {
		pQucik = pQucik.Next.Next
		pSlow = pSlow.Next
		if pQucik == nil || pSlow == nil {
			return nil
		}
		if pQucik == pSlow {
			p := head
			for {
				if p == pSlow {
					return p
				}
				p = p.Next
				pSlow = pSlow.Next
			}
		}
	}
	return nil
}

// todo 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 0 || head == nil {
		return head
	}
	last := 0
	cur := head
	for cur != nil {
		last++
		cur = cur.Next
	}
	last -= n
	i := 0
	if last == 0 {
		return head.Next
	}

	var pre *ListNode
	cur = head
	for cur != nil {
		if i == last {
			break
		}
		i, pre, cur = i+1, cur, cur.Next
	}
	if pre != nil && cur != nil {
		pre.Next = cur.Next
	}
	return head
}

// todo 160. 相交链表 —— 双指针做法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// todo 61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	last := 0
	cur := head
	for cur != nil {
		last++
		cur = cur.Next
	}
	k = k % last
	index := last - k
	if k == 0 {
		return head
	}

	var pre *ListNode
	i := 0
	cur = head
	for cur != nil {
		if i == index {
			break
		}
		i++
		pre = cur
		cur = cur.Next
	}
	if pre != nil {
		pre.Next = nil
	}

	answer := cur
	if cur != nil {
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = head
	}
	return answer
}

// todo 61. 旋转链表 —— 最优
// 举例: 12345  slice: 12345 2345 345 45 5
// 先让head成环，123451234512345...., 然后用 45 拼 head, 也就是 4512345, 然后截断 3 的尾巴
func rotateRight_(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	nodeSlice := []*ListNode{}
	for {
		nodeSlice = append(nodeSlice, tmp)
		if tmp.Next != nil {
			tmp = tmp.Next
		} else {
			break
		}
	}
	length := len(nodeSlice)
	t := k % length
	if t == 0 {
		return head
	}
	start := nodeSlice[length-t]
	tmp.Next = head // 成环
	nodeSlice[length-t-1].Next = nil
	return start
}

// todo 430. 扁平化多级双向链表
// func flatten(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}
// 	cur := root
// 	for cur != nil {
// 		if cur.Child != nil {
// 			tmp := cur.Next
// 			childNode := flatten(cur.Child)
// 			lastChild := childNode
// 			for lastChild != nil && lastChild.Next != nil {
// 				lastChild = lastChild.Next
// 			}
//
// 			cur.Next = childNode
// 			childNode.Prev = cur
// 			childNode.Next = tmp
// 			if tmp != nil {
// 				tmp.Prev = cur.Next
// 			}
// 			cur = tmp
//
// 		} else {
// 			cur = cur.Next
// 		}
// 	}
// 	return root
// }

// todo 430. 扁平化多级双向链表 —————— 错误答案 我也不知道为什么
func flatten_(root *Node) *Node {
	if root == nil {
		return root
	}
	var pre *Node
	var flag = false
	var f func(r *Node)
	f = func(r *Node) {
		if r == nil {
			return
		}
		var tmp *Node
		if r.Child != nil {
			tmp = r.Next
			r.Next = r.Child
			r.Child.Prev = r
			r.Child = nil
		}

		if flag {
			flag = !flag
			pre.Next = r
			r.Prev = pre
		}
		if r.Next == nil {
			flag = !flag
			pre = r
		}

		f(r.Child)
		f(r.Next)
		f(tmp)
	}
	f(root)
	root.Prev = nil
	return root
}

// todo 430. 扁平化多级双向链表 —————— 我人都傻了
func flatten__(root *Node) *Node {
	if root == nil {
		return root
	}
	res := &Node{
		Next: root,
	}
	flattenDFS(res, root)
	res.Next.Prev = nil
	return res.Next
}
func flattenDFS(prev, curr *Node) *Node {
	if curr == nil {
		return prev
	}
	curr.Prev = prev
	prev.Next = curr

	tmp := curr.Next
	tail := flattenDFS(curr, curr.Child)
	curr.Child = nil
	return flattenDFS(tail, tmp)
}

// todo 138. 复制带随机指针的链表 ———— DFS 带随机的节点 深拷贝后不能重复创建
func getCloneNode(visitedMap map[*Node]*Node, p *Node) *Node {
	if v, ok := visitedMap[p]; ok {
		return v
	} else {
		var v *Node
		if p != nil {
			v = &Node{Val: p.Val}
		}
		visitedMap[p] = v
		return v
	}
}
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	vistedMap := make(map[*Node]*Node)
	oldNode := head
	newNode := getCloneNode(vistedMap, oldNode)
	newHead := newNode
	for oldNode != nil {
		newNode.Next = getCloneNode(vistedMap, oldNode.Next)
		newNode.Random = getCloneNode(vistedMap, oldNode.Random)
		oldNode = oldNode.Next
		newNode = newNode.Next
	}
	return newHead
}

// todo 138. 复制带随机指针的链表 ———— MS
func flatten___(root *Node) *Node {
	cur := root
	for cur != nil {
		if cur.Child != nil {
			cur = flat(cur)
		}
		cur = cur.Next
	}
	return root
}
func flat(head *Node) *Node {
	// 将子链表拼到父链表后面
	cur := head
	tmpNext := cur.Next
	cur.Next = cur.Child
	cur.Child.Prev = cur
	cur.Child = nil
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = tmpNext
	if tmpNext != nil {
		tmpNext.Prev = cur
	}
	return head
}

// todo 138. 复制带随机指针的链表 ———— 最优解
func flatten____(root *Node) *Node {
	current := root
	for current != nil {
		if current.Child == nil {
			current = current.Next
		} else {
			childStart, childEnd := current.Child, current.Child
			for childEnd.Next != nil {
				childEnd = childEnd.Next
			}

			next := current.Next
			current.Next = childStart
			childStart.Prev = current
			childEnd.Next = next
			if next != nil {
				next.Prev = childEnd
			}

			current.Child = nil

		}
	}
	return root
}

// todo 138. 复制带随机指针的链表 ———— 迭代解法
func copyRandomList_(head *Node) *Node {
	helper := map[*Node]*Node{}
	dummyhead := &Node{Val: 0, Next: head}
	cur, pre := head, dummyhead
	for cur != nil {
		t := &Node{Val: cur.Val}
		pre.Next = t
		pre = t
		helper[cur] = t
		cur = cur.Next
	}
	cur = head
	pre = dummyhead.Next
	for cur != nil {
		pre.Random = helper[cur.Random]
		cur = cur.Next
		pre = pre.Next
	}
	return dummyhead.Next
}
func main() {
	// obj := Constructor()
	// obj.AddAtHead(7)
	// obj.p()
	// obj.AddAtHead(2)
	// obj.p()
	// obj.AddAtHead(1)
	// obj.p()
	// obj.AddAtIndex(3, 0)
	// obj.p()
	// obj.DeleteAtIndex(2)
	// obj.p()
	// obj.AddAtHead(6)
	// obj.p()
	// obj.AddAtTail(4)
	// obj.p()
	// obj.Get(4)
	// obj.p()
	// obj.AddAtHead(4)
	// obj.p()
	// obj.AddAtIndex(5, 0)
	// obj.p()
	// obj.AddAtHead(6)
	// obj.p()

	// head := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val: 5,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// rotateRight_(head, 2)

	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}

	n1.Next = n2
	n2.Prev = n1
	n2.Next = n3
	n3.Prev = n2
	n3.Next = n4
	n4.Next = n5
	n4.Prev = n3
	n5.Next = n6
	n5.Prev = n4
	n6.Next = nil
	n6.Prev = n5

	n3.Child = n7
	n7.Next = n8
	// n7.Prev = n3
	n8.Next = n9
	n8.Prev = n7
	n9.Next = n10
	n10.Prev = n9

	n8.Child = n11
	n11.Next = n12
	// n11.Prev = n8
	n12.Prev = n11
	// flatten(n1)
}
