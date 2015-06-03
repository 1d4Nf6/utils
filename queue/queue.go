package queue

import "errors"

type Value interface{}

type Queue interface {
	Push(v Value)
	PushAll([]Value)
	Pop() (interface{}, error)
	PopAll() []Value
	Len() int
}

type SQueue struct {
	head *Node
	tail *Node
	ln   int
}

type Node struct {
	next *Node
	prev *Node
	val  interface{}
}

func (pq *SQueue) Push(val Value) {
	node := &Node{nil, nil, val}

	if pq.head == nil {
		pq.head = node
		pq.tail = node
	} else {
		node.next = pq.head
		pq.head.prev = node
		pq.head = node
	}

	pq.ln++
}

func (pq *SQueue) Pop() (interface{}, error) {
	if pq.tail == nil {
		return -1, errors.New("SQueue empty")
	}
	ret := pq.tail.val
	pq.tail = pq.tail.prev

	// If we are empty now, clear the head as well
	if pq.tail == nil {
		pq.head = nil
	} else {
		pq.tail.next = nil
	}

	pq.ln--

	return ret, nil
}

func (pq *SQueue) PopAll() []Value {
	vals := []Value{}
	for v, e := pq.Pop(); e == nil; v, e = pq.Pop() {
		vals = append(vals, v)
	}

	return vals
}

func (pq *SQueue) PushAll(vals []Value) {
	for _, v := range vals {
		switch v.(type) {
		case int:
			pq.Push(v.(int))
		case string:
			pq.Push(v.(string))
		}
	}
}

func (pq *SQueue) Len() int {
	return pq.ln
}

func NewSQueue() *SQueue {
	return &SQueue{nil, nil, 0}
}
