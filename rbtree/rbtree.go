/* Package rbtree implements a Red-Black Tree.
The algorithm has been adopted from CLR.
*/

package rbtree

import "fmt"

type Val interface {
	LessThan(Val) bool
	GreaterThan(Val) bool
	Equals(Val) bool
}
type StringVal string
type IntVal int

type NodeVals []Val

func (a StringVal) LessThan(b Val) bool {
	return a < b.(StringVal)
}

func (a StringVal) GreaterThan(b Val) bool {
	return a > b.(StringVal)
}

func (a StringVal) Equals(b Val) bool {
	return a == b.(StringVal)
}

func (a IntVal) LessThan(b Val) bool {
	return a < b.(IntVal)
}

func (a IntVal) GreaterThan(b Val) bool {
	return a > b.(IntVal)
}

func (a IntVal) Equals(b Val) bool {
	return a == b.(IntVal)
}

type Tree interface {
	Walk() NodeVals
	Insert(Val)
	Height() int
	Size() int
	Delete(Val) error
	isBalanced() bool
}

type RBTree struct {
	root     *Node
	Sentinel *Node
	sz       int
}

type Color uint8

const (
	RED   Color = 0
	BLACK Color = 1
)

type Node struct {
	left  *Node
	right *Node
	p     *Node
	color Color
	val   Val
}

func (tree *RBTree) Size() int {
	return tree.sz
}

func (tree *RBTree) Walk() NodeVals {
	var vals = make([]Val, tree.sz, tree.sz)

	if tree.sz > 0 {
		tree.walk(tree.root, vals, 0)
	}

	return vals
}

func NewRBTree() *RBTree {
	sentinel := &Node{nil, nil, nil, BLACK, nil}
	return &RBTree{sentinel, sentinel, 0}
}

func (tree *RBTree) Insert(val Val) {
	tree.insertNode(val)
}

func (tree *RBTree) insertNode(val Val) {
	z := &Node{tree.Sentinel, tree.Sentinel, tree.Sentinel, RED, val}
	y := tree.Sentinel
	x := tree.root

	for x != tree.Sentinel {
		y = x
		if val.LessThan(x.val) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.p = y

	if y == tree.Sentinel {
		tree.root = z
	} else if z.val.LessThan(y.val) {
		y.left = z
	} else {
		y.right = z
	}

	tree.sz++

	tree.fixUp(z)
	return
}

func (tree *RBTree) Delete(val Val) error {
	var err error

	if _, err = tree.delNode(tree.root, val); err == nil {
		tree.sz--
	}

	return err
}

func (tree *RBTree) height(node *Node) int {
	if node == tree.Sentinel {
		return -1
	}

	ht := 0

	hLst := tree.height(node.left)
	hRst := tree.height(node.right)

	if hLst > hRst {
		ht = hLst + 1
	} else {
		ht = hRst + 1
	}

	return ht
}

func (tree *RBTree) Height() int {
	return tree.height(tree.root)
}

func (tree *RBTree) fixUp(z *Node) {
	for z.p.color == RED {
		if z.p == z.p.p.left {
			y := z.p.p.right
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.right {
					z = z.p
					tree.rotateLeft(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				tree.rotateRight(z.p.p)
			}
		} else {
			y := z.p.p.left
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					tree.rotateRight(z.p.p)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				tree.rotateLeft(z)
			}
		}
	}

	tree.root.color = BLACK
}

func (tree *RBTree) rotateLeft(x *Node) {
	y := x.right
	x.right = y.left

	if y.left != tree.Sentinel {
		y.left.p = x
	}

	y.p = x.p

	if x.p == tree.Sentinel {
		tree.root = y
	} else if x == x.p.left {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.left = x
	x.p = y
}

func (tree *RBTree) rotateRight(x *Node) {
	y := x.left
	x.left = y.right

	if y.right != tree.Sentinel {
		y.right.p = x
	}

	y.p = x.p

	if x.p == tree.Sentinel {
		tree.root = y
	} else if x == x.p.right {
		x.p.right = y
	} else {
		x.p.left = y
	}
	y.right = x
	x.p = y
}

func (tree *RBTree) isBalanced() bool {
	lht := tree.height(tree.root.left)
	rht := tree.height(tree.root.right)

	if lht > rht+1 || rht > lht+1 {
		return false
	}
	return true
}

func (tree *RBTree) walk(node *Node, vals []Val, idx int) int {
	if node == tree.Sentinel {
		return idx
	}

	idx = tree.walk(node.left, vals, idx)
	vals[idx] = node.val
	idx = tree.walk(node.right, vals, idx+1)

	return idx
}

func (tree *RBTree) nextInOrder(node *Node) *Node {
	for node.left != tree.Sentinel {
		node = node.left
	}
	return node
}

func (tree *RBTree) delNode(root *Node, val Val) (*Node, error) {
	err := fmt.Errorf("Value %v not found", val)
	if root == tree.Sentinel {
		return root, err
	}

	if val.LessThan(root.val) {
		root.left, err = tree.delNode(root.left, val)
	} else if val.GreaterThan(root.val) {
		root.right, err = tree.delNode(root.right, val)
	} else {
		if root.right == tree.Sentinel {
			return root.left, nil
		} else if root.left == tree.Sentinel {
			return root.right, nil
		} else {
			nnode := tree.nextInOrder(root.right)
			root.val = nnode.val
			root.right, err = tree.delNode(root.right, nnode.val)
		}
	}
	return root, err
}
