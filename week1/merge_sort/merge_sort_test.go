package main

import (
	"reflect"
	"testing"
)

func testReturnsSortedArray(t *testing.T) {
	unsorted := []int{4, 3, 2, 1}
	sorted := Sort(unsorted)
	if !reflect.DeepEqual(sorted, []int{1, 2, 3, 4}) {
		t.Error("expected sorted array, got", sorted)
	}
}

func TestSingleElementsArrayJustReturned(t *testing.T) {
	unsorted := []int{5}
	sorted := Sort(unsorted)
	if !reflect.DeepEqual(sorted, unsorted) {
		t.Error("single element array expected to be returned unchanged")
	}
}

func TestSortElementsArray(t *testing.T) {
	unsorted := []int{7, 8, 2, 1}
	sorted := Sort(unsorted)
	if !reflect.DeepEqual(sorted, []int{1, 2, 7, 8}) {
		t.Error("expected sorted array, got", sorted)
	}
}
