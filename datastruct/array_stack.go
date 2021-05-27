package datastruct

import "sync"

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (st *ArrayStack) Push(v string) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.array = append(st.array, v)
	st.size = st.size + 1
}

//出栈
func (st *ArrayStack) Pop() string {
	st.lock.Lock()
	defer st.lock.Unlock()
	if st.size == 0 {
		panic("empty")
	}

	v := st.array[st.size-1]
	newArray := make([]string, st.size-1, st.size-1)

	for i := 0; i < st.size-1; i++ {
		newArray[i] = st.array[i]
	}

	st.array = newArray

	st.size = st.size - 1
	return v
}

//获取顶部元素
func (st *ArrayStack) Peek() string {
	if st.size == 0 {
		panic("empty")
	}

	v := st.array[st.size-1]
	return v
}

func (st *ArrayStack) Size() int {
	return st.size
}

func (st *ArrayStack) IsEmpty() bool {
	return st.size == 0
}
