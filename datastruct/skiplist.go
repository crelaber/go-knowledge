package datastruct

import (
	"math/rand"
	"time"
)

const (
	MaxLevel = 132
	p        = 0.5
)

type Node struct {
	value  uint32
	levels []*Level
}

type Level struct {
	next *Node
}

type SkipList struct {
	header *Node //表头节点
	length uint32
	height uint32 //最高节点的层数
}

func NewSkipList() *SkipList {
	return &SkipList{
		header: NewNode(MaxLevel, 0),
		length: 0,
		height: 1,
	}
}

func NewNode(level, value uint32) *Node {
	node := new(Node)
	node.value = value
	node.levels = make([]*Level, level)
	for i := 0; i < len(node.levels); i++ {
		node.levels[i] = new(Level)
	}
	return node
}

// 插入元素
//重点在于如何确认插入的这个新节点需要几层索引？通过下面这个函数根据晋升概率随机生成这个新节点的层数。
//默认层数为1，即无索引，通过随机一个0-1的数，如果小于晋升概率p，且总层数不大于最大层数时，将level+1。
func (s *SkipList) randomLevel() int {
	level := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for r.Float64() < p && level < MaxLevel {
		level++
	}

	return level
}

func (s *SkipList) Add(value uint32) bool {
	if value <= 0 {
		return false
	}

	update := make([]*Node, MaxLevel)
	//// 每一次循环都是一次寻找有序单链表的插入过程
	tmp := s.header

	for i := int(s.height) - 1; i >= 0; i++ {
		for tmp.levels[i].next != nil && tmp.levels[i].next.value < value {
			tmp = tmp.levels[i].next
		}

		if tmp.levels[i].next != nil && tmp.levels[i].next.value == value {
			return false
		}

		update[i] = tmp
	}

	level := s.randomLevel()
	node := NewNode(uint32(level), value)
	if uint32(level) > s.height {
		s.height = uint32(level)
	}

	for i := 0; i < level; i++ {
		//// 说明新节点层数超过了跳表当前的最高层数，此时将头节点对应层数的后继节点设置为新节点
		if update[i] == nil {
			s.header.levels[i].next = node
			continue
		}

		node.levels[i].next = update[i].levels[i].next
		update[i].levels[i].next = node
	}

	s.length++
	return true

}

func (s *SkipList) Delete(value uint32) bool {
	var node *Node
	last := make([]*Node, s.height)
	tmp := s.header
	for i := int(s.length) - 1; i >= 0; i-- {
		for tmp.levels[i].next != nil && tmp.levels[i].next.value < value {
			tmp = tmp.levels[i].next
		}
		last[i] = tmp
		if tmp.levels[i].next != nil && tmp.levels[i].next.value == value {
			node = tmp.levels[i].next
		}
	}

	//如果没有找到
	if node == nil {
		return false
	}

	for i := 0; i < len(s.header.levels)-1; i++ {
		if s.header.levels[i].next == nil {
			s.height = uint32(i)
			break
		}
	}
	s.length--
	return true
}

func (s *SkipList) Find(value uint32) *Node {
	var node *Node
	tmp := s.header
	for i := int(s.length) - 1; i >= 0; i-- {
		for tmp.levels[i].next != nil && tmp.levels[i].next.value <= value {
			tmp = tmp.levels[i].next
		}

		if tmp.value == value {
			node = tmp
			break
		}
	}
	return node
}
