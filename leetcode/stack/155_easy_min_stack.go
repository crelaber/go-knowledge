package stack

//最小栈
//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//push(x) —— 将元素 x 推入栈中。
//pop() —— 删除栈顶的元素。
//top() —— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。

type MinStack struct {
	min   int
	stack []int
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	if len(ms.stack) == 0 {
		ms.min = x
	}
	if x < ms.min {
		ms.min = x
	}
	ms.stack = append(ms.stack, x)
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	if len(ms.stack) > 0 {
		ms.min = ms.stack[0]
	}
	for _, elem := range ms.stack {
		if ms.min > elem {
			ms.min = elem
		}
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min
}
