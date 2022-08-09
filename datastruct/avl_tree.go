package datastruct

//参考：https://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6b14vcdn63
//AVL树是一棵严格自平衡的二叉查找树，1962年，发明者 Adelson-Velsky 和 Landis 发表了论文，以两个作者的名字命名了该数据结构，这是较早发明的平衡二叉树。
//定义如下：
//1、首先它是一棵二叉查找树。
//2、任意一个节点的左右子树最大高度差为1。
//由于树特征定义，我们可以计算出其高度 h 的上界 h<=1.44log(n)，也就是最坏情况下，树的高度约等于 1.44log(n)。
//假设高度 h 的AVL树最少有 f(h) 个节点，因为左右子树的高度差不能大于1，所以左子树和右子树最少节点为： f(h-1)，f(h-2)。
//因此，树根节点加上左右子树的节点，满足公式 f(h) = 1 + f(h-1) + f(h-2)，初始条件 f(0)=0,f(1)=1。
//经过数学的推算可以得出 h<=1.44log(n)
//树的高度被限制于 1.44log(n)， 所以查找元素时使用二分查找，最坏查找 1.44log(n) 次，此时最坏时间复杂度为 1.44log(n)，去掉常数项，时间复杂度为：log(n)。
//为了维持AVL树的特征，每次添加和删除元素都需要一次或多次旋转来调整树的平衡。调整的依据来自于二叉树节点的平衡因子：节点的左子树与右子树的高度差称为该节点的平衡因子，约束范围为 [-1，0，1]。

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

// BalanceFactor 计算平衡因子
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

// RightRotation 右旋操作
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

// LeftRotation 左旋操作
func LeftRotation(Root *AVLTreeNode) *AVLTreeNode {
	Pivot := Root.Right
	B := Pivot.Left
	Pivot.Left = Root
	Root.Right = B
	Root.UpdateTreeHeight()
	Pivot.UpdateTreeHeight()
	return Pivot
}

// LeftRightRotation 先左后右旋操作
func LeftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

// RightLeftRotation 先右后左旋操作
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

// Delete 删除指定节点
func (tree *AVLTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	tree.Root = tree.Root.Delete(value)
}

//当删除的值不等于当前节点的值时，在相应的子树中递归删除，递归过程中会自底向上维护AVL树的特征。
//1、小于删除的值 value < node.Value，在左子树中递归删除：node.Left = node.Left.Delete(value)。
//2、大于删除的值 value > node.Value，在右子树中递归删除：node.Right = node.Right.Delete(value)。
//因为删除后可能因为旋转调整，导致树根节点变了，这时会返回新的树根，递归删除后需要将返回的新根节点赋予原来的老根节点。

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
		} else {
			//只有左子树或者只有右子树
			//只有一个子树，该子树也只有一个节点，将该节点替换被删除的节点，然后置空子树
			if node.Left != nil {
				//第三种情况，删除的节点只有左子树，因为树的特征，可以知道左子树其实只有一个节点，它本身，否则高度就等于2了
				node.Value = node.Left.Value
				node.Times = node.Left.Times
				node.Height = 1
				node.Left = nil
			} else if node.Right != nil {
				//第四种情况，
			}
		}
		return node
	}
	//左右子树删除后需要平衡
	var newNode *AVLTreeNode
	// 相当于删除了右子树的节点，左边比右边高了，不平衡
	if node.BalanceFactor() == 2 {
		if node.Left.BalanceFactor() >= 0 {
			newNode = RightRotation(node)
		} else {
			newNode = LeftRotation(node)
		}
	} else if node.BalanceFactor() == -2 {
		if node.Right.BalanceFactor() <= 0 {
			newNode = LeftRotation(node)
		} else {
			newNode = RightRotation(node)
		}
	}
	if newNode == nil {
		node.UpdateTreeHeight()
		return node
	} else {
		newNode.UpdateTreeHeight()
		return newNode
	}
}
