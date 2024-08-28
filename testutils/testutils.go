package testutils

import (
	"fmt"
	"testing"
)

func Assert[T comparable](t *testing.T, varname string, expected T, got T) {
	if expected != got {
		t.Fatalf("Expected '%s' to be: %v, instead got: %v", varname, expected, got)
	}
}

func CompareSlices[T comparable](sliceA []T, sliceB []T) error {
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
