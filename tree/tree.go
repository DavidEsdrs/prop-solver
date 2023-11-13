package tree

import "fmt"

type Tree[T any] struct {
	mapping map[string][]bool
	root    *Node[T]
}

func NewTree[T any](root *Node[T], mapping map[string][]bool) Tree[T] {
	return Tree[T]{
		root:    root,
		mapping: mapping,
	}
}

func (t *Tree[T]) Root() *Node[T] {
	return t.root
}

func (t *Tree[T]) Mapping() map[string][]bool {
	return t.mapping
}

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (n *Node[T]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (t *Tree[T]) BreadthFirstTraversal() []T {
	var result []T
	if t.root == nil {
		return result
	}

	queue := []*Node[T]{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func (t *Tree[T]) InOrderTraversal() {
	inOrder(t.root)
}

func inOrder[T any](node *Node[T]) {
	if node == nil {
		return
	}

	inOrder(node.Left)
	fmt.Printf("%#v ", node.Value)
	inOrder(node.Right)
}
