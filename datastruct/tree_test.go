package datastruct

import (
	"fmt"
	"testing"
)

var inst = new(TreeInst)

func TestTreeOrder(t *testing.T) {
	root := &TreeNode{
		Data: "A",
	}

	root.Left = &TreeNode{
		Data: "B",
	}
	root.Right = &TreeNode{
		Data: "C",
	}

	root.Left.Left = &TreeNode{
		Data: "D",
	}

	root.Left.Right = &TreeNode{
		Data: "E",
	}

	root.Right.Left = &TreeNode{
		Data: "F",
	}
	fmt.Println("\n先序遍历：")
	inst.PreOrder(root)
	fmt.Println("\n中序遍历：")
	inst.MidOrder(root)
	fmt.Println("\n后序遍历：")
	inst.PostOrder(root)
	fmt.Println("\n层次遍历：")
	inst.LayOrder(root)

}
