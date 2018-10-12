package index

import (
	"fmt"
	"DA/4_queue/queue"
)

// IdxSort1 索引排序
func IdxSort1(a []int, n int) {
	if n < 1 {
		return
	}

	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	// fmt.Printf("idx: %v\n", index)

	for i := 1; i < n; i++ {

		j := i
		for j != 0 {
			if a[i] < a[j - 1] {
				index[i]--
				index[j - 1]++
			}
			j--
		}

		// fmt.Printf("idx: %v\n", index)
	}

	fmt.Printf("idx: %v\n", index)

	modified := true
	for modified {
		modified = false

		for i := 0; i < n; i++ {
			if index[i] != i {
				modified = true

				j := index[i]
				a[i], a[j] = a[j], a[i]
				index[i], index[j] = index[j], index[i]

				// fmt.Printf("src: %v\n", a)
				// fmt.Printf("idx: %v\n", index)
			}
		}
	}
}

// IdxSort2 索引排序
func IdxSort2(a []int, n int) {
	if n < 1 {
		return
	}

	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	// fmt.Printf("idx: %v\n", index)

	for i := 1; i < n; i++ {
		j := i
		for j != 0 {
			if a[index[j]] >= a[index[j - 1]] {
				break
			}

			index[j], index[j - 1] = index[j - 1], index[j]
			j--
		}

		// fmt.Printf("idx: %v\n", index)
	}

	fmt.Printf("idx: %v\n", index)

	lq := queue.LqInit()
	for i := 0; i < n; i++ {
		lq.Push(a[index[i]])
	}

	cnt := lq.Length
	for i := 0; i < cnt; i++ {
		a[i] = lq.Pop().(int)
	}
}

// IdxTest ...
func IdxTest(array []int) {
	fmt.Println("IndexTest begin")
	// data := make([]int, len(array))
	// copy(data, array)
	data := []int{3, 8, 6, 2, 7}
	fmt.Printf("src: %v\n", data)
	// IdxSort1(data, len(data))
	IdxSort2(data, len(data))
	fmt.Printf("dst: %v\n", data)
	fmt.Println("IndexTest end")
}