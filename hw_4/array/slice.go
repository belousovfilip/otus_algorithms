package array

type Slice struct {
	items []int
}

func (a *Slice) Size() int {
	return len(a.items)
}

func (a *Slice) Get(i int) int {
	return a.items[i]
}

func (a *Slice) Add(v int) {
	a.items = append(a.items, v)
}

func (a *Slice) AddByIndex(i, v int) {
	if a.Size() == i {
		a.items = append(a.items, v)
		return
	}
	items := append(a.items[:i+1], a.items[i:]...)
	items[i] = v
	a.items = items
}

func (a *Slice) Remove(i int) int {
	v := a.Get(i)
	a.items = append(a.items[:i], a.items[i+1:]...)
	return v
}

func InitSlice() Array {
	return &Slice{
		items: make([]int, 0),
	}
}
