package set

import (
	"testing"

	"github.com/davidpogosian/ds/testutils"
)

func TestNewEmpty(t *testing.T) {
	s := NewEmpty[int]()
	testutils.Assert(t, "s.Size()", 0, s.Size())
}

func TestNewFromSlice(t *testing.T) {
	t.Run("UninitializedSlice", func(t *testing.T) {
		var slice []int
		s := NewFromSlice(slice)
		testutils.Assert(t, "s.Size()", 0, s.Size())
		// Ensure Set still works.
		s.Add(1)
		testutils.Assert(t, "s.Size()", 1, s.Size())
		testutils.Assert(t, "s.Contains(1)", true, s.Contains(1))
	})

	t.Run("InitializedSlice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		s := NewFromSlice(slice)
		testutils.Assert(t, "s.Size()", 3, s.Size())
		testutils.Assert(t, "s.Contains(1)", true, s.Contains(1))
		testutils.Assert(t, "s.Contains(2)", true, s.Contains(2))
		testutils.Assert(t, "s.Contains(3)", true, s.Contains(3))
	})
}

func TestString(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int]()
		testutils.Assert(t, "s.String()", "[]", s.String())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "s.String()", "[1 2 3]", s.String())
	})
}

func TestCopy(t *testing.T) {
	s1 := NewFromSlice([]int{1, 2, 3})
	s2 := s1.Copy()
	s2.Add(4)
	testutils.Assert(t, "s1.Size()", 3, s1.Size())
	testutils.Assert(t, "s2.Size()", 4, s2.Size())
}

func TestContains(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "s.Contains(2)", true, s.Contains(2))
	})

	t.Run("NotExists", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "s.Contains(4)", false, s.Contains(4))
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int]()
		testutils.Assert(t, "s.Size()", 0, s.Size())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "s.Size()", 3, s.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int]()
		testutils.Assert(t, "s.IsEmpty()", true, s.IsEmpty())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "s.IsEmpty()", false, s.IsEmpty())
	})
}

func TestClear(t *testing.T) {
	s := NewFromSlice([]int{1, 2, 3})
	s.Clear()
	testutils.Assert(t, "s.Size()", 0, s.Size())
}

func TestToSlice(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := NewEmpty[int]()
		testutils.Assert(t, "len(s.ToSlice())", 0, len(s.ToSlice()))
	})

	t.Run("NotEmpty", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		testutils.Assert(t, "len(s.ToSlice())", 3, len(s.ToSlice()))
	})
}

func TestAdd(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		s := NewEmpty[int]()
		s.Add(1)
		testutils.Assert(t, "s.Contains(1)", true, s.Contains(1))
	})
}

func TestRemove(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		s.Remove(1)
		testutils.Assert(t, "s.Contains(1)", false, s.Contains(1))
	})
}

func TestUnion(t *testing.T) {
	t.Run("NoIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{4, 5, 6})
		union := s1.Union(s2)
		testutils.Assert(t, "union.Size()", 6, union.Size())
	})

	t.Run("PartialIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{3, 4, 5})
		union := s1.Union(s2)
		testutils.Assert(t, "union.Size()", 5, union.Size())
	})

	t.Run("FullIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{1, 2, 3})
		union := s1.Union(s2)
		testutils.Assert(t, "union.Size()", 3, union.Size())
		testutils.Assert(t, "union.Contains(1)", true, union.Contains(1))
		testutils.Assert(t, "union.Contains(2)", true, union.Contains(2))
		testutils.Assert(t, "union.Contains(3)", true, union.Contains(3))
	})
}

func TestIntersection(t *testing.T) {
	t.Run("NoIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{4, 5})
		intersection := s1.Intersection(s2)
		testutils.Assert(t, "intersection.Size()", 0, intersection.Size())
	})

	t.Run("PartialIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{2, 3, 4, 5})
		intersection := s1.Intersection(s2)
		testutils.Assert(t, "intersection.Size()", 2, intersection.Size())
		testutils.Assert(t, "intersection.Contains(2)", true, intersection.Contains(2))
		testutils.Assert(t, "intersection.Contains(3)", true, intersection.Contains(3))
	})

	t.Run("FullIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{1, 2, 3})
		intersection := s1.Intersection(s2)
		testutils.Assert(t, "intersection.Size()", 3, intersection.Size())
		testutils.Assert(t, "intersection.Contains(1)", true, intersection.Contains(1))
		testutils.Assert(t, "intersection.Contains(2)", true, intersection.Contains(2))
		testutils.Assert(t, "intersection.Contains(3)", true, intersection.Contains(3))
	})
}

func TestDifference(t *testing.T) {
	t.Run("NoIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{4, 5, 6})
		difference := s1.Difference(s2)
		testutils.Assert(t, "difference.Size()", 3, difference.Size())
	})

	t.Run("SomeIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{3, 4, 5})
		difference := s1.Difference(s2)
		testutils.Assert(t, "difference.Size()", 2, difference.Size())
	})

	t.Run("FullIntersection", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{1, 2, 3})
		difference := s1.Difference(s2)
		testutils.Assert(t, "difference.Size()", 0, difference.Size())
	})
}

func TestIsSubset(t *testing.T) {
	t.Run("ProperSubset", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		isSubset := s.IsSubset(NewFromSlice([]int{1, 2, 3, 4}))
		testutils.Assert(t, "isSubset", true, isSubset)
	})

	t.Run("Subset", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		isSubset := s.IsSubset(NewFromSlice([]int{1, 2, 3}))
		testutils.Assert(t, "isSubset", true, isSubset)
	})

	t.Run("NotSubset", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		isSubset := s.IsSubset(NewFromSlice([]int{1, 2, 4}))
		testutils.Assert(t, "isSubset", false, isSubset)
	})
}

func TestIsSuperset(t *testing.T) {
	t.Run("Superset", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		isSuperset := s.IsSuperset(NewFromSlice([]int{1, 2}))
		testutils.Assert(t, "isSuperset", true, isSuperset)
	})

	t.Run("NotSuperset", func(t *testing.T) {
		s := NewFromSlice([]int{1, 2, 3})
		isSuperset := s.IsSuperset(NewFromSlice([]int{1, 2, 4}))
		testutils.Assert(t, "isSuperset", false, isSuperset)
	})
}

func TestEquals(t *testing.T) {
	t.Run("Equals", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{1, 2, 3})
		equals := s1.Equals(s2)
		testutils.Assert(t, "equals", true, equals)
	})

	t.Run("NotEqualsSize", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{1, 2})
		equals := s1.Equals(s2)
		testutils.Assert(t, "equals", false, equals)
	})

	t.Run("NotEquals", func(t *testing.T) {
		s1 := NewFromSlice([]int{1, 2, 3})
		s2 := NewFromSlice([]int{1, 2, 4})
		equals := s1.Equals(s2)
		testutils.Assert(t, "equals", false, equals)
	})
}
