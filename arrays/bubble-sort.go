package arrays

// BubbleSort sorts an array of integers using the bubble sort algorithm
// The algorithm works by repeatedly swapping the adjacent elements if they are in the wrong order
// The time complexity of the algorithm is O(n^2) in the worst case and O(n) in the best case
// The space complexity of the algorithm is O(1)
func BubbleSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		hasSwapped := false
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				hasSwapped = true
			}
		}
		if !hasSwapped {
			break
		}
	}
}
