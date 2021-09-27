package tree

//题目：二叉搜索树中2个节点被错误的交换，请在不改变其结构的情况下，恢复这棵树
//解题思路：
//这题目先按照先跟遍历一次可以找到这两个问题的节点，因为先访问根节点，然后左孩子，右孩子。用先根遍历二叉搜索树的时候，根节点比左子树都要大，根节点比右子树都要小
//所以左子树比根节点大的话就会出现乱序；根节点比右子树大的话，也是出现了乱序。遍历过程中在左子树中如果出现了前一次遍历的节点大于此次根节点的值，则出现了错误节点记录下来。继续遍历，直到找到了第二个这样的节点
//最后交换这两个节点的时候，只是交换他们的值即可

type RecoverTreeNode struct {
	Val   int
	Left  *RecoverTreeNode
	Right *RecoverTreeNode
}

func RecoverTree(root *RecoverTreeNode) {
	var prev, target1, target2 *RecoverTreeNode
	_, target1, target2 = inOrderTraverse(root, prev, target1, target2)
	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}
}

func inOrderTraverse(root, prev, target1, target2 *RecoverTreeNode) (*RecoverTreeNode, *RecoverTreeNode, *RecoverTreeNode) {
	if root == nil {
		return prev, target1, target2
	}
	prev, target1, target2 = inOrderTraverse(root.Left, prev, target1, target2)
	if prev != nil && prev.Val > root.Val {
		if target1 == nil {
			target1 = prev
		}
		target2 = root
	}
	prev = root
	prev, target1, target2 = inOrderTraverse(root.Right, prev, target1, target2)
	return prev, target1, target2
}
