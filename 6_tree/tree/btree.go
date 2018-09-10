package tree

import (
	"fmt"
	"DA/3_stack/stack"
	"DA/4_queue/queue"
)

// BTree 二叉树
type BTree struct {
	Root *Node
	Size int
}

// PreorderTraverse ...
func (btree *BTree) PreorderTraverse() {
	if btree.Root != nil {
		fmt.Print("Preorder Traverse: ")
		preorderTraverseImpl(btree.Root)
		fmt.Println()
	}
}

// InorderTraverse ...
func (btree *BTree) InorderTraverse() {
	if btree.Root != nil {
		fmt.Print("Inorder Traverse: ")
		inorderTraverseImpl(btree.Root)
		fmt.Println()
	}
}

// PostorderTraverse ...
func (btree *BTree) PostorderTraverse() {
	if btree.Root != nil {
		fmt.Print("Postorder Traverse: ")
		postorderTraverseImpl(btree.Root)
		fmt.Println()
	}
}

// LevelTraverse ...
func (btree *BTree) LevelTraverse() {
	p := btree.Root
	if p != nil {
		fmt.Print("Level Traverse: ")
		oq := queue.OqInit()
		oq.Push(p)
		for !oq.Empty() {
			node, _ := oq.Pop().(*Node)
			fmt.Printf("%v ", node.Data)

			if node.Left != nil {
				oq.Push(node.Left)
			}

			if node.Right != nil {
				oq.Push(node.Right)
			}
		}

		fmt.Println()
	}
}

func getHeightImpl(node *Node, level int, height *int) {
	if node != nil {
		if *height < level {
			*height = level
		}

		getHeightImpl(node.Left, level + 1, height)
		getHeightImpl(node.Right, level + 1, height)
	}
}

// GetHeight ...
func (btree *BTree) GetHeight() int {
	height := 0
	if btree.Root != nil {
		getHeightImpl(btree.Root, 1, &height)
	}

	return height
}

// GetParent ...
func (btree *BTree) GetParent(data Item) *Node {
	p := btree.Root
	if p != nil {
		oq := queue.OqInit()
		oq.Push(p)
		for !oq.Empty() {
			node, _ := oq.Pop().(*Node)
			if node.Data == data {
				return node.Parent
			}

			if node.Left != nil {
				oq.Push(node.Left)
			}

			if node.Right != nil {
				oq.Push(node.Right)
			}
		}
	}

	return nil
}

// GetParent2 ...
func (btree *BTree) GetParent2(data Item) *Node {
	p := btree.Root
	if p != nil {
		if p.Data == data {
			return nil
		}

		os := stack.OsInit()
		for p != nil || !os.Empty() {
			if p != nil {
				if (p.Left != nil && p.Left.Data == data) ||
				   (p.Right != nil && p.Right.Data == data) {
					return p
				}

				os.Push(p)
				p = p.Left
			} else {
				p = os.Pop().(*Node)
				p = p.Right
			}
		}
	}

	return nil
}

// PreorderWithoutRecursion ...
func (btree *BTree) PreorderWithoutRecursion() {
	p := btree.Root
	if p != nil {
		fmt.Print("Preorder Without Recursion: ")

		os := stack.OsInit()
		for p != nil || !os.Empty() {
			if p != nil {
				fmt.Printf("%v ", p.Data)
				
				if p.Right != nil {
					os.Push(p.Right)
				}

				p = p.Left
			} else {
				p = os.Pop().(*Node)
			}
		}

		fmt.Println()
	}
}

// InorderWithoutRecursion ...
func (btree *BTree) InorderWithoutRecursion() {
	p := btree.Root
	if p != nil {
		fmt.Print("Inorder Without Recursion: ")

		os := stack.OsInit()
		for p != nil || !os.Empty() {
			if p != nil {
				os.Push(p)
				p = p.Left
			} else {
				p = os.Pop().(*Node)
				fmt.Printf("%v ", p.Data)
				p = p.Right
			}
		}

		fmt.Println()
	}
}

// PostorderWithoutRecursion ...
func (btree *BTree) PostorderWithoutRecursion() {
	p := btree.Root
	if p != nil {
		fmt.Print("Postorder Without Recursion: ")

		os := stack.OsInit()
		for p != nil || !os.Empty() {
			if p != nil {
				if (p.Right != nil && p.Right.Reserve1 == 1) ||
				   (p.Left != nil && p.Left.Reserve1 == 1) {
					fmt.Printf("%v ", p.Data)
					p.Reserve1 = 1
					p, _ = os.Pop().(*Node)
				} else {
					os.Push(p)
					if p.Right != nil {
						os.Push(p.Right)
					}
					p = p.Left
				}
			} else {
				p = os.Pop().(*Node)
				if p.Left == nil && p.Right == nil {
					fmt.Printf("%v ", p.Data)
					p.Reserve1 = 1
					p, _ = os.Pop().(*Node)
				} else {
					os.Push(p)
					if p.Right != nil {
						os.Push(p.Right)
					}
					p = p.Left
				}
			}
		}

		fmt.Println()
	}
}

// BTreeInit ...
// 	     1
// 	   /   \
// 	  2		3
//  /   \	  \
// 4	 5	   6
//            /
//          7
func BTreeInit() *BTree {
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 3}
	node4 := &Node{Data: 4}
	node5 := &Node{Data: 5}
	node6 := &Node{Data: 6}
	node7 := &Node{Data: 7}

	node1.Left = node2
	node1.Right = node3
	node2.Parent = node1
	node3.Parent = node1

	node2.Left = node4
	node2.Right = node5
	node4.Parent = node2
	node5.Parent = node2

	node3.Right = node6
	node6.Parent = node3

	node6.Left = node7
	node7.Parent = node6

	btree := &BTree {node1, 7}
	return btree
}

// BTreeTest ...
func BTreeTest() {
	fmt.Println("BTreeTest begin")
	btree := BTreeInit()
	btree.PreorderTraverse()
	btree.InorderTraverse()
	btree.PostorderTraverse()
	btree.LevelTraverse()
	fmt.Printf("Height: %d\n", btree.GetHeight())
	
	target := 1
	parent := btree.GetParent(target)
	if parent != nil {
		fmt.Printf("parent of %v is: %v\n", target, parent.Data)
	} else {
		fmt.Printf("parent of %v is: %v\n", target, parent)
	}

	target = 6
	parent = btree.GetParent2(target)
	if parent != nil {
		fmt.Printf("parent of %v is: %v\n", target, parent.Data)
	} else {
		fmt.Printf("parent of %v is: %v\n", target, parent)
	}

	btree.PreorderWithoutRecursion()
	btree.InorderWithoutRecursion()
	btree.PostorderWithoutRecursion()

	fmt.Println("BTreeTest end")
}