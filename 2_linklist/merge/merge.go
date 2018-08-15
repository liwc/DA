package merge

import (
	"DA/2_linklist/linklist"
)

func merge1(ll1, ll2 *linklist.LList) *linklist.LList {
	if ll1.Head == nil || ll1.Head.Next == nil {
		return ll2
	}
	if ll2.Head == nil || ll2.Head.Next == nil {
		return ll1
	}

	ll := linklist.Init()
	ll.Duplicate = false

	p := ll1.Head.Next
	q := ll2.Head.Next
	for {
		if p == nil && q == nil {
			break
		} else if p == nil && q != nil {
			ll.Append(q.Data)
			q = q.Next
		} else if p != nil && q == nil {
			ll.Append(p.Data)
			p = p.Next
		} else {
			pData, _ := p.Data.(int)
			qData, _ := q.Data.(int)
			if pData < qData {
				ll.Append(p.Data)
				p = p.Next
			} else {
				ll.Append(q.Data)
				q = q.Next
			}
		}
	}
	
	return ll
}

func merge2(ll1, ll2 *linklist.LList) *linklist.LList {
	if ll1.Head == nil || ll1.Head.Next == nil {
		return ll2
	}
	if ll2.Head == nil || ll2.Head.Next == nil {
		return ll1
	}

	prev := ll1.Head
	p := ll1.Head.Next
	q := ll2.Head.Next
	for q != nil {
		if p == nil {
			ll1.Append(q.Data)
			q = q.Next
		} else {
			pData, _ := p.Data.(int)
			qData, _ := q.Data.(int)
			if pData < qData {
				prev = p
				p = p.Next
			} else if pData == qData {
				q = q.Next
			} else if pData > qData{
				rear := q.Next
				q.Next = p
				prev.Next = q
				prev = q
				q = rear
			}
		}
	}
	
	return ll1
}

// Test ...
func Test() {
	ll1 := linklist.Init()
	for i := 0; i < 10; i = i + 2 {
		ll1.Append(i)
	}
	ll1.Output()

	ll2 := linklist.Init()
	for i := 0; i < 10; i = i + 3 {
		ll2.Append(i)
	}
	ll2.Output()

	ll := merge2(ll1, ll2)
	ll.Output()
}