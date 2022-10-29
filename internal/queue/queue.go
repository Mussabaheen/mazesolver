package queue

import "fmt"

func (q *Queue) Enqueue(element Node) {
	q.Elements = append(q.Elements, element)
	q.Size++
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) Dequeue() Node {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return Node{}
	}
	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	q.Size--
	return element
}

func (q *Queue) Len() int {
	return q.Size
}
