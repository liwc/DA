package line

import (
	"fmt"
	// "DA/4_queue/queue"
	"strconv"
)

// RadixSort 基数排序
func RadixSort(a []int, n int) {
	if n <= 1 {
		return
	}

	max := a[0]
	for i := 1; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	r := len(strconv.Itoa(max))
	for i := 1; i <= r; i++ {
		tmp := make([]int, n)
		copy(tmp, a)
		cnt := make([]int, 10)

		d := 1
		for j := 1; j < i; j++ {
			d *= 10
		}

		for k := 0; k < n; k++ {
			cnt[a[k] / d % 10]++
		}

		for k := 1; k < 10; k++ {
			cnt[k] += cnt[k - 1]
		}

		for k := n - 1; k >= 0; k-- {
			cnt[tmp[k] / d % 10]--
			pos := cnt[tmp[k] / d % 10]
			a[pos] = tmp[k]
		}
	}
}

// RadixTest ...
func RadixTest(array []int) {
	fmt.Println("RadixTest begin")
	// data := make([]int, len(array))
	// copy(data, array)
	data := []int{123, 234, 45, 111, 6, 128}
	fmt.Printf("src: %v\n", data)
	RadixSort(data, len(data))
	fmt.Printf("dst: %v\n", data)
	fmt.Println("RadixTest end")
}