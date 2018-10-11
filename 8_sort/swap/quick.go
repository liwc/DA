package swap

import (
	"fmt"
	"DA/3_stack/stack"
)

// QuickSort10 快速排序
func QuickSort10(a []int, n int) {
	if n <= 1 {
		return
	}

	pivot := a[0]
	i := 0
	j := n - 1
	for i < j {
		for i < j && a[j] >= pivot {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
			// fmt.Printf("tmp: %v\n", a)
		}

		for i < j && a[i] <= pivot {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
			// fmt.Printf("tmp: %v\n", a)
		}
	}
	a[i] = pivot
	// fmt.Printf("tmp: %v\n", a)

	QuickSort10(a[:i], i)
	QuickSort10(a[i + 1:], n - i - 1)
}

// QuickSort11 快速排序
func QuickSort11(a []int, n int) {
	if n <= 1 {
		return
	}

	pivot := a[n - 1]
	i := 0
	j := n - 1
	for i < j {
		for i < j && a[i] <= pivot {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}

		for i < j && a[j] >= pivot {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
	}
	a[i] = pivot

	QuickSort10(a[:i], i)
	QuickSort10(a[i + 1:], n - i - 1)
}

// QuickSort20 快速排序
func QuickSort20(a []int, low, high int) {
	if low >= high {
		return
	}

	pos := partition(a, low, high)
	QuickSort20(a, low, pos - 1)
	QuickSort20(a, pos + 1, high)
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i + 1], a[high] = a[high], a[i + 1]
	return i + 1
}

// QuickSort30 快速排序
func QuickSort30(a []int, low, high int) {
	if low >= high {
		return
	}

	os := stack.OsInit()
	pos := partition(a, low, high)
	if low < pos - 1 {
		os.Push(pos - 1)
		os.Push(low)		
	}
	if high > pos + 1 {
		os.Push(high)
		os.Push(pos + 1)
	}

	for !os.Empty() {
		l := os.Pop().(int)
		h := os.Pop().(int)
		pos = partition(a, l, h)
		if l < pos - 1 {
			os.Push(pos - 1)
			os.Push(l)		
		}
		if h > pos + 1 {
			os.Push(h)
			os.Push(pos + 1)
		}
	}
}

// QuickTest ...
func QuickTest(array []int) {
	fmt.Println("QuickTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	// QuickSort10(data, len(data))
	// QuickSort11(data, len(data))
	// QuickSort20(data, 0, len(data) - 1)
	QuickSort30(data, 0, len(data) - 1)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("QuickTest end")
}