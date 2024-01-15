package array

type Array interface {
	Size() int
	Get(i int) int
	Add(v int)
	AddByIndex(i, v int)
	Remove(i int) int
}
