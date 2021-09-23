package datastruct

type AVLTree struct {
	Root *AVLTreeNode //数的根节点
}

type AVLTreeNode struct {
	Value  int64
	Times  int64        //值出现的次数
	Height int64        //该节点作为数的根节点，树的高度，方便计算平衡因子
	Left   *AVLTreeNode //左子树
	Right  *AVLTreeNode //右子树
}

func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

//更新数的高度
func (node *AVLTreeNode) UpdateTreeHeight() {
	if node == nil {
		return
	}

	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	//高度上自己加一层
	node.Height = maxHeight + 1
}

// 计算平衡因子
func (node *AVLTreeNode) BalanceFactor() int64 {
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}

//右旋操作
func RightRotation(Root *AVLTreeNode) *AVLTreeNode {
	//只有Pivot和B，Root位置变了
	Pivot := Root.Left
	B := Pivot.Right
	Pivot.Right = Root
	Root.Left = B
	// 只有Root和Pivot变化了高度
	Root.UpdateTreeHeight()
	Pivot.UpdateTreeHeight()
	return Pivot
}

//左旋操作
func LeftRotation(Root *AVLTreeNode) *AVLTreeNode {
	Pivot := Root.Right
	B := Pivot.Left
	Pivot.Left = Root
	Root.Right = B
	Root.UpdateTreeHeight()
	Pivot.UpdateTreeHeight()
	return Pivot
}

// 先左后右旋操作
func LeftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

//先右后左旋操作
func RightLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Right = RightRotation(node.Right)
	return LeftRotation(node)
}

func (node *AVLTreeNode) Add(value int64) *AVLTreeNode {
	if node == nil {
		return &AVLTreeNode{Value: value, Height: 1}
	}
	// 如果值重复，什么都不用做，直接更新次数
	if node.Value == value {
		node.Times = node.Times + 1
		return node
	}

	var newTreeNode *AVLTreeNode
	// 插入的值大于节点值，要从右子树继续插入
	if value > node.Value {
		node.Right = node.Right.Add(value)
		// 平衡因子，插入右子树后，要确保树根左子树的高度不能比右子树低一层。
		factor := node.BalanceFactor()
		// 右子树的高度变高了，导致左子树-右子树的高度从-1变成了-2
		if factor == 2 {
			if value > node.Right.Value {
				newTreeNode = LeftRotation(node)
			} else {
				newTreeNode = RightRotation(node)
			}
		}
	} else {
		// 插入的值小于节点值，要从左子树继续插入
		node.Left = node.Left.Add(value)
		// 平衡因子，插入左子树后，要确保树根左子树的高度不能比右子树高一层。
		factor := node.BalanceFactor()
		if factor == 2 {
			if value < node.Left.Value {
				newTreeNode = RightRotation(node)
			} else {
				newTreeNode = LeftRotation(node)
			}
		}
	}

	if newTreeNode == nil {
		// 表示什么旋转都没有，根节点没变，直接刷新树高度
		node.UpdateTreeHeight()
		return node
	}
	// 旋转了，树根节点变了，需要刷新新的树根高度
	newTreeNode.UpdateTreeHeight()
	return newTreeNode
}

func (tree *AVLTree) FindMinValue() *AVLTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *AVLTreeNode) FindMinValue() *AVLTreeNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}
	// 一直左子树递归
	return node.Left.FindMinValue()
}

func (tree *AVLTree) FindMaxValue() *AVLTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()

}

func (node *AVLTreeNode) FindMaxValue() *AVLTreeNode {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}
	return node.Right.FindMaxValue()
}

//// 查找指定节点
func (tree *AVLTree) Find(value int64) *AVLTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(value)
}

func (node *AVLTreeNode) Find(value int64) *AVLTreeNode {
	// 如果该节点刚刚等于该值，那么返回该节点
	if value == node.Value {
		return node
	} else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Left.Find(value)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Right.Find(value)
	}
}

//删除指定节点
func (tree *AVLTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	tree.Root = tree.Root.Delete(value)
}

func (node *AVLTreeNode) Delete(value int64) *AVLTreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		// 从左子树开始删除
		node.Left = node.Left.Delete(value)
		// 删除后要更新该子树高度
		node.Left.UpdateTreeHeight()
	} else if value > node.Value {
		node.Right = node.Right.Delete(value)
		node.Right.UpdateTreeHeight()
	} else {
		// 找到该值对应的节点
		// 该节点没有左右子树
		// 第一种情况，删除的节点没有儿子，直接删除即可。
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// 该节点有两棵子树，选择更高的哪个来替换
		// 第二种情况，删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，
		// 如果左子树更高，选择左子树中最大的节点，也就是左子树最右边的叶子节点，
		// 如果右子树更高，选择右子树中最小的节点，也就是右子树最左边的叶子节点。
		// 最后，删除这个叶子节点。
		if node.Left != nil && node.Right != nil {
			// 左子树更高，拿左子树中最大值的节点替换
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}
				// 最大值的节点替换被删除节点
				node.Value = maxNode.Value
				node.Times = maxNode.Times
				// 把最大的节点删掉
				node.Left = node.Left.Delete(maxNode.Value)
				node.Left.UpdateTreeHeight()
			} else {
				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}
				node.Value = minNode.Value
				node.Times = minNode.Times
				node.Right = node.Right.Delete(minNode.Value)
				node.Right.UpdateTreeHeight()
			}
		}
	}
}
