// Package graph provides algorithms on graph.
package graph

import (
	"github.com/RVetas/Go-Algorithms/structures/node"
)

// WalkDFS recursively walks the graph depth-first calling fn for each node, including starting position.
func WalkDFS[T any](head *node.Node[T], fn func(node *node.Node[T])) {
	if head == nil {
		return
	}

	// To prevent cycling we keep visitedNodes
	visitedNodes := make(map[*node.Node[T]]bool)
	walkDFS(head, visitedNodes, fn)
}

// WalkBFS iteratively walks the graph breadth-first calling fn for each node, including starting position.
func WalkBFS[T any](head *node.Node[T], fn func(node *node.Node[T])) {
	if head == nil {
		return
	}

	// To prevent cycling we keep visitedNodes
	visitedNodes := make(map[*node.Node[T]]bool)
	walkBFS(head, visitedNodes, fn)
}

func walkDFS[T any](node *node.Node[T], visitedNodes map[*node.Node[T]]bool, fn func(node *node.Node[T])) {
	if val, ok := visitedNodes[node]; val && ok {
		return
	}
	visitedNodes[node] = true

	if fn != nil {
		fn(node)
	}
	for _, val := range node.Children {
		walkDFS(val, visitedNodes, fn)
	}
}
func walkBFS[T any](head *node.Node[T], visitedNodes map[*node.Node[T]]bool, fn func(node *node.Node[T])) {
	queue := make([]*node.Node[T], 0, 5)
	queue = append(queue, head)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if val, ok := visitedNodes[current]; val && ok {
			continue
		}
		if fn != nil {
			fn(current)
		}
		queue = append(queue, current.Children...)
	}
}
