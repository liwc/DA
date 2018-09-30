package selection

import (
	"fmt"
)

// SimpleSort 简单选择排序
func SimpleSort(data []int) {
	n := len(data)
	for i := 0; i < n - 1; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[index] {
				index = j
			}
		}

		if index != i {
			data[i], data[index] = data[index], data[i]
		}
	}
}

// SimpleTest ...
func SimpleTest(array []int) {
	fmt.Println("SimpleTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	SimpleSort(data)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("SimpleTest end")
}