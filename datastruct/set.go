package datastruct

import (
	"sync"
)

type Set struct {
	m   map[int]struct{} //map 的值我们不使用，所以值定义为空结构体 struct{}，因为空结构体不占用内存空间
	len int
	sync.RWMutex
}

func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

//添加数据
//首先，加并发锁，实现线程安全，然后往结构体 s *Set 里面的内置 map 添加该元素：item，元素作为字典的键，会自动去重。同时，集合大小重新生成。
//时间复杂度等于字典设置键值对的复杂度，哈希不冲突的时间复杂度为：O(1)，否则为 O(n)，可看哈希表实现一章。
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
	s.len = len(s.m)
}

//先加并发锁，然后删除 map 里面的键：item。
//时间复杂度等于字典删除键值对的复杂度，哈希不冲突的时间复杂度为：O(1)，否则为 O(n)，可看哈希表实现一章。
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	if s.len == 0 {
		return
	}
	delete(s.m, item)
	s.len = len(s.m)
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return s.len
}

//时间复杂度：O(1)。
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

//将原先的 map 释放掉，并且重新赋一个空的 map。
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

//将集合转化为列表
//时间复杂度：O(n)。
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
