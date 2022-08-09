package datastruct

import "fmt"

//二叉查找树，又叫二叉排序树，二叉搜索树，是一种有特定规则的二叉树，定义如下：
//它是一棵二叉树，或者是空树。
//左子树所有节点的值都小于它的根节点，右子树所有节点的值都大于它的根节点。
//左右子树也是一棵二叉查找树。
//参考：https://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6b0vdr7bm2

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	Value int64                 //值
	Times int64                 //值出现的次数
	Left  *BinarySearchTreeNode //左子树
	Right *BinarySearchTreeNode //右子树
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

// Add 添加元素
//一个节点代表一个元素，节点的 Value 值是用来进行二叉查找的关键，当 Value 值重复时，我们将值出现的次数 Times 加 1。添加元素代码如下：
func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{
			Value: value,
		}
		return
	}

	tree.Root.Add(value)
}

// FindMinValue 查找最大之和最小值
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

// FindMaxValue 查找最大值
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

// Find 查找指定元素
func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(value)
}

// MinOrder 中序遍历实现排序
func (tree *BinarySearchTree) MinOrder() {
	tree.Root.MinOrder()
}

// Delete 删除指定的元素
//首先查找到要删除元素的节点：tree.Root.Find(value)，然后找到该节点父亲：tree.Root.FindParent(value)
//四种不同情况对删除节点进行补位。核心在于，第三种情况下，删除的节点有两个子树情况下，需要用右子树中最小的节点来替换被删除节点。
func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	node := tree.Root.Find(value)
	if node == nil {
		//不存在该值，直接返回
		return
	}
	//查找出该值的父节点
	parent := tree.Root.FindParent(value)
	if parent == nil && node.Left == nil && node.Right == nil {
		//置空后直接返回
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		//第二种情况，删除的节点有父亲节点，但是美哦与子树
		//如果删除的节点是父亲的左儿子，直接删除该值
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			//删除的原来是父亲的有儿子，直接将该值删除即可
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		//第三种情况，删除的节点下有两个子树，因为右子树的值比左子树大，那么右子树中最小元素来替换删除的节点
		//右子树的最小元素，只要一直往右子树的左边一直找就可以找到，替换后二叉树的性质又满足了
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		tree.Delete(minNode.Value)
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		//第四种情况，只有一个子树，那么该子树直接替换被删除的节点即可
		//父亲为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
		}

		if node.Left != nil {
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			//如果删除的节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}

// FindParent 查找指定元素的父节点
func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindParent(value)
}

// Add 节点添加元素
//如果添加元素时是棵空树，那么初始化根节点。然后添加的值和根节点比较，判断是要插入到根节点左子树还是右子树，还是不用插入。
//当值比根节点小时，元素要插入到根节点的左子树中，当值比根节点大时，元素要插入到根节点的右子树中，相等时不插入，只更新次数。
//然后再分别对根节点的左子树和右子树进行递归操作即可。
func (node *BinarySearchTreeNode) Add(value int64) {
	if value < node.Value {
		if node.Left == nil { //如果不存在则新建节点
			node.Left = &BinarySearchTreeNode{
				Value: value,
			}
		} else { //负责循环添加
			node.Left.Add(value)
		}
	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{
				Value: value,
			}
		} else {
			node.Right.Add(value)
		}
	} else {
		//值如果相同，则不需要添加，直接将次数加一
		node.Times = node.Times + 1
	}
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	if node.Left == nil {
		return node
	}
	//一直递归左子树
	return node.Left.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.FindMaxValue()
}

func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if value == node.Value {
		return node
	} else if value < node.Value {
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(value)
	} else {
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(value)
	}
}

// FindParent 寻找指定元素的父节点
func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	if value < node.Value {
		leftTree := node.Left
		if leftTree == nil {
			return nil
		}
		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value)
		}
	} else {
		rightTree := node.Right
		if rightTree == nil {
			return nil
		}
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}
}

func (node *BinarySearchTreeNode) MinOrder() {
	if node == nil {
		return
	}
	node.Left.MinOrder()
	//按照次数打印根节点
	for i := 0; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}
	node.Right.MinOrder()
}
