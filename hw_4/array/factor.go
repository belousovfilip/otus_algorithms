package array

type Factor struct {
	items []int
	cap   int
	size  int
}

func (a *Factor) Size() int {
	return a.size
}

func (a *Factor) Get(i int) int {
	return a.items[i]
}

func (a *Factor) Add(v int) {
	a.resizeIfFull()
	a.items[a.Size()] = v
	a.size++
}

func (a *Factor) AddByIndex(i, v int) {
	a.resizeIfFull()
	a.size++
	for j := a.size - 2; j >= i; j-- {
		a.items[j+1] = a.items[j]
	}
	a.items[i] = v
}

func (a *Factor) Remove(i int) int {
	v := a.Get(i)
	for j := i + 1; j < a.Size(); j++ {
		a.items[j-1] = a.items[j]
	}
	a.items[a.Size()-1] = 0
	a.size--
	return v
}
func (a *Factor) resize() {
	c := a.cap * 2
	a.cap = c
	resized := make([]int, c)
	for i := 0; i < a.Size(); i++ {
		resized[i] = a.items[i]
	}
	a.items = resized
}

func (a *Factor) resizeIfFull() {
	if a.size == a.cap {
		a.resize()
	}
}

func InitFactor() Array {
	return &Factor{
		items: make([]int, 1),
		cap:   1,
		size:  0,
	}
}
