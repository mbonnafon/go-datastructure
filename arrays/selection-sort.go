package arrays

// SelectionSort sorts an array of integers using the selection sort algorithm
// The algorithm works by repeatedly finding the minimum element from the unsorted portion
// and placing it at the beginning of the unsorted portion
// The time complexity of the algorithm is O(n^2) in all cases (best, average, and worst)
// The space complexity of the algorithm is O(1)
func SelectionSort(arr []int) {
	cursor := 0
	for cursor < len(arr)-1 {
		minIndex := cursor
		for i := cursor + 1; i < len(arr); i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			}
		}
		arr[cursor], arr[minIndex] = arr[minIndex], arr[cursor]
		cursor++
	}
}
