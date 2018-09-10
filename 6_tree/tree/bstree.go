package tree

import (
	"fmt"
	"DA/3_stack/stack"
	"DA/4_queue/queue"
)

// BSTree 二叉搜索树
type BSTree struct {
	Root *Node
	Size int
}

// Empty 判断树空
func (bstree *BSTree) Empty() bool {
	if bstree.Root == nil {
		return true
	}
	
	return false
}

// Insert 插入节点
func (bstree *BSTree) Insert(data Item) bool {
	if bstree.Empty() {
		bstree.Root = &Node{Data: data}
		bstree.Size++
		return true
	}

	p := bstree.Root
	for p != nil {
		pData, _ := p.Data.(int)
		qData, _ := data.(int)
		if pData == qData {
			return false
		} else if pData > qData {
			if p.Left != nil {
				p = p.Left
			} else {
				p.Left = &Node{Data: data, Parent: p}
				break
			}
		} else if pData < qData {
			if p.Right != nil {
				p = p.Right
			} else {
				p.Right = &Node{Data: data, Parent: p}
				break
			}
		}
	}

	bstree.Size++
	return true
}

// InorderTraverse 中序遍历
func (bstree *BSTree) InorderTraverse() {
	if bstree.Root != nil {
		fmt.Print("Inorder Traverse: ")
		inorderTraverseImpl(bstree.Root)
		fmt.Println()
	}
}

// Leaf 叶子节点数
func (bstree *BSTree) Leaf() int {
	leafNum := 0
	if !bstree.Empty() {
		oq := queue.OqInit()
		oq.Push(bstree.Root)
		for !oq.Empty() {
			p, _ := oq.Pop().(*Node)
			if p.Left == nil && p.Right == nil {
				leafNum++
			}

			if p.Left != nil {
				oq.Push(p.Left)
			}

			if p.Right != nil {
				oq.Push(p.Right)
			}
		}
	}
	return leafNum
}

// Search 查找
func (bstree *BSTree) Search(data Item) bool {
	if !bstree.Empty() {
		oq := queue.OqInit()
		oq.Push(bstree.Root)
		for !oq.Empty() {
			p, _ := oq.Pop().(*Node)
			if p.Data == data {
				return true
			}

			if p.Left != nil {
				oq.Push(p.Left)
			}

			if p.Right != nil {
				oq.Push(p.Right)
			}
		}
	}

	return false
}

// Parent 获取父节点
func (bstree *BSTree) Parent(data Item) Item {
	if bstree.Search(data) {
		oq := queue.OqInit()
		oq.Push(bstree.Root)
		for !oq.Empty() {
			p, _ := oq.Pop().(*Node)
			if p.Data == data {
				if p.Parent == nil {
					return nil
				}
				return p.Parent.Data
			}

			if p.Left != nil {
				oq.Push(p.Left)
			}

			if p.Right != nil {
				oq.Push(p.Right)
			}
		}
	}

	return nil
}

// Predecessor 获取前驱
func (bstree *BSTree) Predecessor(data Item) Item {
	var q *Node
	if bstree.Search(data) {
		p := bstree.Root
		os := stack.OsInit()
		for p != nil || !os.Empty() {
			if p != nil {
				if p.Right != nil {
					os.Push(p.Right)
				}
				os.Push(p)
				p = p.Left
			} else {
				p, _ = os.Pop().(*Node)
				if p.Data == data {
					break
				}

				q = p
				p = p.Right
			}
		}
	}

	if q == nil {
		return nil
	}
	return q.Data
}

// Successor 获取后继
func (bstree *BSTree) Successor(data Item) Item {
	var q *Node
	if bstree.Search(data) {
		p := bstree.Root
		os := stack.OsInit()
		for p != nil || !os.Empty() {
			if p != nil {
				if p.Right != nil {
					os.Push(p.Right)
				}
				os.Push(p)
				p = p.Left
			} else {
				p, _ = os.Pop().(*Node)
				if q != nil && q.Data == data {
					q = p
					break
				}

				q = p
				p = p.Right
			}
		}
	}

	if q == nil || q.Data == data {
		return nil
	}
	return q.Data
}

