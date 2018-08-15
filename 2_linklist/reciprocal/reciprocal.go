package reciprocal

import (
	"fmt"
	"DA/2_linklist/linklist"
)

func reciprocalK(ll *linklist.LList, k int) int {
	if ll.Head == nil || ll.Head.Next == nil {
		return -1
	}

	p := ll.Head
	for i := 0; i < k; i++ {
		p = p.Next
		if p == nil {
			return -2
		}
	}

	q := ll.Head.Next
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	value, ok := q.Data.(int)
	if ok {
		return value
	}

	return 0
}

// Test ...
func Test() {
	ll := linklist.Init()
	for i := 0; i < 10; i++ {
		ll.Append(i)
	}
	ll.Output()

	fmt.Println(reciprocalK(ll, 5))
	fmt.Println()
}