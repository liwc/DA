package graph

import (
	"fmt"
	"DA/3_stack/stack"
	"DA/4_queue/queue"
)

// Fs 优先遍历
type Fs struct {
	Am *AdjMatrix
	Al *AdjList
}

// Depth 深度优先
func (fs *Fs) Depth(start int) {
	numV := fs.Am.NumV
	if start >= numV {
		return
	}

	os := stack.OsInit()
	visited := make([]int, numV)
	
	for start < numV {
		if visited[start] == 0 {
			visited[start] = 1
			fmt.Printf("%d ", start)
		}

		tail := start
		head := 0
		for head < numV {
			if visited[head] == 0 && 
			   fs.Am.Array[tail][head] == 1 {
				os.Push(start)
				start = head
				break
			}

			head++
		}

		if head >= numV {
			if !os.Empty() {
				start = os.Pop().(int)
			} else {
				pos := 0
				for pos < numV {
					if visited[pos] == 0 {
						start = pos;
						break
					}

					pos++
				}

				if pos >= numV {
					break
				}
			}
		}
	}

	fmt.Println()
}

// Breadth 广度优先
func (fs *Fs) Breadth(start int) {
	numV := fs.Am.NumV
	if start >= numV {
		return
	}

	oq := queue.OqInit()
	visited := make([]int, numV)
	if visited[start] == 0 {
		visited[start] = 1
		fmt.Printf("%d ", start)
		oq.Push(start)
	}

	for !oq.Empty() {
		pos := oq.Pop().(int)
		p := fs.Al.Array[pos].Next
		for p != nil {
			if visited[p.Adjvtx] == 0 {
				visited[p.Adjvtx] = 1
				fmt.Printf("%d ", p.Adjvtx)
				oq.Push(p.Adjvtx)
			}

			p = p.Next
		}

		if oq.Empty() {
			for pos = 0; pos < numV; pos++ {
				if visited[pos] == 0 {
					visited[pos] = 1
					fmt.Printf("%d ", pos)
					oq.Push(pos)
				}
			}
		}
	}

	fmt.Println()
}

// FsTest ...
func FsTest() {
	fmt.Println("FsTest begin")

	am := AdjMatrixInit(false, true, 5)
	am.SetEdge(0, 1, 1)
	am.SetEdge(0, 2, 1)
	am.SetEdge(1, 3, 1)
	am.SetEdge(3, 0, 1)
	am.SetEdge(4, 2, 1)
	am.Output()

	al := AdjListInit(false, true, 5)
	al.SetEdge(0, 1, 1)
	al.SetEdge(0, 2, 1)
	al.SetEdge(1, 3, 1)
	al.SetEdge(3, 0, 1)
	al.SetEdge(4, 2, 1)
	al.Output()

	fs := &Fs{
		Am: am,
		Al: al,
	}

	fmt.Println("Depth:")
	for i := 0; i < fs.Am.NumV; i++ {
		fs.Depth(i)
	}

	fmt.Println("Breadth:")
	for i := 0; i < fs.Al.NumV; i++ {
		fs.Breadth(i)
	}

	fmt.Println("FsTest end")
}