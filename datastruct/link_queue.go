package datastruct

import "sync"

type LinkQueue struct {
	root *LinkQueueNode
	size int
	lock sync.Mutex
}

type LinkQueueNode struct {
	Next  *LinkQueueNode
	Value string
}

func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.root == nil {
		queue.root = new(LinkQueueNode)
		queue.root.Value = v
	} else {
		newNode := new(LinkQueueNode)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		nowNode.Next = newNode
	}
	queue.size = queue.size + 1
}

//链表第一个节点出队即可，时间复杂度为：O(1)。
func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}

	topNode := queue.root
	v := topNode.Value
	queue.root = topNode.Next
	queue.size = queue.size - 1
	return v
}
