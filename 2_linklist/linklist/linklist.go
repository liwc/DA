package linklist

import (
	"fmt"
	"DA/2_linklist/linknode"
)

// LList ...
type LList struct {
	Head *linknode.LNode
	Tail *linknode.LNode
	Length int
	PrtHead bool
	Duplicate bool
}

// Output ...
func (ll *LList) Output() {
	if ll.Head == nil {
		return;
	}

	prtLen := ll.Length 
	if ll.PrtHead {
		prtLen = ll.Length + 1
	}

	iter := ll.Head
	for i := 0; i < prtLen; i++ {
		if iter != nil {
			fmt.Printf("%v ", iter.Data)
			iter = iter.Next
		}
	}
	
	fmt.Println()
}

// Insert ...
func (ll *LList) Insert(pos int, data linknode.Item) {
	
}

// Delete ...
func (ll *LList) Delete(pos int) linknode.Item {
	return nil
}

// Update ...
func (ll *LList) Update(pos int, data linknode.Item) {
	
}

// Search ...
func (ll *LList) Search(data linknode.Item) int {
	index := -1

	return index
}

// Append ...
func (ll *LList) Append(data linknode.Item) {
	if !ll.Duplicate && ll.Tail.Data == data {
		return
	}

	ll.Length++

	node := &linknode.LNode{Data: data}
	ll.Tail.Next = node
	ll.Tail = node
}

// IgnoreHead ...
func (ll *LList) IgnoreHead() {
	ll.PrtHead = false

	ll.Tail.Next = ll.Head.Next
	ll.Head = ll.Tail.Next
}

// RecoverHead ...
func (ll *LList) RecoverHead() {
	ll.PrtHead = true

	node := &linknode.LNode{}
	node.Next = ll.Tail.Next
	ll.Tail.Next = node
	ll.Head = ll.Tail.Next
}

// Init ...
func Init() *LList {
	node := &linknode.LNode{}
	ll := &LList{
		Head: node, 
		Tail: node,
		Length: 0,
		PrtHead: true,
		Duplicate: true,
	}

	return ll
}

// Test ...
func Test () {
	ll := Init()
	for i := 0; i < 3; i++ {
		ll.Append(i)
	}
	ll.Output()

	ll.IgnoreHead()
	ll.Output()

	ll.RecoverHead()
	ll.Output()

	fmt.Println()
}