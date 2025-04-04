// Package graph provides algorithms on graph.
package graph

import (
	"github.com/RVetas/Go-Algorithms/structures/node"
)

// Walk recursively walks the graph calling fn for each node, including starting position.
func Walk[T any](head *node.Node[T], fn func(node *node.Node[T])) {
	if head == nil {
		return
	}

	// To prevent cycling we keep visitedNodes
	visitedNodes := make(map[*node.Node[T]]bool)
	walk(head, visitedNodes, fn)
}

func walk[T any](node *node.Node[T], visitedNodes map[*node.Node[T]]bool, fn func(node *node.Node[T])) {
	if val, ok := visitedNodes[node]; val && ok {
		return
	}
	visitedNodes[node] = true

	if fn != nil {
		fn(node)
	}
	for _, val := range node.Children {
		walk(val, visitedNodes, fn)
	}
}
