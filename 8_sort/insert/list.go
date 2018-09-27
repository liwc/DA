package insert

import (
	"fmt"
)

type node struct {
	data int
	next int
	prev int
}

// ListSort 表插入
func ListSort(data []int) {
	length := len(data)

	array := make([]node, length + 1)
	for i := 0; i < length; i++ {
		array[i + 1].data = data[i]
		array[i + 1].next = 0
	}
	array[0].data = 1000
	array[0].next = 1

	for i := 2; i < length + 1; i++ {
		q := 0
		p := array[q].next
		for array[i].data >= array[p].data {
			q = p
			p = array[q].next
		}
		array[i].next = p
		array[q].next = i
	}

	// p := array[0].next
	// for i := 0; i < length && p != 0; i++ {
	// 	data[i] = array[p].data
	// 	p = array[p].next
	// }

	p := array[0].next
	for i := 1; i <= length; i++ {
		for p < i {
			p = array[p].next
		}

		q := array[p].next
		if p != i {
			temp := array[i]
			array[i] = array[p]
			array[p] = temp

			array[i].next = p
		}

		p = q
	}

	for i := 0; i < length; i++ {
		data[i] = array[i + 1].data
	}
}

// ListTest ...
func ListTest(array []int) {
	fmt.Println("ListTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	ListSort(data)
	fmt.Printf("dst: %v\n", data)
	fmt.Println("ListTest end")
}
