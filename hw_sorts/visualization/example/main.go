package main

import (
	tm "github.com/buger/goterm"
	"otus_algo/hw_sorts/visualization"
)

func main() {
	l := []int{54, 6, 3, 2, 6, 1, 2, 12, 5, 30, 16}
	bubble(l)
	//shell(l)
}

func shell(l []int) {
	v := visualization.InitSorting(l, 500)
	v.Print()
	for gap := len(l) / 2; gap > 0; gap /= 2 {
		for j := gap; j < len(l); j++ {
			for i := j; i >= gap; i -= gap {
				v.Compare(i-gap, i)
				if l[i-gap] > l[i] {
					l[i-gap], l[i] = l[i], l[i-gap]
					v.Swap(i-gap, i)
					v.Print()
					continue
				}
				v.Print()
				break
			}
		}
	}
	tm.MoveCursor(0, len(l)+3)
	tm.Flush()
}

func bubble(l []int) {
	v := visualization.InitSorting(l, 700)
	v.Print()
	for i := len(l) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			v.Compare(j, j+1)
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
				v.Swap(j, j+1)
			}
			v.Print()
		}
	}
	tm.MoveCursor(0, len(l)+3)
	tm.Flush()
}
