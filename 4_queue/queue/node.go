package queue

// Item ...
type Item interface {}

// Node ...
type Node struct {
	Data Item
	Next *Node
	Prev *Node

	Reserve1 Item
	Reserve2 Item
	Reserve3 Item
}

// SIZE ...
const SIZE = 5