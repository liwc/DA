package linknode

// Item ...
type Item interface {}

// LNode ...
type LNode struct {
	Data Item
	Next *LNode

	Reserve1 Item
	Reserve2 Item
	Reserve3 Item
}
