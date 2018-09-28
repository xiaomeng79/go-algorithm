package queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func New() *Queue {
	q := &Queue{}
	q.queue = make([]interface{}, 0)
	q.len = 0
	q.lock = new(sync.Mutex)
	return q
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *Queue) Shift() (el interface{}) {
	el, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) Push(el interface{}) {
	q.queue = append(q.queue, el)
	q.len++
	return
}

func (q *Queue) Peek() (el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue[0]
}
