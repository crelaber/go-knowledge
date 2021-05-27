package datastruct

import "sync"

//双端列表，也可以叫双端队列
//参考地址：http://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6at149ks6s

type DoubleList struct {
	head *DoubleListNode
	tail *DoubleListNode
	len  int
	lock sync.Mutex
}

type DoubleListNode struct {
	pre   *DoubleListNode
	next  *DoubleListNode
	value string
}

func (node *DoubleListNode) GetValue() string {
	return node.value
}

func (node *DoubleListNode) GetPre() *DoubleListNode {
	return node.pre
}

func (node *DoubleListNode) GetNext() *DoubleListNode {
	return node.next
}

func (node *DoubleListNode) HasNext() bool {
	return node.pre == nil
}

func (node *DoubleListNode) HasPre() bool {
	return node.next == nil
}

func (node *DoubleListNode) IsNil() bool {
	return node == nil
}

func (list *DoubleList) AddNodeFromHead(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n > list.len {
		panic("index out of bound")
	}
	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}

	newNode := new(DoubleListNode)
	newNode.value = v

	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		pre := node.pre
		if pre.IsNil() {
			newNode.next = node
			node.pre = newNode
			list.head = newNode
		} else {
			// 将新节点插入到定位到的节点之前
			// 定位到的节点的前驱节点 pre 现在链接到新节点上
			pre.next = newNode
			newNode.pre = pre
			// 定位到的节点的后驱节点 node.next 现在链接到新节点上
			node.next.pre = newNode
			newNode.next = node.next
		}
	}
	list.len = list.len + 1
}

//// 添加节点到链表尾部的第N个元素之后，N=0表示新节点成为新的尾部
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过列表长度，一定找不到，panic
	if n > list.len {
		panic("index out of bound")
	}

	node := list.tail
	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	newNode := new(DoubleListNode)
	newNode.value = v
	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		next := node.next
		// 如果定位到的节点后驱为nil，那么定位到的节点为链表尾部，需要换尾部
		if next.IsNil() {
			node.next = newNode
			newNode.pre = node
			// 新节点成为尾部
			list.tail = newNode
		} else {
			newNode.pre = node
			node.next = newNode

			newNode.next = next
			next.pre = newNode
		}
	}
	list.len = list.len + 1
}
