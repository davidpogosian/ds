package priority_queue

import (
	"testing"

	"github.com/davidpogosian/ds/comparators"
	"github.com/davidpogosian/ds/testutils"
)

func TestNewEmpty(t *testing.T) {
	pq := NewEmpty[int, string](comparators.ComparatorInt, false)
	testutils.Assert(t, "pq.Size()", 0, pq.Size())
}

func TestEnqueue(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		pq.Enqueue(1, "low priority")
		pq.Enqueue(2, "medium priority")
		pq.Enqueue(3, "high priority")
		testutils.Assert(t, "pq.Size()", 3, pq.Size())
	})

	t.Run("Concurrent", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			pq.Enqueue(1, "x")
			return nil
		})
		testutils.Assert(t, "pq.Size()", 1000, pq.Size())
	})
}

func TestPeek(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		_, _, err := pq.Peek()
		if err == nil {
			t.Fatal("Performed peek on an empty PriorityQueue")
		}
	})

	t.Run("NotEmpty", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		pq.Enqueue(1, "low")
		pq.Enqueue(2, "medium")
		pq.Enqueue(3, "high")
		three, high, err := pq.Peek()
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "three", 3, three)
		testutils.Assert(t, "high", "high", high)
	})
}

func TestExtractTop(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		t.Run("Max", func(t *testing.T) {
			pq := NewEmpty[int, string](comparators.ComparatorInt, false)
			pq.Enqueue(1, "low priority")
			pq.Enqueue(2, "medium priority")
			pq.Enqueue(3, "high priority")
			three, highpriority, err := pq.ExtractTop()
			if err != nil {
				t.Fatal(err)
			}
			testutils.Assert(t, "pq.Size()", 2, pq.Size())
			testutils.Assert(t, "three", 3, three)
			testutils.Assert(t, "highpriority", "high priority", highpriority)
		})

		t.Run("Min", func(t *testing.T) {
			pq := NewEmpty[int, string](comparators.ComparatorInt, true)
			pq.Enqueue(1, "low priority")
			pq.Enqueue(2, "medium priority")
			pq.Enqueue(3, "high priority")
			one, lowpriority, err := pq.ExtractTop()
			if err != nil {
				t.Fatal(err)
			}
			testutils.Assert(t, "pq.Size()", 2, pq.Size())
			testutils.Assert(t, "one", 1, one)
			testutils.Assert(t, "lowpriority", "low priority", lowpriority)
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, true)
		for i := 0; i < 1000; i++ {
			pq.Enqueue(1, "Hi")
		}
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			_, _, err := pq.ExtractTop()
			return err
		})
		testutils.Assert(t, "pq.Size()", 0, pq.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		testutils.Assert(t, "pq.IsEmpty()", true, pq.IsEmpty())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		pq.Enqueue(1, "low")
		pq.Enqueue(2, "medium")
		pq.Enqueue(3, "high")
		testutils.Assert(t, "pq.IsEmpty()", false, pq.IsEmpty())
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		testutils.Assert(t, "pq.Size()", 0, pq.Size())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		pq := NewEmpty[int, string](comparators.ComparatorInt, false)
		pq.Enqueue(1, "low")
		pq.Enqueue(2, "medium")
		pq.Enqueue(3, "high")
		testutils.Assert(t, "pq.Size()", 3, pq.Size())
	})
}

func TestClear(t *testing.T) {
	pq := NewEmpty[int, string](comparators.ComparatorInt, false)
	pq.Enqueue(1, "low")
	pq.Enqueue(2, "medium")
	pq.Enqueue(3, "high")
	pq.Clear()
	testutils.Assert(t, "pq.Size()", 0, pq.Size())
}

func TestCopy(t *testing.T) {
	pq1 := NewEmpty[int, string](comparators.ComparatorInt, true)
	pq1.Enqueue(2, "Stwing")
	pq2 := pq1.Copy()
	testutils.Assert(t, "pq2.Size()", 1, pq2.Size())
	pq1.Enqueue(3, "Awso Stwing")
	testutils.Assert(t, "pq2.Size()", 1, pq2.Size())
}
