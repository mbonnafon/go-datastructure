package weightedmaxheap

// Weighted Max-Heap navigation in an array-backed binary tree:
//
// For a node at index i:
//   - Parent:      (i - 1) / 2
//   - Left child:  2*i + 1
//   - Right child: 2*i + 2
//
// Visual example (values shown as {value, weight}):
//
//	              {a,7}[0]
//	              /       \
//	        {b,5}[1]     {c,6}[2]
//	        /      \       /
//	   {d,3}[3] {e,1}[4] {f,4}[5]
//
//	Array: [ {a,7} | {b,5} | {c,6} | {d,3} | {e,1} | {f,4} ]
//	Index:     0       1       2       3       4       5
//
//	Each parent's weight ≥ its children's weights:
//	  7 ≥ 5, 7 ≥ 6, 5 ≥ 3, 5 ≥ 1, 6 ≥ 4  ✓
//
//	{a,7}[0] → parent: -     	   left: [1]{b,5}    right: [2]{c,6}
//	{b,5}[1] → parent: [0]{a,7}    left: [3]{d,3}    right: [4]{e,1}
//	{c,6}[2] → parent: [0]{a,7}    left: [5]{f,4}    right: -
//	{d,3}[3] → parent: [1]{b,5}    left: -           right: -
//	{e,1}[4] → parent: [1]{b,5}    left: -           right: -
//	{f,4}[5] → parent: [2]{c,6}    left: -           right: -
//
// siftUp:   compare nodes[i] with nodes[(i-1)/2] (parent), swap if child is larger, then i = (i-1)/2
// siftDown: compare nodes[i] with nodes[2*i+1] and nodes[2*i+2] (children), swap with the largest, then i = largest child index
// Always check that the child index is < len(nodes) before accessing it.

type MaxHeap struct {
	nodes []Node
}

type Node struct {
	Value  int
	Weight int
}

// Size returns the number of nodes in the heap
func (h *MaxHeap) Size() int {
	return len(h.nodes)
}

// Swap swaps the nodes at indices i and j
func (h *MaxHeap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// Push adds a node to the heap
func (h *MaxHeap) Push(node Node) {
	h.nodes = append(h.nodes, node)
	index := h.Size() - 1
	h.siftUp(index)
}

// Pop removes and returns the root node (the node with the largest weight)
func (h *MaxHeap) Pop() Node {
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

// Peek returns the root node (the node with the largest weight) without removing it from the heap
func (h *MaxHeap) Peek() Node {
	if h.Size() == 0 {
		return Node{}
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
		if h.nodes[index].Weight <= h.nodes[parentIndex].Weight {
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
		if h.nodes[index].Weight >= h.nodes[largestChildIndex].Weight {
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
	if h.nodes[leftChildIndex].Weight > h.nodes[rightChildIndex].Weight {
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
