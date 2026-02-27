package arrays

// InsertionSort sorts an array of integers using the insertion sort algorithm
// The algorithm works by repeatedly inserting the current element into the correct position in the sorted portion
// The time complexity of the algorithm is O(n) in the best case and O(n^2) in the average and worst cases
// The space complexity of the algorithm is O(1)
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
