package main

import (
	"DA/8_sort/insert"
	"DA/8_sort/selection"
)

func main() {
	array := []int{3, 5, 6, 8, 4, 7, 9, 2, 0, 1}
	insert.DirectTest(array)
	insert.SwapTest(array)
	insert.HalfTest(array)
	insert.BinaryTest(array)
	insert.ListTest(array)
	insert.ListHalfTest(array)
	insert.ShellTest(array)

	selection.SimpleTest(array)
	selection.TreeTest(array)
	selection.HeapTest(array)
}
