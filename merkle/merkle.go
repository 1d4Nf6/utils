package merkle

import (
	"hash"

	rbt "github.com/1d4Nf6/utils/rbtree"
)

type NodeVal struct {
	val int
	h   hash.Hash
}

type MerkleTree rbt.RBTree

func (v1 NodeVal) LessThan(v2 rbt.Val) bool {
	return v1.val < v2.(NodeVal).val
}

func (v1 NodeVal) GreaterThan(v2 rbt.Val) bool {
	return v1.val > v2.(NodeVal).val
}

func (v1 NodeVal) Equals(v2 rbt.Val) bool {
	return v1.val == v2.(NodeVal).val
}

func NewMerkleTree() *MerkleTree {
	return (*MerkleTree)(rbt.NewRBTree())
}

func (tree *MerkleTree) Walk() rbt.NodeVals {
	return (*rbt.RBTree)(tree).Walk()
}

func (tree *MerkleTree) Insert(mn NodeVal) {
	(*rbt.RBTree)(tree).Insert(mn)
}

func (tree *MerkleTree) Size() int {
	return (*rbt.RBTree)(tree).Size()
}
