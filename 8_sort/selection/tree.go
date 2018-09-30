package selection

import (
	"fmt"
)

// 树形选择排序
type rec struct {
	data int
	index int
	active bool
}

func pow2(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 2	
	}

	return res
}

func fixup(tree []rec, pos int) {
	if pos % 2 == 1 {
		tree[(pos - 1) / 2] = tree[pos + 1]
	} else {
		tree[(pos - 1) / 2] = tree[pos - 1]
	}

	pos = (pos - 1) / 2
	for pos != 0 {
		brother := pos + 1
		if pos % 2 == 0 {
			brother = pos - 1
		}

		if !tree[pos].active || !tree[brother].active {
			if tree[pos].active {
				tree[(pos - 1) / 2] = tree[pos];
			} else {
				tree[(pos - 1) / 2] = tree[brother];
			}
		} else {
			if tree[pos].data <= tree[brother].data {
				tree[(pos - 1) / 2] = tree[pos];
			} else {
				tree[(pos - 1) / 2] = tree[brother];
			}
		}

		pos = (pos - 1) / 2
	}
}

// TreeSort 树形选择排序
func TreeSort(data []int, n int) {
	leaf := 0
	for i := 0; ; i++ {
		leaf = pow2(i)
		if leaf >= n {
			break 
		}
	}

	tree := make([]rec, leaf * 2 - 1)
	for i := 0; i < leaf; i++ {
		if i < n {
			tree[leaf - 1 + i].data = data[i]
			tree[leaf - 1 + i].index = i
			tree[leaf - 1 + i].active = true
		} else {
			tree[leaf - 1 + i].active = false
		}
	}

	index := leaf -1
	for index != 0 {
		i := index
		for i < index * 2 {
			if !tree[i + 1].active || tree[i + 1].data > tree[i].data {
				tree[(i - 1) / 2] = tree[i]
			} else {
				tree[(i - 1) / 2] = tree[i + 1]
			}

			i += 2
		}

		index = (index - 1) / 2
	}

	index = 0
	for index < n - 1 {
		data[index] = tree[0].data
		tree[leaf - 1 + tree[0].index].active = false
		fixup(tree, leaf - 1 + tree[0].index)
		index++
	}
	data[n - 1] = tree[0].data
}

// TreeTest ...
func TreeTest(array []int) {
	fmt.Println("TreeTest begin")
	data := make([]int, len(array))
	copy(data, array)
	fmt.Printf("src: %v\n", data)
	TreeSort(data, len(data))
	fmt.Printf("dst: %v\n", data)
	fmt.Println("TreeTest end")
}