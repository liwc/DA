package usage

import (
	"fmt"
	"DA/4_queue/queue"
)

// PriorQueue ...
type PriorQueue struct {
	Length int
	Buff [queue.SIZE]queue.Node
}

// Empty ...
func (pq *PriorQueue) Empty() bool {
	if pq.Length == 0 {
		fmt.Println("queue empty")
		return true
	}
	return false
}

// Full ...
func (pq *PriorQueue) Full() bool {
	if pq.Length == queue.SIZE {
		fmt.Println("queue full")
		return true
	}
	return false
}

// Clear ...
func (pq *PriorQueue) Clear() {
	pq.Length = 0
}

// Push ...
func (pq *PriorQueue) Push(node queue.Node) bool {
	var i = 0
	for ; i < pq.Length; i++ {
		if pq.Buff[i].Reserve1 == node.Reserve1 {
			pq.Buff[i].Reserve2 = node.Reserve2
			return true
		}
	}

	if pq.Full() {
		return false
	} 

	pq.Buff[pq.Length] = node
	pq.Length++
	return true
}

// Pop ...
func (pq *PriorQueue) Pop() queue.Node {
	pos := pq.Find()
	if pos == -1 {
		return queue.Node{}
	}

	node := queue.Node{
		Reserve1: pq.Buff[pos].Reserve1,
		Reserve2: pq.Buff[pos].Reserve2,
	}

	pq.Length--
	for i := pos; i < pq.Length; i++ {
		pq.Buff[i] = pq.Buff[i + 1]
	}
	return node
}

// Find ...
func (pq *PriorQueue) Find() int {
	if pq.Empty() {
		return -1
	}

	var pos = 0
	for i := 1; i < pq.Length; i++ {
		prior1, _ := pq.Buff[pos].Reserve2.(int)
		prior2, _ := pq.Buff[i].Reserve2.(int)
		if prior1 > prior2 {
			pos = i
		}
	}

	return pos
}

// Output ...
func (pq *PriorQueue) Output() {
	if pq.Empty() {
		return
	}

	fmt.Printf("id\tprior\n")
	for i := 0; i < pq.Length; i++ {
		fmt.Printf("%v\t%v\n", pq.Buff[i].Reserve1, pq.Buff[i].Reserve2)
	}

	// fmt.Println()
}

// PqInit ...
func PqInit() *PriorQueue {
	pq := &PriorQueue{}
	return pq
}

// PqTest 优先队列
func PqTest() {
	fmt.Println("prior queue test begin")
	pq := PqInit()
	pq.Output()

	pq.Push(queue.Node{Reserve1: 0, Reserve2: 2,})
	pq.Push(queue.Node{Reserve1: 1, Reserve2: 1,})
	pq.Push(queue.Node{Reserve1: 2, Reserve2: 0,})
	pq.Push(queue.Node{Reserve1: 2, Reserve2: 3,})
	pq.Output()

	pq.Pop()
	pq.Output()
	
	fmt.Println()
	fmt.Printf("prior queue test end\n\n")
}