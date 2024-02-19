package sort

import (
	"container/list"
)

func Bucket(arr []int) {
	maxValue := arr[0]
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}
	maxValue++
	buckets := make([]*list.List, maxValue*len(arr)/maxValue)
	for _, v := range arr {
		i := v * len(arr) / maxValue
		if buckets[i] == nil {
			buckets[i] = list.New()
		}
		buckets[i].PushFront(v)
		current := buckets[i].Front()
		for l := current; l.Next() != nil; l = l.Next() {
			cV := current.Value.(int)
			nV := current.Next().Value.(int)
			if cV < nV {
				continue
			}
			current.Value = nV
			current.Next().Value = cV
			current = current.Next()
		}
	}
	var j int
	for _, bucket := range buckets {
		if bucket == nil {
			continue
		}
		for l := bucket.Front(); l != nil; l = l.Next() {
			arr[j] = l.Value.(int)
			j++
		}
	}
}
