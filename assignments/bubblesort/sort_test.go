package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func IsSorted(sli []int) bool {
	for idx, elem := range sli {
		switch idx {
		case 0:
			continue
		default:
			if elem < sli[idx-1] {
				return false
			}
		}
	}
	return true
}

func SlicesEqual[T comparable](sliA, sliB []T) bool {
	if len(sliA) != len(sliB) {
		return false
	}
	for idx, elem := range sliA {
		if sliB[idx] != elem {
			return false
		}
	}
	return true
}

func TestSwap(t *testing.T) {
	a := []int{1, 2}
	Swap(a, 0)
	assert.True(t, reflect.DeepEqual(a, []int{2, 1}))

	a = []int{1, 2} // reset
	Swap(a, 1)      // no change (out of range)
	assert.True(t, reflect.DeepEqual(a, []int{1, 2}))
	Swap(a, 2) // no change (out of range)
	assert.True(t, reflect.DeepEqual(a, []int{1, 2}))
	Swap(a, -1) // no change (out of range)
	assert.True(t, reflect.DeepEqual(a, []int{1, 2}))
}

func TestSort(t *testing.T) {
	intsli := rand.Perm(100)
	sorted := make([]int, 100)
	copy(sorted, intsli)
	sort.Ints(sorted)
	BubbleSort(intsli)
	assert.True(t, reflect.DeepEqual(intsli, sorted))
}
