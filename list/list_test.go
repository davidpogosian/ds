package list

import (
	"testing"

	"github.com/davidpogosian/ds/testutils"
)

func TestNewEmpty(t *testing.T) {
	l := NewEmpty[int]()
	testutils.Assert(t, "l.Size()", 0, l.Size())
	testutils.Assert(t, "l.String()", "[]", l.String())
}

func TestNewFromSlice(t *testing.T) {
	l := NewFromSlice([]int{1, 2, 3})
	testutils.Assert(t, "l.Size()", 3, l.Size())
	testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
}

func TestInsertFront(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		l := NewEmpty[int]()
		l.InsertFront(1)
		testutils.Assert(t, "l.String()", "[1]", l.String())
		l.InsertFront(2)
		testutils.Assert(t, "l.String()", "[2 1]", l.String())
		l.InsertFront(3)
		testutils.Assert(t, "l.String()", "[3 2 1]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int]()
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
		l := NewEmpty[int]()
		l.InsertBack(1)
		testutils.Assert(t, "l.String()", "[1]", l.String())
		l.InsertBack(2)
		testutils.Assert(t, "l.String()", "[1 2]", l.String())
		l.InsertBack(3)
		testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int]()
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
				l := NewEmpty[int]()
				err := l.InsertPosition(1, 0)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1]", l.String())
			})

			t.Run("Index1", func(t *testing.T) {
				l := NewEmpty[int]()
				err := l.InsertPosition(1, 1)
				if err == nil {
					t.Fatal("Inserted into an empty List at position 1.")
				}
				testutils.Assert(t, "l.String()", "[]", l.String())
			})
		})

		t.Run("NotEmpty", func(t *testing.T) {
			t.Run("Front", func(t *testing.T) {
				l := NewFromSlice([]int{2, 3})
				err := l.InsertPosition(1, 0)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
			})

			t.Run("Index1", func(t *testing.T) {
				l := NewFromSlice([]int{1, 3})
				err := l.InsertPosition(2, 1)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
			})

			t.Run("Back", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2})
				err := l.InsertPosition(3, 2)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2 3]", l.String())
			})

			t.Run("Index3", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2})
				err := l.InsertPosition(3, 3)
				if err == nil {
					t.Fatal("Inserted into a List of size 2 at position 3.")
				}
			})
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int]()
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
		l := NewEmpty[int]()
		l.Reverse()
		testutils.Assert(t, "l.String()", "[]", l.String())
	})

	t.Run("Size1", func(t *testing.T) {
		l := NewFromSlice([]int{1})
		l.Reverse()
		testutils.Assert(t, "l.String()", "[1]", l.String())
	})

	t.Run("Size3", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3})
		l.Reverse()
		testutils.Assert(t, "l.String()", "[3 2 1]", l.String())
	})
}

func TestRemoveFront(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3})
		l.RemoveFront()
		testutils.Assert(t, "l.String()", "[2 3]", l.String())
		l.RemoveFront()
		testutils.Assert(t, "l.String()", "[3]", l.String())
		l.RemoveFront()
		testutils.Assert(t, "l.String()", "[]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int]()
		for i := 0; i < 1000; i++ {
			l.InsertBack(i)
		}
		testutils.Assert(t, "l.Size()", 1000, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			err := l.RemoveFront()
			return err
		})
		testutils.Assert(t, "l.Size()", 0, l.Size())
	})
}

func TestRemoveBack(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		l := NewFromSlice([]int{1, 2, 3})
		l.RemoveBack()
		testutils.Assert(t, "l.String()", "[1 2]", l.String())
		l.RemoveBack()
		testutils.Assert(t, "l.String()", "[1]", l.String())
		l.RemoveBack()
		testutils.Assert(t, "l.String()", "[]", l.String())
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int]()
		for i := 0; i < 1000; i++ {
			l.InsertBack(i)
		}
		testutils.Assert(t, "l.Size()", 1000, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			err := l.RemoveBack()
			return err
		})
		testutils.Assert(t, "l.Size()", 0, l.Size())
	})
}

func TestRemovePosition(t *testing.T) {
	t.Run("Sequential", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			l := NewEmpty[int]()
			err := l.RemovePosition(0)
			if err == nil {
				t.Fatal("Removed from an empty List.")
			}
		})

		t.Run("NotEmpty", func(t *testing.T) {
			t.Run("Front", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3})
				err := l.RemovePosition(0)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[2 3]", l.String())
			})

			t.Run("Index1", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3})
				err := l.RemovePosition(1)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 3]", l.String())
			})

			t.Run("Back", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3})
				err := l.RemovePosition(2)
				if err != nil {
					t.Fatal(err)
				}
				testutils.Assert(t, "l.String()", "[1 2]", l.String())
			})

			t.Run("Index3", func(t *testing.T) {
				l := NewFromSlice([]int{1, 2, 3})
				err := l.RemovePosition(3)
				if err == nil {
					t.Fatal("Removed from a List of size 3 at position 3.")
				}
			})
		})
	})

	t.Run("Concurrent", func(t *testing.T) {
		l := NewEmpty[int]()
		for i := 0; i < 1000; i++ {
			l.InsertBack(i)
		}
		testutils.Assert(t, "l.Size()", 1000, l.Size())
		testutils.ConcurrentOperations(t, 10, 100, func() error {
			err := l.RemovePosition(0)
			return err
		})
		testutils.Assert(t, "l.Size()", 0, l.Size())
	})
}
