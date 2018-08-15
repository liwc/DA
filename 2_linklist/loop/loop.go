package loop

import (
	"fmt"
	"DA/2_linklist/linklist"
)

func hasloop(llist *linklist.LList) bool {
	p := llist.Head
	q := llist.Head

	for p != nil && q != nil {
		p = p.Next
		q = q.Next
		if q != nil {
			q = q.Next
		}

		if p == q {
			return true
		}
	}

	return false
}

// Test ...
func Test() {
	ll := linklist.Init()
	for i := 0; i < 10; i++ {
		ll.Append(i)
	}
	//ll.Tail.Next = ll.Head
	
	has := hasloop(ll)
	if has {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println()
}