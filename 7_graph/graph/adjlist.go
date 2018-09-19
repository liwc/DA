package graph

import (
	"fmt"
)

// Edge 边
type Edge struct {
	Adjvtx int
	Weight int
	Next *Edge
}

// Vertex 顶点
type Vertex struct {
	ID int
	Next *Edge
}

// AdjList 邻接表
type AdjList struct {
	IsWeighted bool
	IsDirected bool
	NumV int
	NumE int
	Array []Vertex
}

// Output ...
func (al *AdjList) Output() {
	for i := 0; i < al.NumV; i++ {
		p := al.Array[i].Next
		for p != nil {
			fmt.Printf("[%d, %d] = %d\t", i, p.Adjvtx, p.Weight)
			p = p.Next
		}
		fmt.Println()
	}

	fmt.Println()
}

// Check ...
func (al *AdjList) Check(tail, head, weight int) bool {
	if tail >= 0 && tail < al.NumV &&
	   head >= 0 && head < al.NumV &&
	   weight >= 0 && weight <= MAXWEIGHT {
		return true
	}

	return false
}

// SetEdge ...
func (al *AdjList) SetEdge(tail, head, weight int) bool {
	if !al.Check(tail, head, weight) {
		return false
	}

	al.setedge(tail, head, weight)
	if !al.IsDirected {
		al.setedge(head, tail, weight)
	}

	return true
}

func (al *AdjList) setedge(tail, head, weight int) {
	if al.Array[tail].Next == nil {
		r := &Edge{
			Adjvtx: head,
			Weight: weight,
			Next: nil,
		}
		al.Array[tail].Next = r
	} else {
		p := al.Array[tail].Next
		q := al.Array[tail].Next
		for p != nil && p.Adjvtx < head {
			q = p
			p = p.Next
		}

		if p != nil && p.Adjvtx == head {
			p.Weight = weight
		} else {
			r := &Edge{
				Adjvtx: head,
				Weight: weight,
				Next: p,
			}

			if p != al.Array[tail].Next {
				q.Next = r
			} else {
				al.Array[tail].Next = r
			}
		}
	}
}

// AdjListInit ...
func AdjListInit(isWeighted, isDirected bool, numV int) *AdjList {
	al := &AdjList{
		IsWeighted: isWeighted,
		IsDirected: isDirected,
		NumV: numV,
	}

	al.Array = make([]Vertex, numV)
	for i := 0; i < numV; i++ {
		al.Array[i].ID = i
		al.Array[i].Next = nil
	}

	return al
}

// AdjListTest ...
func AdjListTest() {
	fmt.Println("AdjListTest begin")
	al := AdjListInit(false, false, 5)
	al.SetEdge(0, 1, 15)
	al.SetEdge(0, 4, 78)
	al.SetEdge(1, 2, 22)
	al.SetEdge(1, 3, 20)
	al.SetEdge(2, 3, 34)
	al.SetEdge(2, 4, 56)
	al.Output()
	al.SetEdge(0, 1, 20)
	al.Output()
	fmt.Println("AdjListTest end")
}