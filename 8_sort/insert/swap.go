package insert

import (
	"fmt"
)

// SwapSort 交换插入
func SwapSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j >= 1; j-- {
			if data[j] >= data[j - 1] {
				break
			}
			data[j], data[j - 1] = data[j - 1], data[j]
		}
	}
}

// SwapTest ...
func SwapTest(array []int) {
	fmt.Println("SwapTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	SwapSort(data)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("SwapTest end")
}