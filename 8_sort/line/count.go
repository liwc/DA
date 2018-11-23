package line

import (
	"fmt"
	// "DA/4_queue/queue"
)

// CountSort 计数排序
func CountSort(a []int, n int) {
	if n < 1 {
		return
	}

	max := a[0]
	for i := 1; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	max = max + 1

	temp := make([]int, n)
	copy(temp, a)
	count := make([]int, max)
	for i := 0; i < n; i++ {
		count[a[i]]++
	}
	fmt.Printf("cnt: %v\n", count)

	for i := 1; i < max; i++ {
		count[i] += count[i - 1]
	}
	fmt.Printf("cnt: %v\n", count)

	for i := n - 1; i >= 0; i-- {
		count[temp[i]]--
		pos := count[temp[i]]
		a[pos] = temp[i]
	}
}

// CountTest ...
func CountTest(array []int) {
	fmt.Println("CountTest begin")
	// data := make([]int, len(array))
	// copy(data, array)
	data := []int{2, 3, 1, 5, 1, 6, 4, 9}
	fmt.Printf("src: %v\n", data)
	CountSort(data, len(data))
	fmt.Printf("dst: %v\n", data)
	fmt.Println("CountTest end")
}