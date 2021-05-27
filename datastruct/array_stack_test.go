package datastruct

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	as := new(ArrayStack)
	as.Push("cat")
	as.Push("dog")
	as.Push("hen")

	fmt.Println("size :", as.Size())
	fmt.Println("pop : ", as.Pop())
	fmt.Println("pop : ", as.Pop())
	fmt.Println("size  : ", as.Size())
	as.Push("drag")
	fmt.Println("pop : ", as.Pop())

}
