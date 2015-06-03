package bstree

import "testing"

func TestInsert(t *testing.T) {
	var bst Tree = new(BSTree)
	vals := []IntVal{5, 1, 8, 2, 9, 0, 7}
	for _, val := range vals {
		bst.Insert(val)
	}

	if l := bst.Size(); l != 7 {
		t.Error("Expected 7, received", l)
	}

	if l := bst.Walk(); l[0] != IntVal(0) || l[6] != IntVal(9) {
		t.Error("Expected [0, 9], received", l)
	}
}

func TestDelete(t *testing.T) {
	var bst Tree = new(BSTree)
	vals := []IntVal{5, 1, 8, 2, 9, 0, 7}
	for _, val := range vals {
		bst.Insert(val)
	}

	bst.Delete(IntVal(0))
	if l := bst.Size(); l != 6 {
		t.Error("Expected 6, received", l)
	}

	bst.Delete(IntVal(2))
	if l := bst.Size(); l != 5 {
		t.Error("Expected 5, received", l)
	}

	bst.Delete(IntVal(5))
	if l := bst.Size(); l != 4 {
		t.Error("Expected 4, received", l)
	}

	if b := bst.Walk(); !b[0].Equals(IntVal(1)) {
		t.Errorf("Expected 1, received %v", b[0])
	}

}
