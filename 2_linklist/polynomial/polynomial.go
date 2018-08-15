package polynomial

import (
	"fmt"
	"DA/2_linklist/linknode"
	"DA/2_linklist/linklist"
)

func dupNode(coef, expn linknode.Item) *linknode.LNode {
	node := &linknode.LNode {
		Reserve1: coef,
		Reserve2: expn,
	}

	return node
}

func dupList(src *linklist.LList) *linklist.LList {	
	if src == nil || src.Head == nil {
		return src
	}

	dst := linklist.Init()
	iter := src.Head.Next
	for iter != nil {
		dst.Tail.Next = dupNode(iter.Reserve1, iter.Reserve2)
		dst.Tail = dst.Tail.Next

		iter = iter.Next
	}

	dst.Tail.Next = nil
	return dst
}

func add(ll1, ll2 *linklist.LList) *linklist.LList {
	if ll2 == nil || ll2.Head == nil {
		return ll1
	}
	if ll1 == nil || ll1.Head == nil {
		return ll2
	}

	ll := linklist.Init()

	p := dupList(ll1).Head.Next
	q := dupList(ll2).Head.Next
	for {
		if p == nil && q == nil {
			break
		} else if p != nil && q == nil {
			ll.Tail.Next = p
			ll.Tail = ll.Tail.Next
			p = p.Next
		} else if p == nil && q != nil {
			ll.Tail.Next = q
			ll.Tail = ll.Tail.Next
			q = q.Next
		} else {
			pExpn, _ := p.Reserve2.(int)
			qExpn, _ := q.Reserve2.(int)
			if pExpn > qExpn {
				ll.Tail.Next = p
				ll.Tail = ll.Tail.Next
				p = p.Next
			} else if pExpn < qExpn {
				ll.Tail.Next = q
				ll.Tail = ll.Tail.Next
				q = q.Next
			} else {
				pCoef, _ := p.Reserve1.(int)
				qCoef, _ := q.Reserve1.(int)
				p.Reserve1 = pCoef + qCoef

				ll.Tail.Next = p
				ll.Tail = ll.Tail.Next
				p = p.Next
				q = q.Next
			}
		}
	}

	ll.Tail.Next = nil
	return ll
}

func sub(ll1, ll2 *linklist.LList) *linklist.LList {
	if (ll2 == nil || ll2.Head == nil) {
		return ll1
	}
	if (ll1 == nil || ll1.Head == nil) {
		ll := dupList(ll2)
		q := ll.Head.Next
		for q != nil {
			qCoef, _ := q.Reserve1.(int)
			q.Reserve1 = qCoef * (-1)
			q = q.Next
		}

		return ll
	}

	ll := linklist.Init()

	p := dupList(ll1).Head.Next
	q := dupList(ll2).Head.Next
	for {
		if p == nil && q == nil {
			break
		} else if p != nil && q == nil {
			ll.Tail.Next = p
			ll.Tail = ll.Tail.Next
			p = p.Next
		} else if p == nil && q != nil {
			qCoef, _ := q.Reserve1.(int)
			q.Reserve1 = -1 * qCoef
			ll.Tail.Next = q
			ll.Tail = ll.Tail.Next
			q = q.Next
		} else {
			pExpn, _ := p.Reserve2.(int)
			qExpn, _ := q.Reserve2.(int)
			if pExpn > qExpn {
				ll.Tail.Next = p
				ll.Tail = ll.Tail.Next
				p = p.Next
			} else if pExpn < qExpn {
				qCoef, _ := q.Reserve1.(int)
				q.Reserve1 = -1 * qCoef
				ll.Tail.Next = q
				ll.Tail = ll.Tail.Next
				q = q.Next
			} else {
				pCoef, _ := p.Reserve1.(int)
				qCoef, _ := q.Reserve1.(int)
				p.Reserve1 = pCoef - qCoef

				ll.Tail.Next = p
				ll.Tail = ll.Tail.Next
				p = p.Next
				q = q.Next
			}
		}
	}

	ll.Tail.Next = nil
	return ll
}

func mul(ll1, ll2 *linklist.LList) *linklist.LList {
	if ll2 == nil || ll2.Head == nil {
		return ll1
	}
	if ll1 == nil || ll1.Head == nil {
		return ll2
	}

	ll := linklist.Init()
	p := ll1.Head.Next
	for p != nil {
		lltemp := linklist.Init()
		pCoef, _ := p.Reserve1.(int)
		pExpn, _ := p.Reserve2.(int)

		q := ll2.Head.Next
		for q != nil {
			qCoef, _ := q.Reserve1.(int)
			qExpn, _ := q.Reserve2.(int)

			node := &linknode.LNode {
				Reserve1: pCoef * qCoef,
				Reserve2: pExpn + qExpn,
			}
			lltemp.Tail.Next = node
			lltemp.Tail = lltemp.Tail.Next

			q = q.Next
		}

		lltemp.Tail.Next = nil
		output(lltemp)
		ll = add(ll, lltemp)
		output(ll)

		p = p.Next
	}

	return ll
}

func output(ll *linklist.LList) {
	if ll == nil || ll.Head == nil || ll.Head.Next == nil {
		return
	}

	p := ll.Head.Next
	fmt.Printf("%dx%d", p.Reserve1, p.Reserve2)

	for p.Next != nil {
		fmt.Print(" + ")
		fmt.Printf("%dx%d", p.Next.Reserve1, p.Next.Reserve2)
		p = p.Next
	}

	fmt.Println()
}

// Test ...
func Test() {
	// ll1
	ll1 := linklist.Init()
	node11 := &linknode.LNode{
		Reserve1: 3,
		Reserve2: 2,
	}
	ll1.Tail.Next = node11
	ll1.Tail = node11

	node12 := &linknode.LNode{
		Reserve1: 2,
		Reserve2: 1,
	}
	ll1.Tail.Next = node12
	ll1.Tail = node12

	node13 := &linknode.LNode{
		Reserve1: -1,
		Reserve2: 0,
	}
	ll1.Tail.Next = node13
	ll1.Tail = node13

	// ll2
	ll2 := linklist.Init()
	node21 := &linknode.LNode{
		Reserve1: 2,
		Reserve2: 3,
	}
	ll2.Tail.Next = node21
	ll2.Tail = node21

	node22 := &linknode.LNode{
		Reserve1: -5,
		Reserve2: 1,
	}
	ll2.Tail.Next = node22
	ll2.Tail = node22

	node23 := &linknode.LNode{
		Reserve1: 4,
		Reserve2: 0,
	}
	ll2.Tail.Next = node23
	ll2.Tail = node23	

	// add
	output(ll1)
	output(ll2)
	lladd := add(ll1, ll2)
	fmt.Print("add: ")
	output(lladd)
	fmt.Println()

	// sub
	output(ll1)
	output(ll2)
	llsub := sub(ll1, ll2)
	fmt.Print("sub: ")
	output(llsub)
	fmt.Println()

	// mul
	output(ll1)
	output(ll2)
	llmul := mul(ll1, ll2)
	fmt.Print("mul: ")
	output(llmul)
	
	fmt.Println()
}