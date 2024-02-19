package sort

func Counting(arr []int) {
	Max := arr[0]
	for _, v := range arr {
		if v > Max {
			Max = v
		}
	}
	Max++
	count := make([]int, Max)
	for _, v := range arr {
		count[v]++
	}
	for i := 1; i < len(count); i++ {
		count[i] = count[i] + count[i-1]
	}
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		res[count[v]-1] = v
		count[v]--
	}
	for i, v := range res {
		arr[i] = v
	}
}
