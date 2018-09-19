package usage

import (
	"fmt"
	"DA/3_stack/stack"
	"DA/7_graph/graph"
)

// Dijkstra 单源最短路径
type Dijkstra struct {
	NumV int
	Am *graph.AdjMatrix
	Pre []int
	Dis []int
	Find []bool
}

// Output ...
func (dj *Dijkstra) Output() {
	dj.Am.Output()
}

// SetEdge ...
func (dj *Dijkstra) SetEdge(tail, head, weight int) {
	dj.Am.SetEdge(tail, head, weight)
}

// DjImpl ...
func (dj *Dijkstra) DjImpl(v int) {
	dj.Find[v] = true
	for i := 0; i < dj.NumV; i++ {
		dj.Pre[i] = v
		dj.Dis[i] = dj.Am.Array[v][i]
	}

	for i := 1; i < dj.NumV; i++ {
		target := i
		d := graph.MAXWEIGHT
		//确定一个最短距离
		for j := 0; j < dj.NumV; j++ {
			if !dj.Find[j] && dj.Dis[j] < d {
				d = dj.Dis[j]
				target = j
			}
		}
		
		dj.Find[target] = true

		//更新剩余顶点的前驱和最短距离
		for k := 0; k < dj.NumV; k++ {
			if !dj.Find[k] {
				d = dj.Dis[target] + dj.Am.Array[target][k]
				if d < dj.Dis[k] {
					dj.Pre[k] = target
					dj.Dis[k] = d
				}
			}
		}
	}

	for i := 0; i < dj.NumV; i++ {
		if i == v {
			continue
		}

		os := stack.OsInit()
		for pos := i; pos != v; {
			os.Push(pos)
			pos = dj.Pre[pos]
		}

		fmt.Printf("[%d, %d]: ", v, i)
		for !os.Empty() {
			fmt.Printf("%v ", os.Pop())
		}
		fmt.Println()
	}
}

// DijkstraInit ...
func DijkstraInit(isWeighted, isDirected bool, numV int) *Dijkstra {
	dj := &Dijkstra{
		NumV: numV,
		Am: graph.AdjMatrixInit(isWeighted, isDirected, numV),
		Pre: make([]int, numV),
		Dis: make([]int, numV),
		Find: make([]bool, numV),
	}

	return dj
}

// DijkstraTest ...
func DijkstraTest() {
	fmt.Println("DijkstraTest begin")
	dj := DijkstraInit(true, true, 5)
	dj.SetEdge(0, 1, 10)
	dj.SetEdge(0, 3, 30)
	dj.SetEdge(0, 4, 90)
	dj.SetEdge(1, 2, 50)
	dj.SetEdge(2, 4, 10)
	dj.SetEdge(3, 2, 20)
	dj.SetEdge(3, 4, 60)
	dj.Output()
	dj.DjImpl(0)
	fmt.Println("DijkstraTest end")
}
