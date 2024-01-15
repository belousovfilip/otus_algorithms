package array

type Single struct {
	items []int
	size  int
}

func (a *Single) Size() int {
	return a.size
}

func (a *Single) Get(i int) int {
	return a.items[i]
}

func (a *Single) Add(v int) {
	a.resize()
	a.items[a.Size()] = v
	a.size++
}

func (a *Single) AddByIndex(i, v int) {
	a.resize()
	a.size++
	for j := a.size - 2; j >= i; j-- {
		a.items[j+1] = a.items[j]
	}
	a.items[i] = v
}

func (a *Single) Remove(i int) int {
	v := a.Get(i)
	for j := i + 1; j < a.Size(); j++ {
		a.items[j-1] = a.items[j]
	}
	a.items[a.Size()-1] = 0
	a.size--
	return v
}
func (a *Single) resize() {
	resized := make([]int, a.size+1)
	for i := 0; i < a.Size(); i++ {
		resized[i] = a.items[i]
	}
	a.items = resized
}

func InitSingle() Array {
	return &Single{
		items: make([]int, 0),
	}
}
