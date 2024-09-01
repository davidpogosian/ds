package list

import (
	"testing"

	"github.com/davidpogosian/ds/comparators"
	"github.com/davidpogosian/ds/testutils"
)

func TestNewEmpty(t *testing.T) {
	l := NewEmpty[int](comparators.ComparatorInt)
	testutils.Assert(t, "l.Size()", 0, l.Size())
	testutils.Assert(t, "l.String()", "[]", l.String())
}

func TestNewFromSlice(t *testing.T) {
	l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
	testutils.Assert(t, "l.Size()", 3, l.Size())
	testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
}

func TestInsertFront(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		l.InsertFront(1)
		testutils.Assert(t, "l.String()", "[1]", l.String())
		l.InsertFront(2)
		testutils.Assert(t, "l.String()", "[2 1]", l.String())
		l.InsertFront(3)
		testutils.Assert(t, "l.String()", "[3 2 1]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "l.Size()", 0, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			l.InsertFront(1)
			return nil
		})
 		testutils.Assert(t, "l.Size()", 1000, l.Size())
	})
}

func TestInsertBack(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		l.InsertBack(1)
		testutils.Assert(t, "l.String()", "[1]", l.String())
		l.InsertBack(2)
		testutils.Assert(t, "l.String()", "[1 2]", l.String())
		l.InsertBack(3)
		testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "l.Size()", 0, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			l.InsertBack(1)
			return nil
		})
 		testutils.Assert(t, "l.Size()", 1000, l.Size())
	})
}

func TestInsertPosition(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			t.Run("Front", func(t *testing.T) {
				l := NewEmpty[int](comparators.ComparatorInt)
				err := l.InsertPosition(1, 0)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1]", l.String())
			})

			t.Run("Index1", func(t *testing.T) {
				l := NewEmpty[int](comparators.ComparatorInt)
				err := l.InsertPosition(1, 1)
				if err == nil {
					t.Fatal("Inserted into an empty List at position 1.")
				}
				testutils.Assert(t, "l.String()", "[]", l.String())
			})
		})

		t.Run("NotEmpty", func(t *testing.T) {
			t.Run("Front", func(t *testing.T) {
				l := NewFromSlice([]int{2, 3}, comparators.ComparatorInt)
				err := l.InsertPosition(1, 0)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
			})

			t.Run("Index1", func(t *testing.T) {
				l := NewFromSlice([]int{1, 3}, comparators.ComparatorInt)
				err := l.InsertPosition(2, 1)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
			})

			t.Run("Back", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2}, comparators.ComparatorInt)
				err := l.InsertPosition(3, 2)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
			})

			t.Run("Index3", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2}, comparators.ComparatorInt)
				err := l.InsertPosition(3, 3)
				if err == nil {
					t.Fatal("Inserted into a List of size 2 at position 3.")
				}
			})
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "l.Size()", 0, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			err := l.InsertPosition(1, 0)
			return err
		})
 		testutils.Assert(t, "l.Size()", 1000, l.Size())
	})
}

func TestReverse(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		l.Reverse()
		testutils.Assert(t, "l.String()", "[]", l.String())
	})

	t.Run("Size1", func(t *testing.T) {
		l := NewFromSlice([]int{1}, comparators.ComparatorInt)
		l.Reverse()
		testutils.Assert(t, "l.String()", "[1]", l.String())
	})

	t.Run("Size3", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		l.Reverse()
		testutils.Assert(t, "l.String()", "[3 2 1]", l.String())
	})
}

func TestRemoveFront(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		one, err := l.RemoveFront()
		testutils.Assert(t, "one", 1, one)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "l.String()", "[2 3]", l.String())
		two, err := l.RemoveFront()
		testutils.Assert(t, "two", 2, two)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "l.String()", "[3]", l.String())
		three, err := l.RemoveFront()
		testutils.Assert(t, "three", 3, three)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "l.String()", "[]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		for i := 0; i < 1000; i++ {
			l.InsertBack(i)
		}
		testutils.Assert(t, "l.Size()", 1000, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			_, err := l.RemoveFront()
			return err
		})
		testutils.Assert(t, "l.Size()", 0, l.Size())
	})
}

