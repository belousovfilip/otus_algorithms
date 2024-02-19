package sort

func Bubble(l []int) {
	for i := len(l) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}
