package insert

import (
	"fmt"
)

// DirectSort 直接插入
func DirectSort(data []int) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i
		for ; j >= 1; j-- {
			if temp >= data[j - 1] {
				break
			}
			data[j] = data[j - 1]
		}
		data[j] = temp
	}
}

// DirectTest ...
func DirectTest(array []int) {
	fmt.Println("DirectTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	DirectSort(data)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("DirectTest end")
}