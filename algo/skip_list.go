package algo

import "fmt"

//跳跃表
//定义：在一组数据中，按照一定规则对数据创建层级索引，通过自上而下的所以查询查找数据的位置
//特点
//1、二分查找针对数据，为了实现不使用针对数组二分查找，产生了跳跃表
//2、跳跃表使用连表
//3、有多级索引
//时间复杂度
//1、修改（增删改）时间复杂度O(lgn)
//2、查询时间复杂度O(lgn)

//空间复杂度
//1、空间复杂度根据索引生成方法有关
//n/2+n/4+...+4+2等比求和Sn=(a1-anq)/1-q O(n-2)约等于O(n)

type SkipListNode struct {
	Data      int
	NextPoint *SkipListNode
	PrePoint  *SkipListNode
	NextLevel *SkipListNode
}

type LinkedList struct {
	Head    *SkipListNode
	Current *SkipListNode
	Tail    *SkipListNode
	Length  int
	Level   int
}

type SkipList struct {
	List        LinkedList
	FirstIndex  LinkedList
	SecondIndex LinkedList
}

func InitSkipList() {
	data := []int{11, 12, 13, 19, 21, 31, 33, 42, 51, 62}
	sl := SkipList{}
	sl.initSkip(data)
	sl.add(11)
	showSkipList(sl)
}

func (sl *SkipList) initSkip(list []int) {
	sl.List = LinkedList{}
	sl.FirstIndex = LinkedList{}
	sl.SecondIndex = LinkedList{}
	var currentNode *SkipListNode
	for i := 0; i < len(list); i++ {
		currentNode = new(SkipListNode)
		currentNode.Data = list[i]
		addNode(sl, currentNode)
	}
}

func showSkipLinkedList(link LinkedList, name int) {
	var currentNode *SkipListNode
	currentNode = link.Head
	for {
		i := 1
		fmt.Println(name, "-Node:", currentNode.Data)
		if currentNode.NextPoint == nil {
			break
		} else {
			currentNode = currentNode.NextPoint
		}

		if name == 1 {
			fmt.Print("---------->")
		} else if name == 2 {
			for i <= 3 {
				fmt.Print("---------->")
				i++
			}
		} else {
			for i <= 7 {
				fmt.Print("---------->")
				i++
			}
		}
	}
	fmt.Println("")
}

func (sl *SkipList) find(x int) {
	var current *SkipListNode
	current = sl.SecondIndex.Head
	if current.Data == x {
		fmt.Println(current.Data)
		return
	}
	if x < current.Data {
		panic("no exist in skip list")
		return
	}
	for {
		if x > current.Data {
			fmt.Println(current.Data)
			current = current.NextPoint
		} else if x < current.Data {
			//吓到底层索引
			fmt.Println(current.Data)
			current = current.PrePoint.NextLevel.NextPoint
		} else {
			fmt.Println(current.Data)
			return
		}
	}
}

func (sl *SkipList) add(x int) {
	var current *SkipListNode
	current = sl.SecondIndex.Head
	if current.Data == x {
		panic("head exists in skipList")
		return
	}

	if x < current.Data {
		newNode2 := new(SkipListNode)
		newNode2.Data = x
		newNode2.NextPoint = sl.SecondIndex.Head
		sl.SecondIndex.Head.PrePoint = newNode2
		sl.SecondIndex.Head = newNode2

		newNode1 := new(SkipListNode)
		newNode1.Data = x
		newNode1.NextPoint = sl.FirstIndex.Head
		sl.FirstIndex.Head.PrePoint = newNode1
		sl.FirstIndex.Head = newNode1

		newNode := new(SkipListNode)
		newNode.Data = x
		newNode.NextPoint = sl.List.Head
		sl.SecondIndex.Head.PrePoint = newNode
		sl.List.Head = newNode
		return
	}

	for {
		if x > current.Data {
			if current.NextPoint == nil {
				if current.NextLevel != nil {
					current = current.NextLevel
				} else {
					//插入
					newNode := new(SkipListNode)
					newNode.Data = x
					current.NextPoint = newNode
					newNode.PrePoint = current
					return
				}
			} else {
				fmt.Println(current.Data)
				current = current.NextPoint
			}
		} else if x < current.Data {
			fmt.Println(current.Data)
			//向下去寻找第一个大于x的值
			if current.PrePoint.NextLevel != nil {
				current = current.PrePoint.NextLevel.NextPoint
			} else {
				newNode := new(SkipListNode)
				newNode.Data = x
				current.PrePoint.NextPoint = newNode
				newNode.NextPoint = current
				current.PrePoint = newNode
				return
			}
		} else {
			fmt.Println(current.Data)
			return
		}
	}
}

//显示跳表
func showSkipList(sl SkipList) {
	showSkipLinkedList(sl.SecondIndex, 3)
	fmt.Println("")
	showSkipLinkedList(sl.FirstIndex, 2)
	fmt.Println("")
	showSkipLinkedList(sl.List, 1)
}

//添加节点
func addNode(skipList *SkipList, node *SkipListNode) {
	insertToLink(&skipList.List, node)
	if skipList.FirstIndex.Length == 0 || ((skipList.List.Length-1)%2 == 0 && skipList.List.Length > 2) {
		newNode := new(SkipListNode)
		newNode.Data = node.Data
		newNode.NextLevel = node
		insertToLink(&skipList.FirstIndex, newNode)
		if skipList.SecondIndex.Length == 0 || ((skipList.FirstIndex.Length-1)%2 == 0 && skipList.FirstIndex.Length > 2) {
			newNode2 := new(SkipListNode)
			newNode2.Data = node.Data
			newNode2.NextLevel = newNode
			insertToLink(&skipList.SecondIndex, newNode2)
		}
	}
}

//插入link list中
func insertToLink(link *LinkedList, node *SkipListNode) {
	if link.Head == nil {
		link.Head = node
		link.Tail = node
		link.Current = node
	} else {
		link.Tail.NextPoint = node
		node.PrePoint = link.Tail
		link.Tail = node
	}
	link.Length++
}
