package usage

import (
	"fmt"
	"DA/4_queue/queue"
)

// Pos ...
type Pos int
const (
	// L - Left
	L Pos = iota
	// R - Right
	R
)

// Deque ...
type Deque struct {
	Left int
	Right int
	Buff [queue.SIZE]queue.Node
	Length int
}

// Empty ...
func (dq *Deque) Empty() bool {
	if dq.Left == dq.Right {
		fmt.Println("deque empty")
		return true
	}
	return false
}

// Full ...
func (dq *Deque) Full() bool {
	if (dq.Right + 1) % queue.SIZE == dq.Left {
		fmt.Println("deque full")
		return true
	}
	return false
}

// Clear ...
func (dq *Deque) Clear() {
	dq.Left = dq.Right
}

// TopAt ...
func (dq *Deque) TopAt(pos Pos) queue.Item {
	if dq.Empty() {
		return nil
	}

	var res queue.Item
	if pos == L {
		res = dq.Buff[dq.Left].Data
	} else {
		res = dq.Buff[(dq.Right - 1 + queue.SIZE) % queue.SIZE].Data
	}

	return res
}

// PushAt ...
func (dq *Deque) PushAt(pos Pos, data queue.Item) bool {
	if dq.Full() {
		return false
	}

	if pos == L {
		dq.Left = (dq.Left - 1 + queue.SIZE) % queue.SIZE
		dq.Buff[dq.Left].Data = data
	} else {
		dq.Buff[dq.Right].Data = data
		dq.Right = (dq.Right + 1) % queue.SIZE
	}

	return true
}

// PopAt ...
func (dq *Deque) PopAt(pos Pos) queue.Item {
	if dq.Empty() {
		return nil
	}

	var res queue.Item
	if pos == L {
		res = dq.Buff[dq.Left].Data
		dq.Left = (dq.Left + 1) % queue.SIZE
	} else {
		dq.Right = (dq.Right - 1 + queue.SIZE) % queue.SIZE
		res = dq.Buff[dq.Right]
	}

	return res
}

// Output ...
func (dq *Deque) Output(pos Pos) {
	if dq.Empty() {
		return
	}

	if pos == L {
		p := dq.Left
		for p != dq.Right {
			fmt.Printf("%v ", dq.Buff[p].Data)
			p = (p + 1) % queue.SIZE
		}
	} else {
		p := dq.Right
		for p != dq.Left {
			p = (p - 1 + queue.SIZE) % queue.SIZE
			fmt.Printf("%v ", dq.Buff[p].Data)
		}
	}

	fmt.Println()
}

// DqInit ...
func DqInit() *Deque{
	return &Deque{}
}

// DqTest ...
func DqTest() {
	fmt.Println("de queue test begin")
	fmt.Printf("de queue test end\n\n")
}