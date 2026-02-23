package weightedminheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap_Pop(t *testing.T) {
	for _, test := range []struct {
		name         string
		nodes        []Node
		expectedPop  Node
		expectedSize int
	}{
		{
			name:         "when the heap is empty pop should return the zero value",
			nodes:        []Node{},
			expectedPop:  Node{},
			expectedSize: 0,
		},
		{
			name: "when the heap has only one node pop should return the node",
			nodes: []Node{
				{Value: 1, Weight: 1},
			},
			expectedPop:  Node{Value: 1, Weight: 1},
			expectedSize: 0,
		},
		{
			name: "should pop the nodes in the correct order when the nodes are in ascending order",
			nodes: []Node{
				{Value: 1, Weight: 1},
				{Value: 2, Weight: 2},
				{Value: 3, Weight: 3},
			},
			expectedPop:  Node{Value: 1, Weight: 1},
			expectedSize: 2,
		},
		{
			name: "should pop the nodes in the correct order when the nodes are in descending order",
			nodes: []Node{
				{Value: 3, Weight: 3},
				{Value: 2, Weight: 2},
				{Value: 1, Weight: 1},
			},
			expectedPop:  Node{Value: 1, Weight: 1},
			expectedSize: 2,
		},
		{
			name: "should pop the nodes in the correct order when the nodes are in random order",
			nodes: []Node{
				{Value: 2, Weight: 2},
				{Value: 3, Weight: 3},
				{Value: 1, Weight: 1},
				{Value: 4, Weight: 4},
			},
			expectedPop:  Node{Value: 1, Weight: 1},
			expectedSize: 3,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			heap := &MinHeap{}
			for _, node := range test.nodes {
				heap.Push(node)
			}
			assert.Equal(t, test.expectedPop, heap.Pop(), "the pop should be the correct node")
			assert.Equal(t, test.expectedSize, heap.Size(), "the size should be correct")
		})
	}
}

func TestMinHeap_Push(t *testing.T) {
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
			expectedPeek: Node{Value: 1, Weight: 1},
		},
		{
			name: "should push the nodes in the correct order when the nodes are in descending order",
			nodes: []Node{
				{Value: 3, Weight: 3},
				{Value: 2, Weight: 2},
				{Value: 1, Weight: 1},
			},
			expectedPeek: Node{Value: 1, Weight: 1},
		},
		{
			name: "should push the nodes in the correct order when the nodes are in random order",
			nodes: []Node{
				{Value: 2, Weight: 2},
				{Value: 3, Weight: 3},
				{Value: 1, Weight: 1},
				{Value: 4, Weight: 4},
			},
			expectedPeek: Node{Value: 1, Weight: 1},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			heap := &MinHeap{}
			for _, node := range test.nodes {
				heap.Push(node)
			}
			assert.Equal(t, test.expectedPeek, heap.Peek())
		})
	}
}
