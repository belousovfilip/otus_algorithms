package sort

func Radix(arr []int) {
	Max := arr[0]
	for _, v := range arr {
		if v > Max {
			Max = v
		}
	}
	for j := 1; j <= Max; j *= 10 {
		count := make([]int, 10)
		for _, v := range arr {
			count[v/j%10]++
		}
		for i := 1; i < len(count); i++ {
			count[i] = count[i] + count[i-1]
		}
		res := make([]int, len(arr))
		for i := len(arr) - 1; i >= 0; i-- {
			v := arr[i]
			res[count[v/j%10]-1] = v
			count[v/j%10]--
		}
		for i, v := range res {
			arr[i] = v
		}
	}
}
