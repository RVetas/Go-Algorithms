package graph

import (
	"fmt"
	"slices"
	"testing"

	"github.com/RVetas/Go-Algorithms/structures/node"
)

func TestWalkDFS(t *testing.T) {
	// given
	graph := sampleGraph()
	expectedResult := []string{"0", "0_0", "0_0_0", "0_0_1", "0_1", "0_1_0", "0_1_1", "0_2", "0_2_0", "0_2_1"}
	result := []string{}

	// when
	WalkDFS(graph, func(n *node.Node[string]) { result = append(result, n.Val) })

	// then
	if !slices.Equal(result, expectedResult) {

		t.Errorf("expected: %v, got: %s", expectedResult, result)
	}
}

func TestWalkBFS(t *testing.T) {
	// given
	graph := sampleGraph()
	expectedResult := []string{"0", "0_0", "0_1", "0_2", "0_0_0", "0_0_1", "0_1_0", "0_1_1", "0_2_0", "0_2_1"}
	result := []string{}
	// when
	WalkBFS(graph, func(n *node.Node[string]) { result = append(result, n.Val) })

	// then
	if !slices.Equal(result, expectedResult) {
		t.Errorf("expected: %v, got: %v", expectedResult, result)
	}
}

// Generates sample graph. Head has N (=3) children, every child has 2 children
func sampleGraph() *node.Node[string] {
	head := &node.Node[string]{
		Val:      "0",
		Parent:   nil,
		Children: []*node.Node[string]{},
	}
	head.Children = generateLayer(3, head)

	for _, child := range head.Children {
		child.Children = generateLayer(2, child)
	}

	return head
}

func generateLayer(childrenCount int, parent *node.Node[string]) []*node.Node[string] {
	result := make([]*node.Node[string], childrenCount)
	for i := 0; i < childrenCount; i++ {
		val := fmt.Sprintf("%s_%d", parent.Val, i)
		node := &node.Node[string]{
			Val:      val,
			Parent:   parent,
			Children: []*node.Node[string]{},
		}
		result[i] = node
	}
	return result
}
