package usage

import (
	"fmt"
	"DA/3_stack/stack"
)

func decimal2binary(value int) *stack.OrderStack {
	os := stack.OsInit()
	for {
		os.Push(value % 2)
		value = value / 2

		if value == 0 {
			break
		}
	}

	return os
}

// ConTest ...
func ConTest() {
	value := 19
	fmt.Printf("%d in binary: ", value)
	os := decimal2binary(value)
	for !os.Empty() {
		fmt.Printf("%v ", os.Pop())
	}

	fmt.Println()
}