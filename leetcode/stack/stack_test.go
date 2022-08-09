package stack

import (
	"fmt"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	path := "/path/to1/"
	path = simplifyPath(path)
	fmt.Println(path)
	path = "/path/./to2/"
	path = simplifyPath(path)
	fmt.Println(path)
	path = "/path/../to3/"
	path = simplifyPath(path)
	fmt.Println(path)
}

//逆波兰表达式
func TestEvaluatePolishNotation(t *testing.T) {
	token := []string{"2", "1", "+", "3", "*"}
	val := evalRPN(token)
	fmt.Println(val)
	token = []string{"4", "13", "5", "/", "+"}
	val = evalRPN(token)
	fmt.Println(val)
	token = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	val = evalRPN(token)
	fmt.Println(val)
}

//验证最小栈
func TestMinStack(t *testing.T) {
	s := MinStack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.GetMin())
	println(s.Top())
}

//队列实现栈
func TestQueueStack(t *testing.T) {
	s := MyStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

//栈实现队列
func TestStackQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

func TestDecodeString(t *testing.T) {
	str := "3[a]2[b]4[F]3[c]"
	newStr := decodeString(str)
	fmt.Println(newStr)
}

func TestFindNextGreaterNum(t *testing.T) {
	findNums := []int{4, 1, 2}
	nums := []int{1, 3, 4, 2}
	ret := bestNextGreaterElement(findNums, nums)
	fmt.Println(ret)
}

func TestBaseballCalPoint(t *testing.T) {
	s := []string{"5", "2", "C", "D", "+"}
	fmt.Println(baseballCalPoints(s))
	s = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(baseballCalPoints(s))
}

func TestDailyTemperature(t *testing.T) {
	s := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(s))
}

func TestBackspaceCompare(t *testing.T) {
	s := "ab#c"
	t1 := "cd#a"
	fmt.Println(backspaceCompare(s, t1))
}
