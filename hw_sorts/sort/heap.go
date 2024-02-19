package sort

func Heap(l []int) {
	for i := len(l)/2 - 1; i >= 0; i-- {
		heapify(i, len(l), l)
	}
	for i := len(l) - 1; i > 0; i-- {
		l[0], l[i] = l[i], l[0]
		heapify(0, i, l)
	}
}

func heapify(root, size int, l []int) {
	P := root
	L := 2*P + 1
	R := L + 1
	if L < size && l[L] > l[P] {
		P = L
	}
	if R < size && l[R] > l[P] {
		P = R
	}
	if P == root {
		return
	}
	l[root], l[P] = l[P], l[root]
	heapify(P, size, l)
}
