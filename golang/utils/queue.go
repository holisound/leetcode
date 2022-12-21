package utils

type Queue struct {
	data []int
}

func (q *Queue) Push(x int) {
	q.data = append(q.data, x)
}

func (q *Queue) Poll() int {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Queue) Front() int {
	return q.data[0]
}

func (q *Queue) Size() int {
	return len(q.data)
}
