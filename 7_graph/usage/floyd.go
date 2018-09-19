package usage

import (
	"fmt"
	"DA/3_stack/stack"
	"DA/7_graph/graph"
)

// Floyd 点对之间最短距离
type Floyd struct {
	NumV int
	Am *graph.AdjMatrix
	Dis [][]int
	Pre [][]int
}

// Output ...
func (floyd *Floyd) Output() {
	floyd.Am.Output()
}

// SetEdge ...
func (floyd *Floyd) SetEdge(tail, head, weight int) {
	floyd.Am.SetEdge(tail, head, weight)
}

// FloydImpl ...
func (floyd *Floyd) FloydImpl() {
	floyd.Dis = floyd.Am.Array
	for i := 0; i < floyd.NumV; i++ {
		for j := 0; j < floyd.NumV; j++ {
			if floyd.Dis[i][j] == graph.MAXWEIGHT {
				floyd.Pre[i][j] = -1
			} else {
				floyd.Pre[i][j] = i
			}
		}
	}

	for i := 0; i < floyd.NumV; i++ {
		for j := 0; j < floyd.NumV; j++ {
			if i == j {
				continue
			}
			
			for k := 0; k < floyd.NumV; k++ {
				d := floyd.Am.Array[i][k] + floyd.Am.Array[k][j]
				if d < floyd.Am.Array[i][j] {
					floyd.Dis[i][j] = d
					floyd.Pre[i][j] = k
				}
			}
		}
	}

	fmt.Println("Dis...")
	for i := 0; i < floyd.NumV; i++ {
		for j := 0; j <floyd.NumV; j++ {
			fmt.Printf("%v\t", floyd.Dis[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Pre...")
	for i := 0; i < floyd.NumV; i++ {
		for j := 0; j <floyd.NumV; j++ {
			fmt.Printf("%v\t", floyd.Pre[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < floyd.NumV; i++ {
		for j := 0; j < floyd.NumV; j++ {
			if i == j {
				continue
			}

			fmt.Printf("[%d, %d]: %d, path: ", i, j, floyd.Dis[i][j])
			os := stack.OsInit()
			os.Push(j)
			for k := floyd.Pre[i][j]; k != i; {
				os.Push(k)
				k = floyd.Pre[i][k]
			}
			os.Push(i)
			
			for !os.Empty() {
				fmt.Printf("%v ", os.Pop())
			}
			fmt.Println()
		}
	}
}

// FloydInit ...
func FloydInit(isWeighted, isDirected bool, numV int) *Floyd {
	floyd := &Floyd{
		NumV: numV,
		Am: graph.AdjMatrixInit(isWeighted, isDirected, numV),
	}

	floyd.Dis = make([][]int, numV)
	floyd.Pre = make([][]int, numV)
	for i := 0; i < numV; i++ {
		floyd.Dis[i] = make([]int, numV)
		floyd.Pre[i] = make([]int, numV)
	}

	return floyd
}

// FloydTest ...
func FloydTest() {
	fmt.Println("FloydTest begin")
	floyd := FloydInit(true, true, 3)
	floyd.SetEdge(0, 2, 2)
	floyd.SetEdge(1, 0, 5)
	floyd.SetEdge(1, 2, 8)
	floyd.SetEdge(2, 1, 3)
	floyd.Output()
	floyd.FloydImpl()
	fmt.Println("FloydTest end")
}