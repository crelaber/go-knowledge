package template

import "container/list"

// LFUCache
//LFU 是 Least Frequently Used 的缩写，即最不经常最少使用，也是一种常用的页面置换算法，选择访问计数器最小的页面予以淘汰。如下图，缓存中每个页面带一个访问计数器。
//根据 LFU 的策略，每访问一次都要更新访问计数器。当插入 B 的时候，发现缓存中有 B，所以增加访问计数器的计数，并把 B 移动到访问计数器从大到小排序的地方。
// 再插入 D，同理先更新计数器，再移动到它排序以后的位置。当插入 F 的时候，缓存中不存在 F，所以淘汰计数器最小的页面的页面，所以淘汰 A 页面。
// 此时 F 排在最下面，计数为 1。
type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (cache *LFUCache) Get(key int) int {
	value, ok := cache.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	cache.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := cache.lists[currentNode.frequency]; !ok {
		cache.lists[currentNode.frequency] = list.New()
	}

	newList := cache.lists[currentNode.frequency]
	newNode := newList.PushBack(currentNode)
	cache.nodes[key] = newNode
	if currentNode.frequency-1 == cache.min && cache.lists[currentNode.frequency].Len() == 0 {
		cache.min++
	}
	return currentNode.value
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}
	// 如果存在，更新访问次数
	if currentValue, ok := cache.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		cache.Get(key)
		return
	}

	// 如果不存在且缓存满了，需要删除
	if cache.capacity == len(cache.nodes) {
		currentList := cache.lists[cache.min]
		frontNode := currentList.Front()
		delete(cache.nodes, frontNode.Value.(*node).key)
		currentList.Remove(frontNode)
	}
	// 新建结点，插入到 2 个 map 中
	cache.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := cache.lists[1]; !ok {
		cache.lists[1] = list.New()
	}
	newList := cache.lists[1]
	newNode := newList.PushBack(currentNode)
	cache.nodes[key] = newNode
}
