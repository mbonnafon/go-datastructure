package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	t.Run("when the array is empty", func(t *testing.T) {
		arr := []int{}
		assert.Equal(t, []int{}, MergeSort(arr))
	})

	t.Run("when the array has one element", func(t *testing.T) {
		arr := []int{1}
		assert.Equal(t, []int{1}, MergeSort(arr))
	})

	t.Run("when the array is already sorted", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		assert.Equal(t, []int{1, 2, 3, 4, 5}, MergeSort(arr))
	})

	t.Run("when the array is in the worst-case scenario", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		assert.Equal(t, []int{1, 2, 3, 4, 5}, MergeSort(arr))
	})

	t.Run("when the array has positive and negative numbers", func(t *testing.T) {
		arr := []int{5, -4, 3, -2, 1}
		assert.Equal(t, []int{-4, -2, 1, 3, 5}, MergeSort(arr))
	})

	t.Run("when the array has same numbers", func(t *testing.T) {
		arr := []int{5, 5, 5, 5, 5}
		assert.Equal(t, []int{5, 5, 5, 5, 5}, MergeSort(arr))
	})

	t.Run("when the array has duplicates", func(t *testing.T) {
		arr := []int{5, 3, 8, 4, 2, 5, 3, 8, 4, 2}
		assert.Equal(t, []int{2, 2, 3, 3, 4, 4, 5, 5, 8, 8}, MergeSort(arr))
	})

}
