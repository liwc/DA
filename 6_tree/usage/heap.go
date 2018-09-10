package usage

import (
	"fmt"
	"DA/6_tree/tree"
)

// DEFAULTSIZE 默认堆大小
const DEFAULTSIZE = 5

// INCREMENT 堆空间每次的增量
const INCREMENT = 5

// HeapType 堆类型
type HeapType int

const (
	// Min 小顶堆
	Min HeapType = iota
	// Max 大顶堆
	Max
)

// Heap 堆定义
type Heap struct {
	Array []tree.Item
	Size int
	Capcity int
	Type HeapType
}

// Empty 堆空的判断
func (heap *Heap) Empty() bool {
	return heap.Size == 0
}

// SiftUp 向上调整
func (heap *Heap) SiftUp(pos int) {
	j := pos
	i := (j - 1) / 2
	for i >= 0 && i != j {
		iData, _ := heap.Array[i].(int)
		jData, _ := heap.Array[j].(int)
		if (heap.Type == Min && iData < jData) ||
		   (heap.Type == Max && iData > jData) {
			break
		}

		heap.Array[i], heap.Array[j] = heap.Array[j], heap.Array[i]
		j = i
		i = (j - 1) / 2
	}
}

// SiftDown 向下调整
func (heap *Heap) SiftDown(pos int) {
	i := pos
	j := i * 2 + 1
	for j < heap.Size {
		iData, _ := heap.Array[i].(int)
		jData, _ := heap.Array[j].(int)
		if j + 1 < heap.Size {
			j1Data, _ := heap.Array[j + 1].(int)
			if (heap.Type == Min && jData > j1Data) ||
			   (heap.Type == Max && jData < j1Data) {
				jData = j1Data
				j = j + 1
			}
		}

		if (heap.Type == Min && iData < jData) ||
		   (heap.Type == Max && iData > jData) {
			break
		}

		heap.Array[i], heap.Array[j] = heap.Array[j], heap.Array[i]
		i = j
		j = i * 2 + 1
	}
}

// Insert 插入新的元素
func (heap *Heap) Insert(data tree.Item) {
	if heap.Size >= heap.Capcity {
		temp := make([]tree.Item, heap.Size)
		copy(temp, heap.Array)

		heap.Capcity += INCREMENT
		heap.Array = make([]tree.Item, heap.Capcity)
		copy(heap.Array, temp)
	}

	heap.Array[heap.Size] = data
	heap.SiftUp(heap.Size)
	heap.Size++
}

// Delete 删除元素
func (heap *Heap) Delete(pos int) {
	if heap.Empty() || pos < 0 || pos >= heap.Size {
		return
	}

	heap.Array[pos] = heap.Array[heap.Size - 1]
	heap.Size--
	heap.SiftDown(pos)
}

// Sort 堆排序
func (heap *Heap) Sort() {
	if heap.Empty() {
		return
	}

	fmt.Print("Sort: ")
	for heap.Size > 0 {
		fmt.Printf("%v ", heap.Array[0])
		heap.Delete(0)
	}
	fmt.Println()
}

// Traverse ...
func (heap *Heap) Traverse() {
	// fmt.Print("Traverse: ")
	for i := 0; i < heap.Size; i++ {
		fmt.Printf("%v ", heap.Array[i])
	}
	fmt.Println()
}

// FromHeap ...
func FromHeap(h *Heap) *Heap {
	return FromArray(h.Array, h.Size, h.Type)
}

// FromArray ...
func FromArray(a []tree.Item, n int, t HeapType) *Heap {
	heap := &Heap{Type: t, Size: n}

	heap.Capcity = DEFAULTSIZE
	for heap.Capcity < n {
		heap.Capcity += INCREMENT
	}
	heap.Array = make([]tree.Item, heap.Capcity)
	copy(heap.Array, a)

	for i := (heap.Size - 2) / 2; i >= 0; i-- {
		heap.SiftDown(i)
	}

	return heap
}

// HeapTest 堆测试
func HeapTest() {
	fmt.Println("HeapTest begin")
	a := []tree.Item{3, 6, 1, 6, 2, 5, 7, 12}
	heap1 := FromArray(a, len(a), Min)
	fmt.Print("heap1: ")
	heap1.Traverse()
	heap2 := FromHeap(heap1)
	fmt.Print("heap2: ")
	heap2.Traverse()

	heap1.Insert(4)
	fmt.Print("heap1: ")
	heap1.Traverse()
	heap1.Insert(1)
	fmt.Print("heap1: ")
	heap1.Traverse()

	heap1.Delete(1)
	fmt.Print("heap1: ")
	heap1.Traverse()

	heap1.Sort()
	
	fmt.Println("HeapTest end")
}