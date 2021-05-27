package datastruct

import "sync"

type LinkStack struct {
	root *LinkStackNode
	size int
	lock sync.Mutex
}

type LinkStackNode struct {
	Next  *LinkStackNode
	Value string
}

func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.root == nil {
		stack.root = new(LinkStackNode)
		stack.root.Value = v
	} else {
		preNode := stack.root
		newNode := new(LinkStackNode)
		newNode.Value = v
		newNode.Next = preNode
		stack.root = newNode
	}

	stack.size = stack.size + 1
}

//出栈
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		panic("empty")
	}

	topNode := stack.root
	v := topNode.Value

	stack.root = topNode.Next
	stack.size = stack.size - 1
	return v
}

func (stack *LinkStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}
	v := stack.root.Value
	return v
}

func (stack *LinkStack) Size() int {
	return stack.size
}

func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
