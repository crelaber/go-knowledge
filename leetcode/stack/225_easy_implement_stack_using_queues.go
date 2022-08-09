package stack

//用队列实现栈
//请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通队列的全部四种操作（push、top、pop 和 empty）。

//实现 MyStack 类：
//void push(int x) 将元素 x 压入栈顶。
//int pop() 移除并返回栈顶元素。
//int top() 返回栈顶元素。
//boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
//注意：
//你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
//你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

type MyStack struct {
	enque []int
	deque []int
}

func NewStack() MyStack {
	return MyStack{[]int{}, []int{}}
}

func (s *MyStack) Push(x int) {
	s.enque = append(s.enque, x)
}

func (s *MyStack) Pop() int {
	length := len(s.enque)
	for i := 0; i < length-1; i++ {
		s.deque = append(s.deque, s.enque[0])
		s.enque = s.enque[1:]
	}
	topEle := s.enque[0]
	s.enque = s.deque
	s.deque = nil
	return topEle
}

func (s *MyStack) Top() int {
	topEle := s.Pop()
	s.enque = append(s.enque, topEle)
	return topEle
}

func (s *MyStack) Empty() bool {
	if len(s.enque) == 0 {
		return true
	}
	return false
}
