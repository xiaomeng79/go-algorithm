package queue

import "container/list"

//使用go标准库的list实现
type ListQueue struct {
	l *list.List
}

func NewListQueue() *ListQueue {
	l := list.New()
	return &ListQueue{l}
}

func (lq *ListQueue) Shift() (el interface{}) {
	e := lq.l.Front()
	el = e.Value
	lq.l.Remove(e)
	return
}

func (lq *ListQueue) Push(el interface{}) {
	lq.l.PushBack(el)
	return
}

func (lq *ListQueue) Len() int {
	return lq.l.Len()
}
