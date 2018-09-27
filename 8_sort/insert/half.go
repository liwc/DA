package insert

import (
	"fmt"
)

// HalfSort 折半插入[]
func HalfSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] > data[i - 1] {
			continue
		}

		low := 0
		high := i - 1
		for low <= high {
			pos := (low + high) / 2
			if data[i] > data[pos] {
				low = pos + 1
			} else if data[i] < data[pos] {
				high = pos - 1
			}
		}

		temp := data[i]
		for j := i; j > low; j-- {
			data[j] = data[j - 1]
		}
		data[low] = temp
		// fmt.Printf("dst: %v\n", data)
	}
}

// HalfSort2 折半插入[)
func HalfSort2(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] > data[i - 1] {
			continue
		}

		low := 0
		high := i
		for low < high {
			pos := (low + high) / 2
			if data[i] > data[pos] {
				low = pos + 1
			} else if data[i] < data[pos] {
				high = pos
			}
		}

		temp := data[i]
		for j := i; j > low; j-- {
			data[j] = data[j - 1]
		}
		data[low] = temp
		// fmt.Printf("dst: %v\n", data)
	}
}

// HalfTest ...
func HalfTest(array []int) {
	fmt.Println("HalfTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	HalfSort(data)
	fmt.Printf("dst: %v\n", data)

	data2 := make([]int, len(array))
	copy(data2, array)
	HalfSort2(data2)
	fmt.Printf("dst: %v\n", data2)
	fmt.Println("HalfTest end")
}