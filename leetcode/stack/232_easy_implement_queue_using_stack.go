package stack

//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//实现 MyQueue 类：
//void push(int x) 将元素 x 推到队列的末尾
//int pop() 从队列的开头移除并返回元素
//int peek() 返回队列开头的元素
//说明：
//你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
//你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

type MyQueue struct {
	s1 Stack
	s2 Stack
}

func NewQueue() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.s1.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Pop()
}

func (q *MyQueue) Peek() int {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}
