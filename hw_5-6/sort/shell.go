package sort

func Shell(l []int) {
	for gap := len(l) / 2; gap > 0; gap /= 2 {
		for j := gap; j < len(l); j++ {
			for i := j; i >= gap && l[i-gap] > l[i]; i -= gap {
				l[i-gap], l[i] = l[i], l[i-gap]
			}
		}
	}
}
