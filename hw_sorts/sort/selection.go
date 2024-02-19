package sort

func Selection(l []int) {
	for i := len(l) - 1; i > 0; i-- {
		var m int
		for j := 1; j <= i; j++ {
			if l[j] > l[m] {
				m = j
			}
		}
		l[i], l[m] = l[m], l[i]
	}
}
