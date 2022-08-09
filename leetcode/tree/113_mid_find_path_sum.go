package tree

//题目：给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径，说明叶子节点是没有子节点的节点
//解题思路：这道题是257和112题的综合加强版
//复杂度：时间复杂度o(n)，空间复杂度o(1

//解法一
func pathSum(root *CommonTreeNode, sum int) [][]int {
	var slice [][]int
	slice = findPath(root, sum, slice, []int(nil))
	return slice
}

func findPath(tree *CommonTreeNode, sum int, slice [][]int, stack []int) [][]int {
	if tree == nil {
		return slice
	}
	sum -= tree.Val
	stack = append(stack, tree.Val)
	if sum == 0 && tree.Left == nil && tree.Right == nil {
		slice = append(slice, append([]int{}, stack...))
		stack = stack[:len(stack)-1]
	}

	slice = findPath(tree.Left, sum, slice, stack)
	slice = findPath(tree.Right, sum, slice, stack)
	return slice
}

//解法二
func pathSum2(root *CommonTreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return [][]int{[]int{root.Val}}
		}
	}

	var res [][]int
	var path []int
	tmpLeft := pathSum(root.Left, sum-root.Val)
	path = append(path, root.Val)
	if len(tmpLeft) > 0 {
		for i := 0; i < len(tmpLeft); i++ {
			tmpLeft[i] = append(path, tmpLeft[i]...)
		}
		res = append(res, tmpLeft...)
	}

	path = []int{}
	tmpRight := pathSum(root.Right, sum-root.Val)
	path = append(path, root.Val)
	if len(tmpRight) > 0 {
		for i := 0; i < len(tmpRight); i++ {
			tmpRight[i] = append(path, tmpRight[i]...)
		}
		res = append(res, tmpRight...)
	}
	return res
}
