package main

import (
	"fmt"
	"DA/6_tree/tree"
	"DA/6_tree/usage"
)

func main() {
	tree.BTreeTest()
	tree.BSTreeTest()
	usage.HeapTest()
	usage.HuffmanTest()
	fmt.Println()
}