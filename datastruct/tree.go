package datastruct

import (
	"fmt"
	"sync"
)

type TreeInst struct {
}

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
//先序遍历：先访问根节点，再访问左子树，最后访问右子树。
func (t TreeInst) PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Print(tree.Data, " ")
	t.PreOrder(tree.Left)
	t.PreOrder(tree.Right)
}

//中序遍历
//中序遍历：先访问左子树，再访问根节点，最后访问右子树。
func (t TreeInst) MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	t.MidOrder(tree.Left)
	fmt.Print(tree.Data, " ")
	t.MidOrder(tree.Right)
}

//后续遍历
//后序遍历：先访问左子树，再访问右子树，最后访问根节点。
func (t TreeInst) PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	t.MidOrder(tree.Left)
	t.MidOrder(tree.Right)
	fmt.Print(tree.Data, " ")
}

type TreeLinkNode struct {
	Next  *TreeLinkNode
	Value *TreeNode
}

type TreeLinkQueue struct {
	root *TreeLinkNode
	size int
	lock sync.Mutex
}

/**
 * 层次遍历较复杂，用到一种名叫广度遍历的方法，需要使用辅助的先进先出的队列。
 * 先将树的根节点放入队列。
 * 从队列里面 remove 出节点，先打印节点值，如果该节点有左子树节点，左子树入栈，如果有右子树节点，右子树入栈。
 * 重复2，直到队列里面没有元素。
 */
func (t TreeInst) LayOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	queue := new(TreeLinkQueue)
	queue.Add(tree)
	for queue.size > 0 {
		element := queue.Remove()
		fmt.Print(element.Data, " ")
		// 左子树非空，入队列
		if element.Left != nil {
			queue.Add(element.Left)
		}
		// 右子树非空，入队列

		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

//入队
func (queue *TreeLinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	// 如果栈顶为空，那么增加节点
	if queue.root == nil {
		queue.root = new(TreeLinkNode)
		queue.root.Value = v
	} else {
		newNode := new(TreeLinkNode)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		nowNode.Next = newNode
	}
	queue.size = queue.size + 1
}

//出队
func (queue *TreeLinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 {
		panic("over limit")
	}
	topNode := queue.root
	v := topNode.Value
	queue.root = topNode.Next
	queue.size = queue.size - 1
	return v
}

func (queue *TreeLinkQueue) Size() int {
	return queue.size
}
