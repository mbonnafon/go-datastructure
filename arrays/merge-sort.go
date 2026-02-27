package arrays

// MergeSort sorts an array of integers using the merge sort algorithm
// The algorithm works by dividing the array into two halves, sorting each half recursively,
// and then merging the two sorted halves back together
//
// Example:
//
//	Input: [5, 2, 8, 1, 9]
//
//	Divide:
//	         [5, 2, 8, 1, 9]
//	         /              \
//	    [5, 2, 8]        [1, 9]
//	    /       \         /    \
//	 [5, 2]     [8]     [1]    [9]
//	 /    \
//	[5]   [2]
//
//	Merge (bottom-up):
//	[5] [2] → [2, 5]
//	[2, 5] [8] → [2, 5, 8]
//	[1] [9] → [1, 9]
//	[2, 5, 8] [1, 9] → [1, 2, 5, 8, 9]
//
//	Output: [1, 2, 5, 8, 9]
//
// The time complexity of the algorithm is O(n log n) in all cases (best, average, and worst)
// The space complexity of the algorithm is O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]

	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)

	return merge(sortedLeft, sortedRight)
}

func merge(left []int, right []int) []int {
	sorted := make([]int, len(left)+len(right))

	for i, j := 0, 0; i+j < len(sorted); {
		if i < len(left) && (j >= len(right) || left[i] < right[j]) {
			sorted[i+j] = left[i]
			i++
		} else {
			sorted[i+j] = right[j]
			j++
		}
	}

	return sorted
}
