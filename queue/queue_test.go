package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewEmpty[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	one, err := q.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if one != 1 {
		t.Fatalf("Mismatch: got %v, wanted %v", one, 1)
	}
	q.Dequeue()
	q.Dequeue()
	_, err = q.Dequeue()
	if err == nil {
		t.Fatalf("Empty Queue dequeue 1")
	}
	_, err = q.Dequeue()
	if err == nil {
		t.Fatalf("Empty Queue dequeue 2")
	}
	size := q.Size()
	if size != 0 {
		t.Fatal("Nonzero size")
	}
}
