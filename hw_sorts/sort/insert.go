package sort

func Insert(l []int) {
	for i := 1; i < len(l); i++ {
		for j := i; j > 0; j-- {
			if l[j-1] > l[j] {
				l[j-1], l[j] = l[j], l[j-1]
				continue
			}
			break
		}
	}
}

func InsertShift(slice []int) {
	for i := 1; i < len(slice); i++ {
		if p := iterateSearch(slice[:i], slice[i]); p != i {
			shift(slice, p, i)
		}
	}
}

func InsertBinary(slice []int) {
	for i := 1; i < len(slice); i++ {
		if p := binarySearch(slice[:i], slice[i]); p != i {
			shift(slice, p, i)
		}
	}
}

func iterateSearch(slice []int, v int) (pos int) {
	for i := len(slice); i > 0; i-- {
		if slice[i-1] < v {
			return i
		}
	}
	return 0
}

func binarySearch(slice []int, v int) int {
	if slice[0] >= v {
		return 0
	}
	last := len(slice) - 1
	if slice[last] <= v {
		return last + 1
	}
	var low int
	var mid int
	var high = last
	for low <= high {
		mid = (low + high) / 2
		if slice[mid] <= v && slice[mid+1] >= v {
			return mid + 1
		}
		if slice[mid] <= v {
			low = mid + 1
			continue
		}
		if slice[mid] >= v {
			high = mid - 1
		}
	}
	return len(slice)
}

func shift(slice []int, a, b int) {
	v := slice[b]
	for i := b - 1; i >= a; i-- {
		slice[i+1] = slice[i]
	}
	slice[a] = v
}
