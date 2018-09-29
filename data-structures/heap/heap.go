package heap

import "sync"

/*
*二叉堆
*假设"第一个元素"在数组中的索引为 0 的话，则父节点和子节点的位置关系如下：
*(01) 索引为i的左孩子的索引是 (2*i+1);
*(02) 索引为i的左孩子的索引是 (2*i+2);
*(03) 索引为i的父结点的索引是 floor((i-1)/2);
 */
type Item interface {
	Less(than Item) bool
}

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type Heap struct {
	sync.Mutex
	data []Item
	min  bool
}

func New() *Heap {
	return &Heap{
		data: make([]Item, 0),
	}
}

//最小堆
func NewMin() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  true,
	}
}

//最大堆
func NewMax() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  false,
	}
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Len() int {
	return len(h.data)
}
func (h *Heap) Get(n int) Item {
	return h.data[n]
}

//根据生成最大堆还是最小堆，返回比较
func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}

//插入元素
func (h *Heap) Insert(n Item) {
	h.Lock()
	defer h.Unlock()
	h.data = append(h.data, n)
	h.siftUp()
	return
}

//取出元素
func (h *Heap) Extract() (el Item) {
	h.Lock()
	defer h.Unlock()
	if h.Len() == 0 {
		return
	}
	el = h.data[0]
	last := h.data[h.Len()-1]
	if h.Len() == 1 {
		h.data = nil
		return
	}
	h.data = append([]Item{last}, h.data[1:h.Len()-1]...)
	h.siftDown()
	return
}

//上升
func (h *Heap) siftUp() {
	for i, parent := h.Len()-1, h.Len()-1; i > 0; i = parent {
		parent = i >> 1 //位移，向下取整
		if h.Less(h.Get(i), h.Get(parent)) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}

//下降
func (h *Heap) siftDown() {
	for i, child := 0, 1; i < h.Len() && i<<1+1 < h.Len(); i = child {
		child = i<<1 + 1
		if child+1 <= h.Len()-1 && h.Less(h.Get(child+1), h.Get(child)) {
			child++
		}

		if h.Less(h.Get(i), h.Get(child)) {
			break
		}

		h.data[i], h.data[child] = h.data[child], h.data[i]
	}
}
