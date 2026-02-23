package maxheap

// Max-Heap navigation in an array-backed binary tree:
//
// For a node at index i:
//   - Parent:      (i - 1) / 2
//   - Left child:  2*i + 1
//   - Right child: 2*i + 2
//
// Visual example:
//
//	              7[0]
//	              /   \
//	            5[1]  6[2]
//	            / \   /
//	         3[3] 1[4] 4[5]
//
//	Array: [ 7 | 5 | 6 | 3 | 1 | 4 ]
//	Index:   0   1   2   3   4   5
//
//	Each parent ≥ its children:
//	  7 ≥ 5, 7 ≥ 6, 5 ≥ 3, 5 ≥ 1, 6 ≥ 4  ✓
//
//	7[0] → parent: -       left: [1]5        right: [2]6
//	5[1] → parent: [0]7    left: [3]3        right: [4]1
//	6[2] → parent: [0]7    left: [5]4        right: -
//	3[3] → parent: [1]5    left: -           right: -
//	1[4] → parent: [1]5    left: -           right: -
//	4[5] → parent: [2]6    left: -           right: -

type MaxHeap struct {
	nodes []int
}

// Size returns the number of nodes in the heap
func (h *MaxHeap) Size() int {
	return len(h.nodes)
}

// Swap swaps the nodes at indices i and j
func (h *MaxHeap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// Push adds a value to the heap
func (h *MaxHeap) Push(value int) {
	h.nodes = append(h.nodes, value)
	index := len(h.nodes) - 1
	h.siftUp(index)
}

// Pop removes and returns the root node (the node with the largest value)
func (h *MaxHeap) Pop() int {
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

// Peek returns the root node (the node with the largest value) without removing it from the heap
func (h *MaxHeap) Peek() int {
	if h.Size() == 0 {
		return 0
	}

	return h.nodes[0]
}

// siftUp ensures the heap property is maintained by moving the node up the heap
func (h *MaxHeap) siftUp(index int) {
	if h.Size() == 0 || index < 0 || index >= h.Size() {
		return
	}

	for index > 0 {
		parentIndex := parent(index)
		if h.nodes[index] <= h.nodes[parentIndex] {
			break
		}
		h.Swap(index, parentIndex)
		index = parentIndex
	}
}

// siftDown ensures the heap property is maintained by moving the node down the heap
func (h *MaxHeap) siftDown(index int) {
	if h.Size() == 0 || index < 0 || index >= h.Size() {
		return
	}

	for h.hasChild(index) {
		largestChildIndex := h.largestChild(index)
		if largestChildIndex == -1 {
			break
		}

		// if the node is larger than or equal to the largest child, then break
		if h.nodes[index] >= h.nodes[largestChildIndex] {
			break
		}

		h.Swap(index, largestChildIndex)

		index = largestChildIndex
	}
}

func (h *MaxHeap) hasChild(index int) bool {
	return leftChild(index) < h.Size() || rightChild(index) < h.Size()
}

func (h *MaxHeap) largestChild(index int) int {
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
	if h.nodes[leftChildIndex] > h.nodes[rightChildIndex] {
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
