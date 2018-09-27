package insert

import (
	"fmt"
)

// ShellSort 希尔排序
func ShellSort(data []int) {
	n := len(data)
	for step := n / 2; step >= 1; step /= 2 {
		for i := step; i < n; i += step {
			for j := i; j > 0; j -= step {
				if data[j] < data[j - step] {
					data[j], data[j - step] = data[j - step], data[j]
				}
			}
		}
	}
}

// ShellTest ...
func ShellTest(array []int) {
	fmt.Println("ShellTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	ShellSort(data)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("ShellTest end")
}