package rbtree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	//var rbt Tree = new(RBTree)
	var rbt Tree = NewRBTree()
	vals := []IntVal{5, 1, 8, 2, 9, 0, 7}
	for _, val := range vals {
		rbt.Insert(val)
	}

	if l := rbt.Size(); l != 7 {
		t.Error("Expected 7, received", l)
	}

	if l := rbt.Walk(); l[0] != IntVal(0) || l[6] != IntVal(9) {
		t.Error("Expected [0, 9], received", l)
	}

	fmt.Println(rbt.Walk(), rbt.isBalanced())
}

func TestBalanced(t *testing.T) {
	var rbt Tree = NewRBTree()
	vals := []IntVal{1, 2, 3}
	for _, val := range vals {
		rbt.Insert(val)
	}

	//if !rbt.isBalanced() {
	//t.Error("Tree not balanced")
	//}
}

func TestDelete(t *testing.T) {
	//var rbt Tree = new(RBTree)
	var rbt Tree = NewRBTree()
	vals := []IntVal{5, 1, 8, 2, 9, 0, 7}
	for _, val := range vals {
		rbt.Insert(val)
	}

	rbt.Delete(IntVal(0))
	if l := rbt.Size(); l != 6 {
		t.Error("Expected 6, received", l)
	}

	rbt.Delete(IntVal(2))
	if l := rbt.Size(); l != 5 {
		t.Error("Expected 5, received", l)
	}

	rbt.Delete(IntVal(5))
	if l := rbt.Size(); l != 4 {
		t.Error("Expected 4, received", l)
	}

	if b := rbt.Walk(); !b[0].Equals(IntVal(1)) {
		t.Errorf("Expected 1, received %v", b[0])
	}
}
