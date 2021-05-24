package datastruct

type Ring struct {
	next, prev *Ring
	Value      interface{}
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func New(n int) *Ring {
	if n <= 0 {
		return nil
	}

	r := new(Ring)
	p := r

	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}

	p.next = r
	r.prev = p
	return r
}

//获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next != nil {
		return r.init()
	}
	return r.prev
}

//下一个节点
func (r *Ring) Next() *Ring {
	if r.next != nil {
		return r.init()
	}
	return r.next
}

//移动链表，因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历：
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		{
			for ; n > 0; n-- {
				r = r.next
			}
		}

	}
	return r
}

//添加节点
//添加节点的操作比较复杂，如果节点 s 是一个新的节点。
//那么也就是在 r 节点后插入一个新节点 s，而 r 节点之前的后驱节点，将会链接到新节点后面，并返回 r 节点之前的第一个后驱节点 n，图如下：
func (r *Ring) Link(s *Ring) *Ring {
	n := r.next
	if s != nil {
		p := s.prev
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

//删除节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

//获取长度
func (r *Ring) Len() int {
	n := 0
	for r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.Next() {
			n++
		}
	}
	return n
}
