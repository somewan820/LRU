package main

import "fmt"

type LRU struct {
	// 当前容量
	size int
	// 最大容量
	limit int
	// 缓存map
	cache map[string]*DoubleLinkList
	// 头尾节点
	head *DoubleLinkList
	tail *DoubleLinkList
}

type DoubleLinkList struct {
	key   string
	value int
	next  *DoubleLinkList
	prev  *DoubleLinkList
}

func (l *LRU) Get(key string) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}

	node := l.cache[key]
	l.moveNodeToHead(node)

	return node.value
}

func (l *LRU) Put(key string, value int) {
	if node, ok := l.cache[key]; !ok {
		node = &DoubleLinkList{key: key, value: value}
		l.cache[key] = node
		l.addNodeToHead(node)
		l.size++
		if l.size > l.limit {
			tailPre := l.tail.prev
			l.removeNode(tailPre)
			delete(l.cache, tailPre.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveNodeToHead(node)
	}
}

func (l *LRU) moveNodeToHead(node *DoubleLinkList) {
	l.removeNode(node)
	l.addNodeToHead(node)
}

func (l *LRU) addNodeToHead(node *DoubleLinkList) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRU) removeNode(node *DoubleLinkList) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (l *LRU) ListLRU() {
	node := l.head.next
	for node != l.tail {
		fmt.Printf("key: %s, value: %d\n", node.key, node.value)
		node = node.next
	}
}

func NewLRU() *LRU {
	l := &LRU{
		limit: 5,
		cache: make(map[string]*DoubleLinkList),
		head: &DoubleLinkList{
			key:   "head",
			value: 0,
		},
		tail: &DoubleLinkList{
			key:   "tail",
			value: 0,
		},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func main() {
	// example1
	l1 := NewLRU()
	l1.Put("key1", 1)
	l1.Put("key2", 2)
	l1.Put("key3", 3)
	l1.Put("key4", 4)
	l1.Put("key5", 5)
	fmt.Println(l1.Get("key2"))
	l1.ListLRU()

	// example2
	l2 := NewLRU()
	l2.Put("key1", 1)
	l2.Put("key2", 2)
	l2.Put("key3", 3)
	l2.Put("key4", 4)
	l2.Put("key5", 5)
	fmt.Println("before put: ")
	l2.ListLRU()
	l2.Put("key6", 6)
	fmt.Println("after put: ")
	l2.ListLRU()
}
