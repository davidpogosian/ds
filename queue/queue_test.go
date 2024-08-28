package queue

import (
	"sync"
	"testing"

	"github.com/davidpogosian/ds/testutils"
)

func TestNewEmpty(t *testing.T) {
	q := NewEmpty[int]()
	testutils.Assert(t, "q.Size()", 0, q.Size())
	testutils.Assert(t, "q.String()", "[]", q.String())
}

func TestNewFromSlice(t *testing.T) {
	t.Run("InitializedSlice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		q := NewFromSlice(slice)
		err := testutils.CompareSlices(slice, q.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("NilSlice", func(t *testing.T) {
		var slice []float64
		q := NewFromSlice(slice)
		testutils.Assert(t, "q.Size()", 0, q.Size())
		testutils.Assert(t, "q.String()", "[]", q.String())
	})

	t.Run("ModifySlice", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		slice := []int{1, 2, 3}
		q := NewFromSlice(slice)
		slice[2] = 99
		err := testutils.CompareSlices(originalSlice, q.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestEnqueue(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		q := NewEmpty[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		testutils.Assert(t, "q.Size()", 3, q.Size())
	})

	t.Run("Concurrent", func(t *testing.T) {
		q := NewEmpty[int]()
		testutils.Assert(t, "q.Size()", 0, q.Size())
		threads := 10
		operations := 100
		var waitGroup sync.WaitGroup
		for i := 0; i < threads; i++ {
			waitGroup.Add(1)
			go func(i int) {
				defer waitGroup.Done()
				for j := 0; j < operations; j++ {
					q.Enqueue(i)
				}
			}(i)
		}
		waitGroup.Wait()
 		testutils.Assert(t, "q.Size()", 1000, q.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		q := NewEmpty[int]()
		testutils.Assert(t, "q.IsEmpty()", true, q.IsEmpty())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		q := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "q.IsEmpty()", false, q.IsEmpty())
	})
}

func TestDequeue(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		q := NewFromSlice([]int{1, 2, 3})
		one, err := q.Dequeue()
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "one", 1, one)
	})

	t.Run("Concurrent", func(t *testing.T) {
		q := NewEmpty[int]()
		for i := 0; i < 1000; i++ {
			q.Enqueue(i)
		}
		testutils.Assert(t, "q.Size()", 1000, q.Size())
		threads := 10
		operations := 100
		errorChannel := make(chan error, 1000)
		var waitGroup sync.WaitGroup
		for i := 0; i < threads; i++ {
			waitGroup.Add(1)
			go func(i int) {
				defer waitGroup.Done()
				for j := 0; j < operations; j++ {
					_, err := q.Dequeue()
					if err != nil {
						errorChannel <- err
					}
				}
			}(i)
		}
		waitGroup.Wait()
		close(errorChannel)
		for err := range errorChannel {
			t.Fatal(err)
		}
 		testutils.Assert(t, "q.Size()", 0, q.Size())
	})
}

func TestPeek(t *testing.T) {
	q := NewFromSlice([]int{1, 2, 3})
	one, err := q.Peek()
	if err != nil {
		t.Fatal(err)
	}
	testutils.Assert(t, "one", 1, one)
	testutils.Assert(t, "q.Size()", 3, q.Size())
}

func TestSize(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		q := NewEmpty[int]()
		testutils.Assert(t, "q.Size()", 0, q.Size())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		q := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "q.Size()", 3, q.Size())
	})
}

func TestClear(t *testing.T) {
	q := NewFromSlice([]int{1, 2, 3})
	q.Clear()
	testutils.Assert(t, "q.Size()", 0, q.Size())
}

func TestContains(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		q := NewFromSlice([]int{1, 2, 3})
		one := q.Contains(2)
		testutils.Assert(t, "one", 1, one)
	})

	t.Run("DoesntExist", func(t *testing.T) {
		q := NewFromSlice([]int{1, 2, 3})
		negativeOne := q.Contains(1099)
		testutils.Assert(t, "negativeOne", -1, negativeOne)
	})
}

func TestCopy(t *testing.T) {
	q1 := NewFromSlice([]int{1, 2, 3})
	q2 := q1.Copy()
	q1.Dequeue()
	testutils.Assert(t, "q2.Size()", 3, q2.Size())
}

func TestToSlice(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		q := NewFromSlice(originalSlice)
		slice := q.ToSlice()
		err := testutils.CompareSlices(originalSlice, slice)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ModifyQueue", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		q := NewFromSlice(originalSlice)
		slice := q.ToSlice()
		slice[0] = 99
		one, err := q.Dequeue()
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "one", 1, one)
	})
}

func TestString(t *testing.T) {
	q := NewFromSlice([]int{1, 2, 3})
	testutils.Assert(t, "q.String()", "[1 2 3]", q.String())
}
