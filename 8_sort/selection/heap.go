package selection

import (
	"fmt"
)

// HeapSort 堆排序
func HeapSort(data []int, n int) {
	if (n <= 1) {
		return
	}

	temp := make([]int, n)
	copy(temp, data)

	createHeap(temp, n)
	for i := 0; i < len(temp); i++ {
		data[i] = delete(temp, &n, 0)
	}
}

func createHeap(data []int, n int) {
	for i := (n - 2) / 2; i >= 0; i-- {
		siftdown(data, n, i)
	}
}

func insert(a []int, n int, data int) {
	a[n] = data
	n++
	siftup(a, n, n - 1)
}

func delete(a []int, n *int, pos int) int {
	res := a[pos]

	a[pos] = a[*n - 1]
	*n--
	siftdown(a, *n, pos)

	return res
}

func siftup(data []int, n int, pos int) {
	j := pos
	i := (j - 1) / 2
	for i >= 0 {
		if data[i] <= data[j] {
			break
		}

		data[i], data[j] = data[j], data[i]
		j = i
		i = (j - 1) / 2
	}
}

func siftdown(data []int, n int, pos int) {
	i := pos
	j := i * 2 + 1
	for j < n {
		if j + 1 < n && data[j + 1] < data[j] {
			j = j + 1
		}

		if data[i] <= data[j] {
			break
		}
		data[i], data[j] = data[j], data[i]

		i = j
		j = i * 2 + 1
	}
}

// HeapTest ...
func HeapTest(array []int) {
	fmt.Println("HeapTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	HeapSort(data, 10)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("HeapTest end")
}