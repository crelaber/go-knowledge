package datastruct

import (
	"fmt"
	"testing"
)

func TestLinkStack(t *testing.T) {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}
