package sort

func Quick(arr []int) {
	qSort(arr, 0, len(arr)-1)
}

func qSort(arr []int, L, R int) {
	if L >= R {
		return
	}
	M := split(arr, L, R)
	qSort(arr, L, M-1)
	qSort(arr, M+1, R)
}

func split(arr []int, L, R int) int {
	P := arr[R]
	M := L - 1
	for j := L; j <= R; j++ {
		if arr[j] <= P {
			M++
			arr[M], arr[j] = arr[j], arr[M]
		}
	}
	return M
}
