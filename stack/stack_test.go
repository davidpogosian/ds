package stack

import (
	"fmt"
	"testing"
)

func assert[T comparable](t *testing.T, varname string, expected T, got T) {
	if expected != got {
		t.Fatalf("Expected '%s' to be: %v, instead got: %v", varname, expected, got)
	}
}

func compareSlices[T comparable](sliceA []T, sliceB []T) error {
	if (len(sliceA) != len(sliceB)) {
		return fmt.Errorf("SliceA and SliceB are of different lengths.")
	}
	for i := range sliceA {
		if (sliceA[i] != sliceB[i]) {
			return fmt.Errorf("Items at index %d do not match.", i)
		}
	}
	return nil
}

func TestNewEmpty(t *testing.T) {
	s := NewEmpty[int]()
	assert(t, "s.Size()", 0, s.Size())
	assert(t, "s.String()", "[]", s.String())
}

func TestNewFromSlice(t *testing.T) {
	t.Run("InitializedSlice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		s := NewFromSlice(slice)
		err := compareSlices(slice, s.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("NilSlice", func(t *testing.T) {
		var slice []float64
		s := NewFromSlice(slice)
		assert(t, "s.Size()", 0, s.Size())
		assert(t, "s.String()", "[]", s.String())
	})

	t.Run("ModifySlice", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		slice := []int{1, 2, 3}
		s := NewFromSlice(slice)
		slice[2] = 99
		err := compareSlices(originalSlice, s.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("EmptyStack", func(t *testing.T) {
		s := NewEmpty[int]()
		_, err := s.Pop()
		if err == nil {
			t.Fatal("Popped from empty stack")
		}
	})

	t.Run("NonemptyStack", func(t *testing.T) {
		s := NewEmpty[int]()
		s.Push(1)
		top, err := s.Pop()
		if err != nil {
			t.Fatal(err)
		}
		assert(t, "top", 1, top)
		assert(t, "s.Size()", 0, s.Size())
	})
}

func TestPush(t *testing.T) {
	t.Run("EmptyStack", func(t *testing.T) {
		s := NewEmpty[int]()
		s.Push(1)
		err := compareSlices([]int{1}, s.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("NonemptyStack", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2})
		s.Push(3)
		err := compareSlices([]int{1, 2, 3}, s.ToSlice())
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestPeek(t *testing.T) {
	t.Run("EmptyStack", func(t *testing.T) {
		s := NewEmpty[int]()
		_, err := s.Peek()
		if err == nil {
			t.Fatal("Peeked empty stack")
		}
	})

	t.Run("NonemptyStack", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2})
		two, err := s.Peek()
		if err != nil {
			t.Fatal(err)
		}
		assert(t, "two", 2, two)
		assert(t, "s.Size()", 2, s.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("EmptyStack", func(t *testing.T) {
		s := NewEmpty[int]()
		assert(t, "s.IsEmpty()", true, s.IsEmpty())
	})

	t.Run("NonemptyStack", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		assert(t, "s.IsEmpty()", false, s.IsEmpty())
	})
}

func TestSize(t *testing.T) {
	t.Run("EmptyStack", func(t *testing.T) {
		s := NewEmpty[int]()
		assert(t, "s.Size()", 0, s.Size())
	})

	t.Run("NonemptyStack", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		assert(t, "s.Size()", 3, s.Size())
	})
}

func TestClear(t *testing.T) {
	t.Run("EmptyStack", func(t *testing.T) {
		s := NewEmpty[int]()
		s.Clear()
		assert(t, "s.Size()", 0, s.Size())
	})

	t.Run("NonemptyStack", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		s.Clear()
		assert(t, "s.Size()", 0, s.Size())
	})
}

func TestContains(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		one := s.Contains(2)
		assert(t, "one", 1, one)
	})

	t.Run("DoesntExist", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		negativeOne := s.Contains(1099)
		assert(t, "negativeOne", -1, negativeOne)
	})
}

func TestToSlice(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		s := NewFromSlice(originalSlice)
		slice := s.ToSlice()
		err := compareSlices(originalSlice, slice)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ModifyStack", func(t *testing.T) {
		originalSlice := []int{1, 2, 3}
		s := NewFromSlice(originalSlice)
		slice := s.ToSlice()
		slice[2] = 99
		three, err := s.Pop()
		if err != nil {
			t.Fatal(err)
		}
		assert(t, "three", 3, three)
	})
}

func TestCopy(t *testing.T) {
	s1 := NewFromSlice([]int{1, 2, 3})
	s2 := s1.Copy()
	s1.Pop()
	assert(t, "s2.Size()", 3, s2.Size())
}

func TestString(t *testing.T) {
	s1 := NewFromSlice([]int{1, 2, 3})
	assert(t, "s1.String()", "[1 2 3]", s1.String())
}
