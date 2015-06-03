package bstree

import "fmt"

type Val interface {
	LessThan(Val) bool
	GreaterThan(Val) bool
	Equals(Val) bool
}
type StringVal string
type IntVal int

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
	Walk() []Val
	Insert(Val)
	Height() int
	Size() int
	Delete(Val) error
	isBalanced() bool
}

type BSTree struct {
	root *Node
	sz   int
}

type Node struct {
	left  *Node
	right *Node
	val   Val
}

func (tree *BSTree) Size() int {
	return tree.sz
}

func (tree *BSTree) Walk() []Val {
	var vals = make([]Val, tree.sz, tree.sz)

	if tree.sz > 0 {
		walk(tree.root, vals, 0)
	}

	return vals
}

func (tree *BSTree) Insert(val Val) {
	tree.root = addNode(tree.root, val)
	tree.sz++
	return
}

func (bst *BSTree) Delete(val Val) error {
	var err error

	if _, err = delNode(bst.root, val); err == nil {
		bst.sz--
	}

	return err
}

func (bst *BSTree) Height() int {
	return height(bst.root)
}

func (bst *BSTree) isBalanced() bool {
	lht := height(bst.root.left)
	rht := height(bst.root.right)

	if lht > rht+1 || rht > lht+1 {
		return false
	}
	return true
}

func walk(node *Node, vals []Val, idx int) int {
	if node == nil {
		return idx
	}

	idx = walk(node.left, vals, idx)
	vals[idx] = node.val
	idx = walk(node.right, vals, idx+1)

	return idx
}

func height(node *Node) int {
	if node == nil {
		return -1
	}

	ht := 0

	hLst := height(node.left)
	hRst := height(node.right)

	if hLst > hRst {
		ht = hLst + 1
	} else {
		ht = hRst + 1
	}

	return ht
}

func nextInOrder(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func addNode(root *Node, val Val) *Node {
	if root == nil {
		return &Node{nil, nil, val}
	}
	if val.LessThan(root.val) {
		root.left = addNode(root.left, val)
	} else {
		root.right = addNode(root.right, val)
	}

	return root
}

func delNode(root *Node, val Val) (*Node, error) {
	err := fmt.Errorf("Value %v not found", val)
	if root == nil {
		return root, err
	}

	if val.LessThan(root.val) {
		root.left, err = delNode(root.left, val)
	} else if val.GreaterThan(root.val) {
		root.right, err = delNode(root.right, val)
	} else {
		if root.right == nil {
			return root.left, nil
		} else if root.left == nil {
			return root.right, nil
		} else {
			nnode := nextInOrder(root.right)
			root.val = nnode.val
			root.right, err = delNode(root.right, nnode.val)
		}
	}
	return root, err
}
