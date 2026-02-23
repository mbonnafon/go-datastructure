package weightedminheap

// Weighted Min-Heap navigation in an array-backed binary tree:
//
// For a node at index i:
//   - Parent:      (i - 1) / 2
//   - Left child:  2*i + 1
//   - Right child: 2*i + 2
//
// Visual example (values shown as {value, weight}):
//
//	              {a,1}[0]
//	              /       \
//	        {b,3}[1]     {c,2}[2]
//	        /      \       /
//	   {d,5}[3] {e,7}[4] {f,4}[5]
//
//	Array: [ {a,1} | {b,3} | {c,2} | {d,5} | {e,7} | {f,4} ]
//	Index:     0       1       2       3       4       5
//
//	Each parent's weight ≤ its children's weights:
//	  1 ≤ 3, 1 ≤ 2, 3 ≤ 5, 3 ≤ 7, 2 ≤ 4  ✓
//
//	{a,1}[0] → parent: -    	   left: [1]{b,3}    right: [2]{c,2}
//	{b,3}[1] → parent: [0]{a,1}    left: [3]{d,5}    right: [4]{e,7}
//	{c,2}[2] → parent: [0]{a,1}    left: [5]{f,4}    right: -
//	{d,5}[3] → parent: [1]{b,3}    left: -           right: -
//	{e,7}[4] → parent: [1]{b,3}    left: -           right: -
//	{f,4}[5] → parent: [2]{c,2}    left: -           right: -

type MinHeap struct {
	nodes []Node
}

type Node struct {
	Value  int
	Weight int
}

// Size returns the number of nodes in the heap
func (h *MinHeap) Size() int {
	return len(h.nodes)
}

// Swap swaps the nodes at indices i and j
func (h *MinHeap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// Push adds a node to the heap
func (h *MinHeap) Push(node Node) {
	h.nodes = append(h.nodes, node)
	index := len(h.nodes) - 1
	h.siftUp(index)
}

// Pop removes and returns the root node (the node with the smallest weight)
func (h *MinHeap) Pop() Node {
	if h.Size() == 0 {
		return Node{}
	}

	root := h.nodes[0]
	lastIndex := h.Size() - 1
	h.nodes[0] = h.nodes[lastIndex]
	h.nodes = h.nodes[:lastIndex]
	h.siftDown(0)
	return root
}

// Peek returns the root node (the node with the smallest weight) without removing it from the heap
func (h *MinHeap) Peek() Node {
	if h.Size() == 0 {
		return Node{}
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
		if h.nodes[index].Weight >= h.nodes[parentIndex].Weight {
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
		if h.nodes[index].Weight <= h.nodes[smallestChildIndex].Weight {
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
	if h.nodes[leftChildIndex].Weight < h.nodes[rightChildIndex].Weight {
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
