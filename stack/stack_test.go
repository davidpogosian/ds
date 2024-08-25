package stack

import (
	"fmt"
	"testing"
)

func same(sliceA []int, sliceB []int) error {
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

func TestPush(t *testing.T) {
	stack := NewEmpty()
	slice := []int{1,2,3}
	for i := range slice {
		stack.Push(slice[i])
	}
	sliceFromStack, err := stack.ToSlice()
	if err != nil {
		t.Fatal(err)
	}
	err = same(slice, sliceFromStack)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPop(t *testing.T) {

}
