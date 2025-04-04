// Package node provides Node structure that represents a node with a value, a parent and children.
// It can be used in representation of such structures as graphs or trees.
package node

// Node structure holds a value, a pointer on parent Node and pointers to its children.
// It can be used in representation of graphs or trees.
type Node[T any] struct {
	Val      T
	Parent   *Node[T]
	Children []*Node[T]
}
