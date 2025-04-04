package graph

import (
	"slices"
	"testing"

	"github.com/RVetas/Go-Algorithms/structures/node"
)

func TestWalk(t *testing.T) {
	// given
	graph := sampleGraph()
	expectedResult := []int{1, 2, 3}
	result := []int{}

	// when
	Walk(graph, func(n *node.Node[int]) { result = append(result, n.Val) })

	// then
	if !slices.Equal(result, expectedResult) {
		t.Errorf("expected: %v, got: %v", expectedResult, result)
	}

}

// Simple graph with a cycle: n1 <-> n2 -> n3
func sampleGraph() *node.Node[int] {
	n1 := &node.Node[int]{Val: 1, Parent: nil, Children: []*node.Node[int]{}}
	n2 := &node.Node[int]{Val: 2, Parent: n1, Children: []*node.Node[int]{n1}}
	n3 := &node.Node[int]{Val: 3, Parent: n2, Children: []*node.Node[int]{}}
	n1.Children = append(n1.Children, n2)
	n2.Children = append(n2.Children, n3)

	return n1
}
