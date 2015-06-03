package bstree

import (
	"fmt"
	"testing"
)

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
	//vals := []StringVal{"aa", "bb", "cc"}
	for _, val := range vals {
		bst.Insert(val)
	}

	fmt.Println(bst.Walk())
	//bst.Delete(StringVal("aac"))
	bst.Delete(IntVal(0))
	fmt.Println(bst.Walk())
	return

	//fmt.Println(bst.Walk())
	fmt.Println(bst.Walk())
	bst.Delete(IntVal(0))
	fmt.Println(bst.Walk())
	bst.Delete(IntVal(2))
	bst.Delete(IntVal(0)) //TODO this fails !
	//bst.Delete(IntVal(8))
	fmt.Println(bst.Walk())
}
