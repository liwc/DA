package insert

import (
	"fmt"
)

// ListHalfSort 表折半插入
func ListHalfSort(data []int) {
	length := len(data)

	array := make([]node, length + 1)
	array[0].data = 1000
	array[0].next = 1
	array[0].prev = 1
	for i := 0; i < length; i++ {
		array[i + 1].data = data[i]
		array[i + 1].next = 0
		array[i + 1].prev = 0
	}

	for i := 2; i <= length; i++ {
		low := array[0].next
		high := array[0].prev
		sortnum := i - 1

		for low != 0 && high != 0 && array[low].data <= array[high].data {
			mid := low
			sortnum = sortnum / 2
			for index := 1; index < sortnum; index++ {
				mid = array[mid].next
			}

			if array[i].data < array[mid].data {
				high = array[mid].prev
			} else {
				low = array[mid].next
			}
		}

		array[array[low].prev].next = i
		array[i].prev = array[low].prev
		array[i].next = low
		array[low].prev = i
	}

	p := array[0].next
	for i := 0; i < length && p != 0; i++ {
		data[i] = array[p].data
		p = array[p].next
	}
}

// ListHalfTest ...
func ListHalfTest(array []int) {
	fmt.Println("ListHalfTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	ListHalfSort(data)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("ListHalfTest end")
}