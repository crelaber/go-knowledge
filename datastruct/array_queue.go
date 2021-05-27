package datastruct

import "sync"

//参考链接：http://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6asrar055v
type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

//入队
func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.array = append(queue.array, v)
	queue.size = queue.size + 1
}

//出队
//时间复杂度是：O(n)。
func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 {
		panic("empty")
	}

	v := queue.array[0]
	newArr := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArr[i-1] = queue.array[i]
	}
	queue.array = newArr
	queue.size = queue.size - 1
	return v
}
