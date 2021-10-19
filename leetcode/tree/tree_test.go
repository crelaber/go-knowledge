package tree

import (
	"fmt"
	"testing"
)

func TestGenerateTree(t *testing.T) {
	trees := generateTrees(20)
	for i, _ := range trees {
		result := BinaryTreeInOrderTraverse(trees[i])
		fmt.Println(result)
	}
}
