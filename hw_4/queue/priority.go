package queue

type Item struct {
	priority int
	value    string
	next     *Item
}

func InitItem(p int, v string) *Item {
	return &Item{priority: p, value: v}
}

type Priority struct {
	list *Item
}

func InitPriority() *Priority {
	return &Priority{}
}
func (q *Priority) enqueue(p int, v string) {
	item := InitItem(p, v)
	if q.list == nil {
		q.list = item
		return
	}
	if p < q.list.priority {
		item.next = q.list
		q.list = item
		return
	}
	for i := q.list; ; i = i.next {
		if i.next == nil {
			i.next = item
			return
		}
		if p < i.next.priority {
			item.next = i.next
			i.next = item
			return
		}
	}
}

func (q *Priority) dequeue() *Item {
	item := q.list
	if item == nil {
		return nil
	}
	q.list = item.next
	return item
}
