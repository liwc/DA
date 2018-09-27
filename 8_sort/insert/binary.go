package insert

import (
	"fmt"
)

// BinarySort 二路插入
func BinarySort(data []int) {
	array := make([]int, len(data))
	first := 0
	last := 0
	array[0] = data[0]

	length := len(data)
	for i := 1; i < length; i++ {
		if data[i] <= array[first] {
			first = (first - 1 + length) % length
			array[first] = data[i]
		} else {
			last++
			array[last] = data[i]

			if array[last] < array[last - 1] {
				for j := last; j != first; {
					prev := (j - 1 + length) % length
					if array[j] >= array[prev] {
						break
					}
					array[j], array[prev] = array[prev], array[j]
					j = prev
				}
			}
		}

		// fmt.Printf("tmp: %v\n", array)
	}

	for i := 0; i < length; i++ {
		data[i] = array[(first + i) % length]
	}
}

// BinarySort2 二路插入
func BinarySort2(data []int) {
	array := make([]int, len(data))
	first := 0
	last := 0
	array[0] = data[0]

	length := len(data)
	for i := 1; i < length; i++ {
		if data[i] <= array[first] {
			first = (first - 1 + length) % length
			array[first] = data[i]
		} else {
			last++
			array[last] = data[i]

			if array[last] < array[last - 1] {
				low := first
				high := last

				for low != high {
					// mid := (low + high) / 2
					// if low > high {
					// 	mid = (low + high + length) / 2 % length
					// }
					num := (high - low + length) % length
					mid := (low + num / 2) % length
					if array[last] > array[mid] {
						low = (mid + 1) % length
					} else {
						high = mid
					}
				}

				temp := array[last]
				for j := last; j != low; {
					prev := (j - 1 + length) % length
					array[j] = array[prev]
					j = prev
				}
				array[low] = temp
			}
		}

		// fmt.Printf("tmp: %v\n", array)
	}

	for i := 0; i < length; i++ {
		data[i] = array[(first + i) % length]
	}
}

// BinaryTest ...
func BinaryTest(array []int) {
	fmt.Println("BinaryTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	BinarySort(data)
	fmt.Printf("dst: %v\n", data)

	data2 := make([]int, len(array))
	copy(data2, array)
	BinarySort2(data2)
	fmt.Printf("dst: %v\n", data2)
	fmt.Println("BinaryTest end")
}