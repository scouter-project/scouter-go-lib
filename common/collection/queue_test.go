package collection

import (
	"fmt"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(4)

	for !queue.IsEmpty() {
		fmt.Printf("%d\n", queue.Dequeue())
	}
	queue.Enqueue(1)
	queue.Enqueue(2)

	fmt.Println(queue.Peek())
	fmt.Println(queue.Peek())
}
