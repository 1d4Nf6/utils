package merkle

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	mt := NewMerkleTree()
	vals := []NodeVal{{val: 5}, {val: 1}, {val: 8}, {val: 2}, {val: 9}, {val: 0}, {val: 7}}
	for _, val := range vals {
		mt.Insert(val)
	}

	if l := mt.Size(); l != 7 {
		t.Error("Expected 7, received", l)
	}

	if l := mt.Walk(); l[0].(NodeVal).val != 0 || l[6].(NodeVal).val != 9 {
		t.Error("Expected [0, 9], received", l)
	}

	//fmt.Println(mt.Walk(), mt.isBalanced())

	fmt.Println(mt.Walk())
}
