package algo

import (
	"fmt"
)

//拓扑排序
//定义：对一些有前后依赖关系的排序算法，是利用有向无环图进行实现，通过局部依赖关系确定全局顺序的算法
//应用场景：编译有序依赖的文件
//两种拓普排序算法
//Kahn算法
//算法逻辑：利用贪心算法，如果两个顶点，顶点b依赖于顶点a，就将a指向b，当一个顶点的入度为0，将这个顶点就是
//最优排序点，并且将顶点从图中移除，将可达顶点的入度减一

//DFS算法
//1、使用深度算法，产生逆向玲姐表先输出其他依赖，最后输出给自己

type graph struct {
	vertex int           //顶点
	list   map[int][]int //链接表边
}

//创建图
func NewGraph(v int) *graph {
	g := new(graph)
	g.vertex = v
	g.list = map[int][]int{}
	i := 0
	for i < v {
		g.list[i] = make([]int, 0)
		i++
	}
	return g
}

func (g *graph) addVertex(t int, s int) {
	g.list[t] = push(g.list[t], s)
}

//取出切片第一个
func pop(list []int) (int, []int) {
	if len(list) > 0 {
		a := list[0]
		b := list[1:]
		return a, b
	} else {
		return -1, list
	}
}

func push(list []int, value int) []int {
	result := append(list, value)
	return result
}

func (g *graph) KhanSort() {
	var inDegree = make(map[int]int)
	var queue []int
	for i := 0; i < g.vertex; i++ {
		for _, m := range g.list[i] {
			inDegree[m]++
		}
	}

	for i := 1; i <= g.vertex; i++ {
		if inDegree[i] == 0 {
			queue = push(queue, i)
		}
	}

	for len(queue) > 0 {
		var now int
		now, queue = pop(queue)
		fmt.Println("->", now)
		for _, k := range g.list[now] {
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = push(queue, k)
			}
		}
	}
}

func (g *graph) DfsSort() {
	inverseList := make(map[int][]int)
	//初始化逆向邻接表
	for i := 1; i < g.vertex; i++ {
		for _, k := range g.list[i] {
			inverseList[k] = append(inverseList[k], i)
		}
	}

	visited := make([]bool, g.vertex+1)
	visited[0] = true
	for i := 1; i <= g.vertex; i++ {
		if visited[i] == false {
			visited[i] = true
			dfs(i, inverseList, visited)
		}
	}
}

func dfs(vertex int, inverseList map[int][]int, visited []bool) {
	for _, w := range inverseList[vertex] {
		if visited[w] == true {
			continue
		} else {
			visited[w] = true
			dfs(w, inverseList, visited)
		}
	}
	fmt.Println("->", vertex)
}
