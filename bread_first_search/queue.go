package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return 0
	}

	item := q.items[0]

	q.items = q.items[1:]

	return item
}
