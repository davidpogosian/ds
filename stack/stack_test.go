package stack

import (
	"sync"
	"testing"

	"github.com/davidpogosian/ds/comparators"
	"github.com/davidpogosian/ds/testutils"
)

func TestNewEmpty(t *testing.T) {
	s := NewEmpty[int](comparators.ComparatorInt)
	testutils.Assert(t, "s.Size()", 0, s.Size())
	testutils.Assert(t, "s.String()", "[]", s.String())
}

func TestNewFromSlice(t *testing.T) {
	t.Run("InitializedSlice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		s := NewFromSlice(slice, comparators.ComparatorInt)
		err := testutils.CompareSlices(slice, s.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("NilSlice", func(t *testing.T) {
		var slice []float64
		s := NewFromSlice(slice, comparators.ComparatorFloat64)
		testutils.Assert(t, "s.Size()", 0, s.Size())
		testutils.Assert(t, "s.String()", "[]", s.String())
	})

	t.Run("ModifySlice", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		slice := []int{1, 2, 3}
		s := NewFromSlice(slice, comparators.ComparatorInt)
		slice[2] = 99
		err := testutils.CompareSlices(originalSlice, s.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			s := NewEmpty[int](comparators.ComparatorInt)
			_, err := s.Pop()
			if err == nil {
				t.Fatal("Popped from empty stack")
			}
		})

		t.Run("NotEmpty", func(t *testing.T) {
			s := NewEmpty[int](comparators.ComparatorInt)
			s.Push(1)
			one, err := s.Pop()
			if err != nil {
				t.Fatal(err)
			}
			testutils.Assert(t, "one", 1, one)
			testutils.Assert(t, "s.Size()", 0, s.Size())
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		s := NewEmpty[int](comparators.ComparatorInt)
		for i := 0; i < 1000; i++ {
			s.Push(i)
		}
		testutils.Assert(t, "s.Size()", 1000, s.Size())
		threads := 10
		operations := 100
		errorChannel := make(chan error, 1000)
		var waitGroup sync.WaitGroup
		for i := 0; i < threads; i++ {
			waitGroup.Add(1)
			go func(i int) {
				defer waitGroup.Done()
				for j := 0; j < operations; j++ {
					_, err := s.Pop()
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
 		testutils.Assert(t, "s.Size()", 0, s.Size())
	})
}

func TestPush(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			s := NewEmpty[int](comparators.ComparatorInt)
			s.Push(1)
			err := testutils.CompareSlices([]int{1}, s.ToSlice())
			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("NotEmpty", func(t *testing.T) {
			s := NewFromSlice([]int{1, 2}, comparators.ComparatorInt)
			s.Push(3)
			err := testutils.CompareSlices([]int{1, 2, 3}, s.ToSlice())
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		s := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "s.Size()", 0, s.Size())
		threads := 10
		operations := 100
		var waitGroup sync.WaitGroup
		for i := 0; i < threads; i++ {
			waitGroup.Add(1)
			go func(i int) {
				defer waitGroup.Done()
				for j := 0; j < operations; j++ {
					s.Push(i)
				}
			}(i)
		}
		waitGroup.Wait()
 		testutils.Assert(t, "s.Size()", 1000, s.Size())
	})
}

func TestPeek(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int](comparators.ComparatorInt)
		_, err := s.Peek()
		if err == nil {
			t.Fatal("Peeked empty stack")
		}
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2}, comparators.ComparatorInt)
		two, err := s.Peek()
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "two", 2, two)
		testutils.Assert(t, "s.Size()", 2, s.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "s.IsEmpty()", true, s.IsEmpty())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		testutils.Assert(t, "s.IsEmpty()", false, s.IsEmpty())
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "s.Size()", 0, s.Size())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		testutils.Assert(t, "s.Size()", 3, s.Size())
	})
}

func TestClear(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int](comparators.ComparatorInt)
		s.Clear()
		testutils.Assert(t, "s.Size()", 0, s.Size())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		s.Clear()
		testutils.Assert(t, "s.Size()", 0, s.Size())
	})
}

func TestFind(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		one := s.Find(2)
		testutils.Assert(t, "one", 1, one)
	})

	t.Run("DoesntExist", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		negativeOne := s.Find(1099)
		testutils.Assert(t, "negativeOne", -1, negativeOne)
	})
}

func TestToSlice(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		s := NewFromSlice(originalSlice, comparators.ComparatorInt)
		slice := s.ToSlice()
		err := testutils.CompareSlices(originalSlice, slice)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ModifyStack", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		s := NewFromSlice(originalSlice, comparators.ComparatorInt)
		slice := s.ToSlice()
		slice[2] = 99
		three, err := s.Pop()
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "three", 3, three)
	})
}

func TestCopy(t *testing.T) {
	s1 := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
	s2 := s1.Copy()
	s1.Pop()
	testutils.Assert(t, "s2.Size()", 3, s2.Size())
}

func TestString(t *testing.T) {
	s := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
	testutils.Assert(t, "s.String()", "[1 2 3]", s.String())
}
