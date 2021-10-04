package trie_1_rzk

type node struct {
	val    byte
	child  []*node
	high   int
	isOver bool
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{
		root: &node{},
	}
}

func (this *Trie) Insert(word string) {
	paths := []byte(word)
	this.insert(this.root, paths, 0)
}

func (this *Trie) insert(root *node, path []byte, high int) {
	if high == len(path) {
		root.isOver = true
		return
	}
	var tmp *node
	for _, v := range root.child {
		if v.val == path[high] {
			tmp = v
			break
		}
	}
	if tmp == nil {
		tmp = &node{
			val:  path[high],
			high: high + 1,
		}
		root.child = append(root.child, tmp)
	}

	this.insert(tmp, path, high+1)
}

func (this *Trie) Search(word string) bool {

	if len(word) == 0 {
		return false
	}

	start := this.root
	index := 0
	path := []byte(word)
	flag := false
	for {

		if index == len(path) {
			return start.isOver
		}

		for _, v := range start.child {
			if path[index] == v.val {
				flag = true
				index++
				start = v
				break
			}
		}
		if !flag {
			return false
		}
		flag = false
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	start := this.root
	index := 0
	path := []byte(prefix)
	flag := false
	for {
		if index == len(path) {
			return true
		}
		for _, v := range start.child {
			if path[index] == v.val {
				flag = true
				index++
				start = v
				break
			}
		}
		if !flag {
			return false
		}
		flag = false
	}
}
