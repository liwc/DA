package merge

import (
	"fmt"
	"DA/4_queue/queue"
	// "DA/8_sort/swap"
)

func mergearray(a []int, low, mid, high int) {
	lq := queue.LqInit()

	i := low
	j := mid + 1
	for i <= mid && j <= high {
		for i <= mid && j <= high && a[i] <= a[j] {
			lq.Push(a[i])
			i++
		}
		for i <= mid && j <= high && a[j] < a[i] {
			lq.Push(a[j])
			j++
		}
	}

	for i <= mid {
		lq.Push(a[i])
		i++
	}
	for j <= high {
		lq.Push(a[j])
		j++
	}

	n := lq.Length
	for k := 0; k < n; k++ {
		a[low + k] = lq.Pop().(int)
	}
}

func mergesort(a []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2
	mergesort(a, low, mid)
	mergesort(a, mid + 1, high)
	mergearray(a, low, mid, high)
}

// MgSort 归并排序
func MgSort(a []int, n int) {
	if n < 1 {
		return
	}

	mergesort(a, 0, n - 1)
}

// MgTest ...
func MgTest(array []int) {
	fmt.Println("MergeTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	MgSort(data, len(data))
	fmt.Printf("dst: %v\n", data)
	fmt.Println("MergeTest end")
}