package weightedmaxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap_Pop(t *testing.T) {
	for _, test := range []struct {
		name          string
		nodes         []Node
		expectedPop   Node
		expectedNodes []Node
	}{
		{
			name:          "when the heap is empty pop should return the zero value",
			nodes:         []Node{},
			expectedPop:   Node{},
			expectedNodes: []Node{},
		},
		{
			name: "when the heap has only one node pop should return the node",
			nodes: []Node{
				{Value: 1, Weight: 1},
			},
			expectedPop:   Node{Value: 1, Weight: 1},
			expectedNodes: []Node{},
		},
		{
			name: "should pop the nodes in the correct order when the nodes are in ascending order",
			nodes: []Node{
				{Value: 3, Weight: 3},
				{Value: 2, Weight: 2},
				{Value: 1, Weight: 1},
			},
			expectedPop: Node{Value: 3, Weight: 3},
			expectedNodes: []Node{
				{Value: 2, Weight: 2},
				{Value: 1, Weight: 1},
			},
		},
		{
			name: "should pop the nodes in the correct order when the nodes are in random order",
			nodes: []Node{
				{Value: 4, Weight: 4},
				{Value: 3, Weight: 3},
				{Value: 2, Weight: 2},
				{Value: 1, Weight: 1},
			},
			expectedPop: Node{Value: 4, Weight: 4},
			expectedNodes: []Node{
				{Value: 3, Weight: 3},
				{Value: 1, Weight: 1},
				{Value: 2, Weight: 2},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			heap := &MaxHeap{nodes: test.nodes}
			assert.Equal(t, heap.Pop(), test.expectedPop, "the pop should be the correct node")
			assert.Equal(t, heap.nodes, test.expectedNodes, "the nodes should be in the correct order")
		})
	}
}

func TestMaxHeap_Push(t *testing.T) {
	for _, test := range []struct {
		name         string
		nodes        []Node
		expectedPeek Node
	}{
		{
			name: "should push the nodes in the correct order when the nodes are in ascending order",
			nodes: []Node{
				{Value: 1, Weight: 1},
				{Value: 2, Weight: 2},
				{Value: 3, Weight: 3},
			},
			expectedPeek: Node{Value: 3, Weight: 3},
		},
		{
			name: "should push the nodes in the correct order when the nodes are in descending order",
			nodes: []Node{
				{Value: 3, Weight: 3},
				{Value: 2, Weight: 2},
				{Value: 1, Weight: 1},
			},
			expectedPeek: Node{Value: 3, Weight: 3},
		},
		{
			name: "should push the nodes in the correct order when the nodes are in random order",
			nodes: []Node{
				{Value: 2, Weight: 2},
				{Value: 3, Weight: 3},
				{Value: 1, Weight: 1},
				{Value: 4, Weight: 4},
			},
			expectedPeek: Node{Value: 4, Weight: 4},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			heap := &MaxHeap{}
			for _, node := range test.nodes {
				heap.Push(node)
			}
			assert.Equal(t, heap.Peek(), test.expectedPeek)
		})
	}
}