func TestRemoveBack(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		three, err := l.RemoveBack()
		testutils.Assert(t, "three", 3, three)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "l.String()", "[1 2]", l.String())
		two, err := l.RemoveBack()
		testutils.Assert(t, "two", 2, two)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "l.String()", "[1]", l.String())
		one, err := l.RemoveBack()
		testutils.Assert(t, "one", 1, one)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "l.String()", "[]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		for i := 0; i < 1000; i++ {
			l.InsertBack(i)
		}
		testutils.Assert(t, "l.Size()", 1000, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			_, err := l.RemoveBack()
			return err
		})
		testutils.Assert(t, "l.Size()", 0, l.Size())
	})
}

func TestRemovePosition(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			l := NewEmpty[int](comparators.ComparatorInt)
			_, err := l.RemovePosition(0)
			if err == nil {
				t.Fatal("Removed from an empty List.")
			}
		})

		t.Run("NotEmpty", func(t *testing.T) {
			t.Run("Front", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
				one, err := l.RemovePosition(0)
				testutils.Assert(t, "one", 1, one)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[2 3]", l.String())
			})

			t.Run("Index1", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
				two, err := l.RemovePosition(1)
				testutils.Assert(t, "two", 2, two)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 3]", l.String())
			})

			t.Run("Back", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
				three, err := l.RemovePosition(2)
				testutils.Assert(t, "three", 3, three)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2]", l.String())
			})

			t.Run("Index3", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
				_, err := l.RemovePosition(3)
				if err == nil {
					t.Fatal("Removed from a List of size 3 at position 3.")
				}
			})
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		for i := 0; i < 1000; i++ {
			l.InsertBack(i)
		}
		testutils.Assert(t, "l.Size()", 1000, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			_, err := l.RemovePosition(0)
			return err
		})
		testutils.Assert(t, "l.Size()", 0, l.Size())
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "l.Size()", 0, l.Size())
	})

	t.Run("Size1", func(t *testing.T) {
		l := NewFromSlice([]int{2}, comparators.ComparatorInt)
		testutils.Assert(t, "l.Size()", 1, l.Size())
	})

	t.Run("Size5", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3, 4, 5}, comparators.ComparatorInt)
		testutils.Assert(t, "l.Size()", 5, l.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "l.IsEmpty()", true, l.IsEmpty())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3, 4, 5}, comparators.ComparatorInt)
		testutils.Assert(t, "l.IsEmpty()", false, l.IsEmpty())
	})
}

func TestClear(t *testing.T) {
	l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
	l.Clear()
	testutils.Assert(t, "l.Size()", 0, l.Size())
}

func TestGet(t *testing.T) {
	t.Run("ValidIndex", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		two, err := l.Get(1)
		if err != nil {
			t.Fatal(err)
		}
		testutils.Assert(t, "two", 2, two)
	})

	t.Run("InvalidIndex", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		_, err := l.Get(100)
		if err == nil {
			t.Fatal("Got item at index 100 from List of length 3.")
		}
	})
}

func TestCopy(t *testing.T) {
	l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
	copy := l.Copy()
	l.InsertFront(100)
	testutils.Assert(t, "copy.Size()", 3, copy.Size())
}

func TestFind(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		one := l.Find(2)
		testutils.Assert(t, "one", 1, one)
	})

	t.Run("NotExists", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		negativeOne := l.Find(90)
		testutils.Assert(t, "negativeOne", -1, negativeOne)
	})
}

func TestToSlice(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		slice := l.ToSlice()
		testutils.Assert(t, "len(slice)", 0, len(slice))
	})

	t.Run("NotEmpty", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		slice := l.ToSlice()
		testutils.Assert(t, "len(slice)", 3, len(slice))
	})
}

func TestToString(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		l := NewEmpty[int](comparators.ComparatorInt)
		testutils.Assert(t, "l.String()", "[]", l.String())
	})

	t.Run("NotEmpty", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3}, comparators.ComparatorInt)
		testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
	})
}
