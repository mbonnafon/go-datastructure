package maxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap_Pop(t *testing.T) {
	for _, test := range []struct {
		name         string
		values       []int
		expectedPop  int
		expectedSize int
	}{
		{
			name:         "when the heap is empty pop should return zero",
			values:       []int{},
			expectedPop:  0,
			expectedSize: 0,
		},
		{
			name:         "when the heap has only one value pop should return the value",
			values:       []int{1},
			expectedPop:  1,
			expectedSize: 0,
		},
		{
			name:         "should pop the values in the correct order when the values are in ascending order",
			values:       []int{1, 2, 3},
			expectedPop:  3,
			expectedSize: 2,
		},
		{
			name:         "should pop the values in the correct order when the values are in descending order",
			values:       []int{3, 2, 1},
			expectedPop:  3,
			expectedSize: 2,
		},
		{
			name:         "should pop the values in the correct order when the values are in random order",
			values:       []int{2, 3, 1, 4},
			expectedPop:  4,
			expectedSize: 3,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			heap := &MaxHeap{}
			for _, val := range test.values {
				heap.Push(val)
			}
			assert.Equal(t, heap.Pop(), test.expectedPop, "the pop should be the correct value")
			assert.Equal(t, heap.Size(), test.expectedSize, "the size should be correct")
		})
	}
}

func TestMaxHeap_Push(t *testing.T) {
	for _, test := range []struct {
		name         string
		values       []int
		expectedPeek int
	}{
		{
			name:         "should push the values in the correct order when the values are in ascending order",
			values:       []int{1, 2, 3},
			expectedPeek: 3,
		},
		{
			name:         "should push the values in the correct order when the values are in descending order",
			values:       []int{3, 2, 1},
			expectedPeek: 3,
		},
		{
			name:         "should push the values in the correct order when the values are in random order",
			values:       []int{2, 3, 1, 4},
			expectedPeek: 4,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			heap := &MaxHeap{}
			for _, val := range test.values {
				heap.Push(val)
			}
			assert.Equal(t, heap.Peek(), test.expectedPeek)
		})
	}
}
