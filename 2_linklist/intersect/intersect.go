package intersect

import (
	"fmt"
	"DA/2_linklist/linklist"
)

func hasIntersection(ll1, ll2 *linklist.LList) bool {
	if ll1.Head == nil || ll1.Head.Next == nil || 
	   ll2.Head == nil || ll2.Head.Next == nil {
		return false
	}

	p := ll1.Head.Next
	for p.Next != nil {
		p = p.Next
	}

	q := ll2.Head.Next
	for q.Next != nil {
		q = q.Next
	}

	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", q)
	return p.Data == q.Data
}

// Test ...
func Test() {
	ll1 := linklist.Init()
	for i := 0; i < 3; i++ {
		ll1.Append(i)
	}
	ll2 := linklist.Init()
	for i := 0; i < 3; i++ {
		ll2.Append(i)
	}

	for i := 3; i < 5; i++ {
		ll1.Append(i)
		ll2.Append(i)
	}

	hasInter := hasIntersection(ll1, ll2)
	if hasInter {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Println()
}