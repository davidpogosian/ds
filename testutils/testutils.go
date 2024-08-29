package testutils

import (
	"fmt"
	"sync"
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

func ConcurrentOperations(t *testing.T, threads int, repetitions int, function func() error) {
	var waitGroup sync.WaitGroup
	errCh := make(chan error, threads * repetitions)
	for i := 0; i < threads; i++ {
		waitGroup.Add(1)
		go func(errCh chan error) {
			defer waitGroup.Done()
			for j := 0; j < repetitions; j++ {
				err := function()
				if err != nil {
					errCh <- err
				}
			}
		}(errCh)
	}
	waitGroup.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			t.Fatal(err)
		}
	}
}
