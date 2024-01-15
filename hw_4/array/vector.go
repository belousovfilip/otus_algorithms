package array

type Vector struct {
	items  []int
	cap    int
	size   int
	vector int
}

func (a *Vector) Size() int {
	return a.size
}

func (a *Vector) Get(i int) int {
	return a.items[i]
}

func (a *Vector) Add(v int) {
	a.resizeIfFull()
	a.items[a.Size()] = v
	a.size++
}

func (a *Vector) AddByIndex(i, v int) {
	a.resizeIfFull()
	a.size++
	for j := a.size - 2; j >= i; j-- {
		a.items[j+1] = a.items[j]
	}
	a.items[i] = v
}

func (a *Vector) Remove(i int) int {
	v := a.Get(i)
	for j := i + 1; j < a.Size(); j++ {
		a.items[j-1] = a.items[j]
	}
	a.items[a.Size()-1] = 0
	a.size--
	return v
}
func (a *Vector) resize() {
	c := a.cap * a.vector
	a.cap = c
	resized := make([]int, c)
	for i := 0; i < a.Size(); i++ {
		resized[i] = a.items[i]
	}
	a.items = resized
}

func (a *Vector) resizeIfFull() {
	if a.size == a.cap {
		a.resize()
	}
}

func InitVector() Array {
	vector := 10
	return &Vector{
		items:  make([]int, 1),
		cap:    1,
		size:   0,
		vector: vector,
	}
}
