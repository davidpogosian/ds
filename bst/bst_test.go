package bst

import (
	"testing"

	"github.com/davidpogosian/ds/comparators"
	"github.com/davidpogosian/ds/testutils"
)

func TestNewEmpty(t *testing.T) {
	bst := NewEmpty[int, string](comparators.ComparatorInt)
	testutils.Assert(t, "bst.Size()", 0, bst.Size())
}

func TestInsert(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(1, "one")
		bst.Insert(2, "two")
		bst.Insert(3, "three")
		testutils.Assert(t, "bst.Size()", 3, bst.Size())
	})

	t.Run("Concurrent", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			bst.Insert(1, "ello")
			return nil
		})
		testutils.Assert(t, "bst.Size()", 1000, bst.Size())
		testutils.Assert(t, "bst.Height()", 999, bst.Height())
	})
}

func TestSearch(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(1, "one")
		bst.Insert(2, "two")
		bst.Insert(3, "three")
		two, err := bst.Search(2)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "two", "two", two)
	})

	t.Run("NotExists", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(1, "one")
		bst.Insert(2, "two")
		bst.Insert(3, "three")
		_, err := bst.Search(123)
		if err == nil {
			t.Fatalf("Searched BST for key that does not exist with no error.")
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		t.Run("Exists", func(t *testing.T) {
			bst := NewEmpty[int, string](comparators.ComparatorInt)
			bst.Insert(1, "one")
			bst.Insert(2, "two")
			bst.Insert(3, "three")
			one, err := bst.Remove(1)
			if err != nil {
				t.Fatal(err)
			}
			testutils.Assert(t, "one", "one", one)
			testutils.Assert(t, "bst.Height()", 1, bst.Height())
			testutils.Assert(t, "bst.Size()", 2, bst.Size())
		})

		t.Run("NotExists", func(t *testing.T) {
			bst := NewEmpty[int, string](comparators.ComparatorInt)
			bst.Insert(1, "one")
			bst.Insert(2, "two")
			bst.Insert(3, "three")
			_, err := bst.Remove(4)
			if err == nil {
				t.Fatal("Removed from BST by key that does not exist.")
			}
			testutils.Assert(t, "bst.Height()", 2, bst.Height())
			testutils.Assert(t, "bst.Size()", 3, bst.Size())
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		for i := 0; i < 1000; i++ {
			bst.Insert(1, "one")
		}
		testutils.ConcurrentOperations(t, 10, 100,  func() error {
			_, err := bst.Remove(1)
			return err
		})
		testutils.Assert(t, "bst.Size()", 0, bst.Size())
		testutils.Assert(t, "bst.Height()", -1, bst.Height())
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		testutils.Assert(t, "bst.Size()", 0, bst.Size())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(1, "one")
		bst.Insert(2, "two")
		bst.Insert(3, "three")
		testutils.Assert(t, "bst.Size()", 3, bst.Size())
	})
}

func TestInOrderTraversal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		s := bst.InOrderTraversal()
		testutils.AssertSlices(t, s, []int{})
	})

	t.Run("NotEmpty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(4, "")
		bst.Insert(8, "")
		bst.Insert(2, "")
		bst.Insert(6, "")
		bst.Insert(3, "")
		bst.Insert(1, "")
		s := bst.InOrderTraversal()
		testutils.AssertSlices(t, s, []int{1, 2, 3, 4, 6, 8})
	})
}

func TestFindMin(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		_, err := bst.FindMin()
		if err == nil {
			t.Fatal("Found min on an empty BST.")
		}
	})

	t.Run("NotEmpty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(4, "")
		bst.Insert(8, "")
		bst.Insert(2, "")
		bst.Insert(6, "")
		bst.Insert(3, "")
		bst.Insert(1, "")
		one, err := bst.FindMin()
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "one", 1, one)
	})
}

func TestFindMax(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		_, err := bst.FindMax()
		if err == nil {
			t.Fatal("Found max on an empty BST.")
		}
	})

	t.Run("NotEmpty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(4, "")
		bst.Insert(8, "")
		bst.Insert(2, "")
		bst.Insert(6, "")
		bst.Insert(3, "")
		bst.Insert(1, "")
		eight, err := bst.FindMax()
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "eight", 8, eight)
	})
}

func TestPreOrderTraversal(t *testing.T) {
	bst := NewEmpty[int, string](comparators.ComparatorInt)
	bst.Insert(4, "")
	bst.Insert(8, "")
	bst.Insert(2, "")
	bst.Insert(6, "")
	bst.Insert(3, "")
	bst.Insert(1, "")
	slice := bst.PreOrderTraversal()
	testutils.AssertSlices(t, slice, []int{4, 2, 1, 3, 8, 6})
}

func TestPostOrderTraversal(t *testing.T) {
	bst := NewEmpty[int, string](comparators.ComparatorInt)
	bst.Insert(4, "")
	bst.Insert(8, "")
	bst.Insert(2, "")
	bst.Insert(6, "")
	bst.Insert(3, "")
	bst.Insert(1, "")
	slice := bst.PostOrderTraversal()
	testutils.AssertSlices(t, slice, []int{1, 3, 2, 6, 8, 4})
}

func TestHeight(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		testutils.Assert(t, "bst.Height()", -1, bst.Height())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(4, "")
		bst.Insert(8, "")
		bst.Insert(2, "")
		bst.Insert(6, "")
		bst.Insert(3, "")
		bst.Insert(1, "")
		testutils.Assert(t, "bst.Height()", 2, bst.Height())
	})
}

func TestClear(t *testing.T) {
	bst := NewEmpty[int, string](comparators.ComparatorInt)
	bst.Insert(1, "")
	bst.Insert(1, "")
	bst.Insert(1, "")
	bst.Clear()
	testutils.Assert(t, "bst.Size()", 0, bst.Size())
	testutils.Assert(t, "bst.Height()", -1, bst.Height())
}

func TestCopy(t *testing.T) {
	t.Run("Relation", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(1, "hi")
		copy := bst.Copy()
		_, err := bst.Remove(1)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "copy.Size()", 1, copy.Size())
		hi, err := copy.Search(1)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "hi", "hi", hi)
	})

	t.Run("Large", func(t *testing.T) {
		bst := NewEmpty[int, string](comparators.ComparatorInt)
		bst.Insert(10, "")
		bst.Insert(8, "")
		bst.Insert(12, "")
		bst.Insert(13, "")
		bst.Insert(11, "")
		bst.Insert(6, "")
		bst.Insert(7, "")
		testutils.AssertSlices(t, bst.PreOrderTraversal(), []int{10, 8, 6, 7, 12, 11, 13})
		copy := bst.Copy()
		bst.Clear()
		testutils.AssertSlices(t, copy.PreOrderTraversal(), []int{10, 8, 6, 7, 12, 11, 13})
	})
}
