package swap

import (
	"fmt"
)

// BubbleSort10 冒泡排序
func BubbleSort10(a []int, n int) {
	for i := 1; i < n; i++ {
		for j := 0; j < n - i; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
}

// BubbleSort11 冒泡排序
func BubbleSort11(a []int, n int) {
	for i := 1; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if a[j - 1] > a[j] {
				a[j - 1], a[j] = a[j], a[j - 1]
			}
		}
	}
}

// BubbleSort20 冒泡排序
func BubbleSort20(a []int, n int) {
	i := n - 1
	flag := true
	for flag {
		flag = false
		
		for j := 0; j < i; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
				flag = true
			}
		}

		i--
	}
}

// BubbleSort21 冒泡排序
func BubbleSort21(a []int, n int) {
	i := 1
	flag := true
	for flag {
		flag = false

		for j := n - 1; j >= i; j-- {
			if a[j - 1] > a[j] {
				a[j - 1], a[j] = a[j], a[j - 1]
				flag = true
			}
		}

		i++
	}
}

// BubbleSort30 冒泡排序
func BubbleSort30(a []int, n int) {
	i := n - 1
	flag := true
	for flag {
		flag = false
		k := 0

		for j := 0; j < i; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
				k = j
				flag = true
			}
		}

		if k == 0 {
			break
		}
		i = k
	}
}

// BubbleSort31 冒泡排序
func BubbleSort31(a []int, n int) {
	i := 1
	flag := true
	for flag {
		flag = false
		k := n - 1

		for j := n - 1; j >= i; j-- {
			if a[j - 1] > a[j] {
				a[j - 1], a[j] = a[j], a[j - 1]
				k = j
				flag = true
			}
		}

		if k == n - 1 {
			break
		}
		i = k
	}
}

// BubbleTest ...
func BubbleTest(array []int) {
	fmt.Println("BubbleTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	// BubbleSort10(data, len(data))
	// BubbleSort11(data, len(data))
	// BubbleSort20(data, len(data))
	// BubbleSort21(data, len(data))
	// BubbleSort30(data, len(data))
	BubbleSort31(data, len(data))
	fmt.Printf("dst: %v\n", data)
	fmt.Println("BubbleTest end")
}