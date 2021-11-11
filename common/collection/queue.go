package collection

import "sync"

type Queue struct {
	Front int32
	Rear  int32
	Data  []interface{}
	sync.RWMutex
}

func (q *Queue) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()
	if q.Front == q.Rear {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	q.Lock()
	defer q.Unlock()
	if q.nextPosition(q.Rear) == q.Front {
		return true
	}
	return false
}

func NewQueue(size int32) *Queue {
	queue := new(Queue)
	queue.Rear = 0
	queue.Front = 0
	queue.Data = make([]interface{}, size+1)
	return queue
}

func (q *Queue) nextPosition(pos int32) int32 {
	if pos == int32(len(q.Data))-1 {
		return 0
	} else {
		return pos + 1
	}
}
func (q *Queue) Enqueue(value interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.Lock()
	defer q.Unlock()
	q.Rear = q.nextPosition(q.Rear)
	q.Data[q.Rear] = value
	return true
}
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	q.Lock()
	defer q.Unlock()
	q.Front = q.nextPosition(q.Front)
	return q.Data[q.Front]
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	q.Lock()
	defer q.Unlock()
	return q.Data[q.nextPosition(q.Front)]
}
