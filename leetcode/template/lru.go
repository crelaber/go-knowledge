package template

import (
	"container/list"
	"errors"
	"sync"
)

type LRU struct {
	Max      int
	DataList *list.List
	DataMap  map[interface{}]*list.Element
	Rwlock   sync.RWMutex
}

type Node struct {
	Key   interface{}
	Value interface{}
}

func NewLRU(len int) *LRU {
	return &LRU{
		Max:      len,
		DataList: list.New(),
		DataMap:  make(map[interface{}]*list.Element),
	}
}

func (l *LRU) Add(key interface{}, value interface{}) error {
	l.Rwlock.Lock()
	defer l.Rwlock.Unlock()
	if l.DataList == nil {
		return errors.New("empty data list")
	}
	//已经存在，修改value后移动到队首位Ubuntu
	if e, ok := l.DataMap[key]; ok {
		e.Value.(*Node).Value = value
		l.DataList.MoveToFront(e)
		return nil
	}
	//不存在，在队首插入节点，在map中存储节点
	ele := l.DataList.PushFront(&Node{
		Key:   key,
		Value: value,
	})
	l.DataMap[key] = ele
	//如果队列超过了最大距离，移除队尾的节点，移除map中的key
	if l.Max != 0 && l.DataList.Len() > l.Max {
		if e := l.DataList.Back(); e != nil {
			l.DataList.Remove(e)
			delete(l.DataMap, e.Value.(*Node).Key)
		}
	}
	return nil
}

func (l *LRU) Remove(key interface{}) bool {
	l.Rwlock.Lock()
	defer l.Rwlock.Unlock()
	if l.DataList == nil {
		return false
	}
	if ele, ok := l.DataMap[key]; ok {
		l.DataList.Remove(ele)
		delete(l.DataMap, ele.Value.(*Node).Key)
		return true
	}
	return false
}

func (l *LRU) Get(key interface{}) (value interface{}, ok bool) {
	l.Rwlock.Lock()
	defer l.Rwlock.Unlock()
	if l.DataList == nil {
		return nil, false
	}
	//如果存在，则先移至队首，再返回value
	if ele, ok := l.DataMap[key]; ok {
		l.DataList.MoveToFront(ele)
		return ele.Value.(*Node).Value, true
	}
	return nil, false
}
