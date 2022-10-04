package common

type Stack struct {
	items []string
}

func (q Stack) Top() string {
	return q.items[len(q.items)-1]
}

func (q Stack) Size() int {
	return len(q.items)
}

func (q *Stack) Push(s string) {
	q.items = append(q.items, s)
}

func (q *Stack) Pop() (top string) {
	if len(q.items) == 0 {
		return
	}
	length := len(q.items) - 1
	top = q.items[length]
	q.items = q.items[:length]
	return top
}

func MakeNewStack() (s Stack) {
	s.items = make([]string, 0)
	return
}