// Minimum 获取最小节点
func (bstree *BSTree) Minimum() Item {
	if !bstree.Empty() {
		p := bstree.Root
		for p.Left != nil {
			p = p.Left
		}
		return p.Data
	}

	return nil
}

// Maximum 获取最大节点
func (bstree *BSTree) Maximum() Item {
	if !bstree.Empty() {
		p := bstree.Root
		for p.Right != nil {
			p = p.Right
		}
		return p.Data
	}

	return nil
}

// Delete 删除节点
func (bstree *BSTree) Delete(data Item) bool {
	if bstree.Search(data) {
		var p *Node

		oq := queue.OqInit()
		oq.Push(bstree.Root)
		for !oq.Empty() {
			p, _ = oq.Pop().(*Node)
			if p.Data == data {
				break
			}

			if p.Left != nil {
				oq.Push(p.Left)
			}

			if p.Right != nil {
				oq.Push(p.Right)
			}
		}
		
		if p.Left == nil {
			// 无左子树
			if p.Parent == nil {
				bstree.Root = p.Right
			} else if p.Parent.Left == p {
				p.Parent.Left = p.Right
			} else if p.Parent.Right == p {
				p.Parent.Right = p.Right
			}

			if p.Right != nil {
				p.Right.Parent = p.Parent
			}
		} else {
			// 有左子树
			q := p.Left
			for q.Right != nil {
				q = q.Right
			}

			if q.Parent.Left == q {
				q.Parent.Left = q.Left
			} else if q.Parent.Right == q {
				q.Parent.Right = q.Left
			}
			if q.Left != nil {
				q.Left.Parent = q.Parent
			}

			q.Left = p.Left
			if p.Left != nil {
				p.Left.Parent = q
			}

			q.Right = p.Right
			if p.Right != nil {
				p.Right.Parent = q
			}
			
			if p.Parent == nil {
				bstree.Root = q
			} else if p.Parent.Left == p {
				p.Parent.Left = q
			} else if p.Parent.Right == p {
				p.Parent.Right = q
			}
			q.Parent = p.Parent
		}

		bstree.Size--
		return true
	}

	return false
}

// BSTreeInit ...
// 	     4
// 	   /   \
// 	  2		5
//  /   \	  \
// 1	 3	   7
//            /
//          6
func BSTreeInit() *BSTree {
	bstree := &BSTree{}
	bstree.Insert(4)
	bstree.Insert(2)
	bstree.Insert(1)
	bstree.Insert(3)
	bstree.Insert(5)
	bstree.Insert(7)
	bstree.Insert(6)
	return bstree
}

// BSTreeTest ...
func BSTreeTest() {
	fmt.Println("BSTreeTest begin")

	bstree := BSTreeInit()
	bstree.InorderTraverse()
	fmt.Printf("Leaf num: %d\n", bstree.Leaf())
	fmt.Printf("Search 3: %v\n", bstree.Search(3))
	fmt.Printf("Search 8: %v\n", bstree.Search(8))
	fmt.Printf("Parent of 3: %v\n", bstree.Parent(3))
	fmt.Printf("Parent of 4: %v\n", bstree.Parent(4))
	fmt.Printf("Predecessor of 1: %v\n", bstree.Predecessor(1))
	fmt.Printf("Predecessor of 7: %v\n", bstree.Predecessor(7))
	fmt.Printf("Successor of 1: %v\n", bstree.Successor(1))
	fmt.Printf("Successor of 7: %v\n", bstree.Successor(7))
	fmt.Printf("Minimum: %v\n", bstree.Minimum())
	fmt.Printf("Maximum: %v\n", bstree.Maximum())

	bstree.Delete(5)
	bstree.InorderTraverse()
	bstree.Delete(4)
	bstree.InorderTraverse()
	fmt.Println("BSTreeTest end")
}