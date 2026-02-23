package minheap

// Min-Heap navigation in an array-backed binary tree:
//
// For a node at index i:
//   - Parent:      (i - 1) / 2
//   - Left child:  2*i + 1
//   - Right child: 2*i + 2
//
// Visual example:
//
//	              1[0]
//	              /   \
//	            3[1]  2[2]
//	            / \   /
//	         5[3] 7[4] 4[5]
//
//	Array: [ 1 | 3 | 2 | 5 | 7 | 4 ]
//	Index:   0   1   2   3   4   5
//
//	Each parent ≤ its children:
//	  1 ≤ 3, 1 ≤ 2, 3 ≤ 5, 3 ≤ 7, 2 ≤ 4  ✓
//
//	1[0] → parent: -       left: [1]3   	 right: [2]2
//	3[1] → parent: [0]1    left: [3]5   	 right: [4]7
//	2[2] → parent: [0]1    left: [5]4   	 right: -
//	5[3] → parent: [1]3    left: -           right: -
//	7[4] → parent: [1]3    left: -           right: -
//	4[5] → parent: [2]2    left: -           right: -

type MinHeap struct {
	nodes []int
}

// Size returns the number of nodes in the heap
func (h *MinHeap) Size() int {
	return len(h.nodes)
}

// Swap swaps the nodes at indices i and j
func (h *MinHeap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// Push adds a value to the heap
func (h *MinHeap) Push(value int) {
	h.nodes = append(h.nodes, value)
	index := len(h.nodes) - 1
	h.siftUp(index)
}

// Pop removes and returns the root node (the node with the smallest value)
func (h *MinHeap) Pop() int {
	if h.Size() == 0 {
		return 0
	}

	root := h.nodes[0]
	lastIndex := h.Size() - 1
	h.nodes[0] = h.nodes[lastIndex]
	h.nodes = h.nodes[:lastIndex]
	h.siftDown(0)
	return root
}

// Peek returns the root node (the node with the smallest value) without removing it from the heap
func (h *MinHeap) Peek() int {
	if h.Size() == 0 {
		return 0
	}

	return h.nodes[0]
}

// siftUp ensures the heap property is maintained by moving the node up the heap
func (h *MinHeap) siftUp(index int) {
	if h.Size() == 0 || index < 0 || index >= h.Size() {
		return
	}

	for index > 0 {
		parentIndex := parent(index)
		if h.nodes[index] >= h.nodes[parentIndex] {
			break
		}
		h.Swap(index, parentIndex)
		index = parentIndex
	}
}

// siftDown ensures the heap property is maintained by moving the node down the heap
func (h *MinHeap) siftDown(index int) {
	if h.Size() == 0 || index < 0 || index >= h.Size() {
		return
	}

	for h.hasChild(index) {
		smallestChildIndex := h.smallestChild(index)
		if smallestChildIndex == -1 {
			break
		}

		// if the node is smaller than or equal to the smallest child, then break
		if h.nodes[index] <= h.nodes[smallestChildIndex] {
			break
		}

		h.Swap(index, smallestChildIndex)

		index = smallestChildIndex
	}
}

func (h *MinHeap) hasChild(index int) bool {
	return leftChild(index) < h.Size() || rightChild(index) < h.Size()
}

func (h *MinHeap) smallestChild(index int) int {
	leftChildIndex := leftChild(index)
	rightChildIndex := rightChild(index)

	// no child
	if leftChildIndex >= h.Size() {
		return -1
	}

	// only left child
	if rightChildIndex >= h.Size() {
		return leftChildIndex
	}

	// both children, then compare them
	if h.nodes[leftChildIndex] < h.nodes[rightChildIndex] {
		return leftChildIndex
	}
	return rightChildIndex
}

// parent returns the index of the parent node
func parent(index int) int {
	return (index - 1) / 2
}

// leftChild returns the index of the left child node
func leftChild(index int) int {
	return 2*index + 1
}

// rightChild returns the index of the right child node
func rightChild(index int) int {
	return 2*index + 2
}
