package usage

import (
	"fmt"
	"DA/3_stack/stack"
	// "DA/6_tree/tree"
)

// HuffNode 节点
type HuffNode struct {
	Weight int
	Left int
	Right int
	Parent int
}

// HuffCode 编码
type HuffCode struct {
	Weight int
	code stack.OrderStack
}

// HtreeInit 树
func HtreeInit(weight []int) []HuffNode {
	length := len(weight)
	if length == 0 {
		return nil
	}

	htree := make([]HuffNode, length * 2 - 1)
	for i := 0; i < length * 2 - 1; i++ {
		htree[i].Left = -1
		htree[i].Right = -1
		htree[i].Parent = -1

		if i < length {
			htree[i].Weight = weight[i]
		} else {
			htree[i].Weight = -1
		}
	}

	for i := length; i < length * 2 - 1; i++ {
		p1 := 0
		p2 := 0
		w1 := 100
		w2 := 100

		for j := 0; j < i; j++ {
			if htree[j].Parent == -1 {
				if htree[j].Weight < w1 {
					p2 = p1
					w2 = w1
					p1 = j
					w1 = htree[j].Weight
				} else if htree[j].Weight < w2 {
					p2 = j
					w2 = htree[j].Weight
				}
			}
		}

		htree[i].Left = p1
		htree[i].Right = p2
		htree[p1].Parent = i
		htree[p2].Parent = i
		htree[i].Weight = htree[p1].Weight + htree[p2].Weight
	}

	return htree
}

// HtreeOutput ...
func HtreeOutput(htree []HuffNode) {
	length := len(htree)
	if length == 0 {
		return
	}

	fmt.Println("Htree")
	for i := 0; i < length; i++ {
		fmt.Printf("%d:\tParent: %v\tLeft: %v\tRight: %v\tWeight: %v\n", 
			i, htree[i].Parent, htree[i].Left, htree[i].Right, htree[i].Weight)
	}
	fmt.Println()
}

// HcodeOutput 编码
func HcodeOutput(htree []HuffNode, length int) {
	if length == 0 {
		return
	}

	hcode := make([]HuffCode, length)
	for i := 0; i < length; i++ {
		child := i
		parent := htree[child].Parent
		for parent != -1 {
			if htree[parent].Left == child {
				// 左孩子的分支是 0
				hcode[i].code.Push(0)
			} else if htree[parent].Right == child {
				// 右孩子的分支是 1
				hcode[i].code.Push(1)
			}

			child = parent
			parent = htree[child].Parent
		}

		hcode[i].Weight = htree[i].Weight
	}

	fmt.Println("hcode")
	for i := 0; i < length; i++ {
		fmt.Printf("%d:\t%v\t", i, hcode[i].Weight)
		for !hcode[i].code.Empty() {
			fmt.Printf("%v", hcode[i].code.Pop())
		}

		fmt.Println()
	}
}

// HuffmanTest ...
func HuffmanTest() {
	fmt.Println("HuffmanTest begin")
	array := []int{6, 2, 3, 9, 12, 24, 10, 8}
	htree := HtreeInit(array)
	HtreeOutput(htree)
	HcodeOutput(htree, len(array))
	fmt.Println("HuffmanTest end")
}