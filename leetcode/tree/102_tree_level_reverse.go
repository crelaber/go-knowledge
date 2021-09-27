package tree

//题目：按照层序遍历一棵树
//解题思路：使用队列

//BFS
func BFS(root *CommonTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*CommonTreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, tmp)
	}
	return res
}

func DFS(root *CommonTreeNode) [][]int {
	var res [][]int
	var dfsLevel func(node *CommonTreeNode, level int)
	dfsLevel = func(node *CommonTreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}
