package sort

func Merge(arr []int) {
	mSort(arr, 0, len(arr)-1)
}

func mSort(arr []int, L, R int) {
	if L >= R {
		return
	}
	M := (L + R) / 2
	mSort(arr, L, M)
	mSort(arr, M+1, R)
	merge(arr, L, M, R)
}

func merge(arr []int, L, M, R int) {
	T := make([]int, R-L+1)
	a, b, m := L, M+1, 0
	for a <= M && b <= R {
		if arr[a] > arr[b] {
			T[m] = arr[b]
			b++
		} else {
			T[m] = arr[a]
			a++
		}
		m++
	}
	for ; a <= M; a++ {
		T[m] = arr[a]
		m++
	}
	for ; b <= R; b++ {
		T[m] = arr[b]
		m++
	}
	for i, v := range T {
		arr[i+L] = v
	}
}
