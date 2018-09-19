package usage

import (
	"fmt"
	"DA/4_queue/queue"
	"DA/7_graph/graph"
)

// Topo 拓扑排序
type Topo struct {
	Am *graph.AdjMatrix
	InDegree []int
}

// Output ...
func (topo *Topo) Output() {
	topo.Am.Output()
}

// SetEdge ...
func (topo *Topo) SetEdge(tail, head, weight int) {
	if topo.Am.SetEdge(tail, head, weight) {
		topo.InDegree[head]++
	}
}

// Sort ...
func (topo *Topo) Sort() {
	numV := topo.Am.NumV
	if numV == 0 {
		return
	}

	oq := queue.OqInit()
	for i := 0; i < numV; i++ {
		if topo.InDegree[i] == 0 {
			oq.Push(i)
		}
	}

	visited := make([]int, numV)
	for !oq.Empty() {
		tail := oq.Pop().(int)
		fmt.Printf("%d\t", tail)
		visited[tail] = 1

		for head := 0; head < numV; head++ {
			if topo.Am.Array[tail][head] == 1 {
				topo.InDegree[head]--
				if topo.InDegree[head] == 0 {
					oq.Push(head)
				}
			}
		}
	}
	fmt.Println()

	for i := 0; i < numV; i++ {
		if visited[i] == 0 {
			fmt.Println("has loop")
			break
		}
	}
}

// TopoInit ...
func TopoInit(isWeighted, isDirected bool, numV int) *Topo {
	topo := &Topo{
		Am: graph.AdjMatrixInit(isWeighted, isDirected, numV),
		InDegree: make([]int, numV),
	}

	return topo
}

// TopoTest ...
func TopoTest() {
	fmt.Println("TopoTest begin")
	topo := TopoInit(false, true, 9)
	topo.Output()
	topo.SetEdge(0, 2, 1)
	topo.SetEdge(0, 7, 1)
	topo.SetEdge(1, 2, 1)
	topo.SetEdge(1, 3, 1)
	topo.SetEdge(1, 4, 1)
	topo.SetEdge(2, 3, 1)
	topo.SetEdge(3, 5, 1)
	topo.SetEdge(3, 6, 1)
	topo.SetEdge(4, 5, 1)
	topo.SetEdge(7, 8, 1)
	topo.SetEdge(8, 6, 1)
	topo.Output()

	topo.Sort()
	fmt.Println("TopoTest end")
}