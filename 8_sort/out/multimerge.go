package out

import (
	"fmt"
	"DA/4_queue/queue"
)

// MultiMergeSort 多路归并
func MultiMergeSort(a [][]int) {
	lq := queue.LqInit()
	pos := make([]int, len(a))
	for {
		i := 0
		for ; i < len(a); i++ {
			if pos[i] < len(a[i]) {
				break
			}
		}

		if i >= len(a) {
			break
		}

		min := i
		data := a[min][pos[min]]
		for i = i + 1; i < len(a); i++ {
			if pos[i] < len(a[i]) && a[i][pos[i]] < data {
				min = i
				data = a[min][pos[min]]
			}
		}

		pos[min]++
		lq.Push(data)
	}

	fmt.Print("dst: ")
	lq.Output()
}

// MultiMergeTest ...
func MultiMergeTest(array []int) {
	fmt.Println("MultiMergeTest begin")
	data := [][]int {
		{3, 5, 7},
		{1, 6, 9},
		{2, 4, 8},
		{0, 12, 14},
		{10, 11, 13, 15},
	}

	fmt.Printf("src: %v\n", data[0])
	for i := 1; i < len(data); i++ {
		fmt.Printf("     %v\n", data[i])
	}
	
	MultiMergeSort(data)
	// fmt.Printf("dst: %v\n", data)
	fmt.Println("MultiMergeTest end")
}