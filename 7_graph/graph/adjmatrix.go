package graph

import (
	"fmt"
)

// MAXWEIGHT 最大权值
const MAXWEIGHT = 100

// AdjMatrix 邻接矩阵
type AdjMatrix struct {
	IsWeighted bool
	IsDirected bool
	NumV int
	NumE int
	Array [][]int
}

// Output ...
func (am *AdjMatrix) Output() {
	for i := 0; i < am.NumV; i++ {
		for j := 0; j < am.NumV; j++ {
			fmt.Printf("%v\t", am.Array[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Check ...
func (am *AdjMatrix) Check(tail, head, weight int) bool {
	if tail >= 0 && tail < am.NumV &&
	   head >= 0 && head < am.NumV &&
	   weight >= 0 && weight <= MAXWEIGHT {
		return true
	}

	return false
}

// SetEdge ...
func (am *AdjMatrix) SetEdge(tail, head, weight int) bool {
	if !am.Check(tail, head, weight) {
		return false
	}

	if am.IsWeighted {
		if am.IsDirected {
			am.Array[tail][head] = weight
		} else {
			am.Array[tail][head] = weight
			am.Array[head][tail] = weight
		}
	} else {
		if am.IsDirected {
			am.Array[tail][head] = 1
		} else {
			am.Array[tail][head] = 1
			am.Array[head][tail] = 1
		}
	}

	return true
}

// AdjMatrixInit ...
func AdjMatrixInit(isWeighted, isDirected bool, numV int) *AdjMatrix{
	am := &AdjMatrix{
		IsWeighted: isWeighted,
		IsDirected: isDirected,
		NumV: numV,
	}

	weight := 0
	if isWeighted {
		weight = MAXWEIGHT
	}

	am.Array = make([][]int, numV)
	for i := 0; i < numV; i++ {
		am.Array[i] = make([]int, numV)
		for j := 0; j < numV; j++ {
			am.Array[i][j] = weight
		}
	}

	return am
}

// AdjMatrixTest ...
func AdjMatrixTest() {
	fmt.Println("AdjMatrixTest begin")
	am := AdjMatrixInit(false, false, 5)
	am.SetEdge(0, 1, 1)
	am.SetEdge(0, 3, 1)
	am.SetEdge(0, 4, 1)
	am.SetEdge(1, 2, 1)
	am.SetEdge(2, 3, 1)
	am.SetEdge(3, 4, 1)
	am.Output()
	am.SetEdge(1, 3, 1)
	am.Output()

	am2 := AdjMatrixInit(true, true, 5)
	am2.SetEdge(0, 1, 15)
	am2.SetEdge(0, 4, 78)
	am2.SetEdge(1, 3, 20)
	am2.SetEdge(2, 1, 22)
	am2.SetEdge(3, 2, 34)
	am2.SetEdge(4, 2, 56)
	am2.Output()
	am2.SetEdge(1, 0, 20)
	am2.Output()
	fmt.Println("AdjMatrixTest end")
}