package datastruct

type Heap struct {
	Size  int
	Array []int
}

func NewHeap(array []int) *Heap {
	h := new(Heap)
	h.Array = array
	return h
}

func (h *Heap) Push(x int) {
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	i := h.Size
	for i > 0 {
		parent := (i - 1) / 2
		if x <= h.Array[parent] {
			break
		}
		h.Array[i] = h.Array[parent]
		i = parent
	}
	h.Array[i] = x
	h.Size++
}
