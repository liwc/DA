package main

import (
	"DA/8_sort/insert"
	"DA/8_sort/selection"
	"DA/8_sort/swap"
	"DA/8_sort/merge"
)

func main() {
	array := []int{3, 5, 6, 8, 4, 7, 9, 2, 0, 1}

	// 插入排序
	insert.DirectTest(array)
	insert.SwapTest(array)
	insert.HalfTest(array)
	insert.BinaryTest(array)
	insert.ListTest(array)
	insert.ListHalfTest(array)
	insert.ShellTest(array)

	// 选择排序
	selection.SimpleTest(array)
	selection.TreeTest(array)
	selection.HeapTest(array)

	// 交换排序
	swap.BubbleTest(array)
	swap.QuickTest(array)

	// 归并排序
	merge.MgTest(array)
}
