package tree

import (
	"fmt"
	"testing"
)

func TestGenerateTree(t *testing.T) {
	trees := generateTrees(10)
	for _, tree := range trees {
		result := BinaryTreeInOrderTraverse(tree)
		fmt.Println(result)
	}
}
