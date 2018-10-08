package priority_queue

import (
	"github.com/xiaomeng79/go-algorithm/data-structures/heap"
	"github.com/xiaomeng79/go-algorithm/data-structures/queue"
)

//优先队列与队列区别:按照优先级自动排序
//抽象的数据类型，具有队列的所有特性，包括基本操作，只是在这基础上添加了内部的一个排序，它本质是一个堆实现的

type Item struct {
	Value    interface{}
	Priority int
}

func NewItem(value interface{}, priority int) (i *Item) {
	return &Item{
		Value:    value,
		Priority: priority,
	}
}

func (x Item) Less(than heap.Item) bool {
	return x.Priority < than.(Item).Priority
}

type PQ struct {
	data heap.Heap
}

func NewMax() (q *PQ) {
	return &PQ{
		data: *heap.NewMax(),
	}
}

func NewMin() (q *PQ) {
	return &PQ{
		data: *heap.NewMin(),
	}
}

func (pq *PQ) Len() int {
	return pq.data.Len()
}

func (pq *PQ) Insert(el Item) {
	pq.data.Insert(heap.Item(el))
}

func (pq *PQ) Extract() (el Item) {
	return pq.data.Extract().(Item)
}

func (pq *PQ) ChangePriority(val interface{}, priority int) {
	var storage = queue.New()

	popped := pq.Extract()

	for val != popped.Value {
		if pq.Len() == 0 {
			panic("Item not found")
		}

		storage.Push(popped)
		popped = pq.Extract()
	}

	popped.Priority = priority
	pq.data.Insert(popped)

	for storage.Len() > 0 {
		pq.data.Insert(storage.Shift().(heap.Item))
	}
}
