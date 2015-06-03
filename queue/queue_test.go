package queue

import "testing"

func TestSqueue(t *testing.T) {
	var pq Queue = NewSQueue()
	if l := pq.Len(); l != 0 {
		t.Error("Expected 0, received", l)
	}

	pq.Push(1)
	pq.Push("abc")
	if l := pq.Len(); l != 2 {
		t.Error("Expected 2, received", l)
	}

	if v, _ := pq.Pop(); v != 1 {
		t.Error("Expected 1, received", v)
	}
	if v, _ := pq.Pop(); v != "abc" {
		t.Error("Expected abc, received", v)
	}
	if l := pq.Len(); l != 0 {
		t.Error("Expected 0, received", l)
	}

	if _, e := pq.Pop(); e == nil {
		t.Error("Expected error QueueEmpty")
	}

	pq.Push(2)
	if l := pq.Len(); l != 1 {
		t.Error("Expected 1, received", l)
	}
	if v, _ := pq.Pop(); v != 2 {
		t.Error("Expected 2, received", v)
	}
}

func TestSqueueAll(t *testing.T) {
	var pq Queue = NewSQueue()
	pq.PushAll([]Value{1, 2, 3, 4})
	if l := pq.Len(); l != 4 {
		t.Error("Expected 4, received", l)
	}
	pq.PushAll([]Value{"abc", "def", "ghi"})
	if l := pq.Len(); l != 7 {
		t.Error("Expected 7, received", l)
	}

	if v := pq.PopAll(); v[0] != 1 {
		t.Error("Expected 1, received", v[0])
	}

	if l := pq.Len(); l != 0 {
		t.Error("Expected 0, received", l)
	}

	pq.PushAll([]Value{5, "klm", 6, "nop"})
	if l := pq.Len(); l != 4 {
		t.Error("Expected 4, received", l)
	}

	if v := pq.PopAll(); v[0] != 5 || v[1] != "klm" {
		t.Error("Expected 5 and klm, received", v[0], v[1])
	}
}
