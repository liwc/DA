package main

import (
	"fmt"
)

func next(n, k, pos int, flag []int) int {
	i := 0
	for {
		pos++
		if pos > n {
			pos = 1
		}

		if flag[pos - 1] == 0 {
			i++
		}

		if i == k {
			break
		}
	}

	return pos
}

func josephus(n, k int) int {
	pos := k % n
	if pos == 0 {
		pos = n
	}

	flag := make([]int, n)
	count := n
	for count > 1 {
		fmt.Printf("%d ", pos)

		flag[pos - 1] = 1
		count--

		pos = next(n, k, pos, flag)
	}	

	return pos
}

func main() {
	fmt.Printf("\nlast: %d\n", josephus(41, 3))
	fmt.Println()
}
