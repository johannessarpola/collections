package gollections

import (
	"cmp"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

type BinaryTree[T cmp.Ordered] struct {
	head *Node[T]
	mu   sync.Mutex
}

func NewBinaryTree[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (r *BinaryTree[T]) withLock(f func()) {
	defer r.mu.Unlock()
	r.mu.Lock()
	f()
}

func insert[T cmp.Ordered](root *Node[T], value T) *Node[T] {
	n := NewNode(value)
	if root == nil {
		return &n
	}
	v, _ := root.Get()
	if value < v {
		root.prev = insert(root.prev, value)
	} else {
		root.next = insert(root.next, value)
	}

	return root
}

func (b *BinaryTree[T]) Insert(values ...T) {
	b.withLock(func() {
		for _, v := range values {
			b.head = insert(b.head, v)
		}
	})
}

func preorder[T cmp.Ordered](root *Node[T], i *atomic.Int32, yield func(int, T) bool) {
	if root == nil {
		return
	}

	// Yield the current node
	index := int(i.Add(1) - 1)
	if !yield(index, root.inner) {
		return
	}

	// Traverse the left subtree
	preorder(root.prev, i, yield)

	// Traverse the right subtree
	preorder(root.next, i, yield)

}

func (l *BinaryTree[T]) Preorder(yield func(int, T) bool) {
	l.withLock(func() {
		i := atomic.Int32{}
		preorder(l.head, &i, yield)
	})
}

// String returns a string visualization of the binary tree.
func (bt *BinaryTree[T]) String() string {
	var sb strings.Builder
	bt.withLock(func() {
		if bt.head == nil {
			sb.WriteString("empty")
			return
		}

		bt.visualizeNode(bt.head, "", true, &sb)
	})

	return sb.String()
}

// visualizeNode helps in the recursive visualization of the binary tree.
func (bt *BinaryTree[T]) visualizeNode(node *Node[T], prefix string, isTail bool, sb *strings.Builder) {
	if node == nil {
		return
	}

	// Append current node value
	sb.WriteString(prefix)
	if isTail {
		sb.WriteString("└── ")
	} else {
		sb.WriteString("├── ")
	}
	sb.WriteString(fmt.Sprintf("%v\n", node.inner))

	// Prepare the prefix for child nodes
	childPrefix := prefix
	if isTail {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	// Recurse on children nodes
	hasPrev := node.prev != nil
	hasNext := node.next != nil

	if hasPrev {
		bt.visualizeNode(node.prev, childPrefix, !hasNext, sb)
	}

	if hasNext {
		bt.visualizeNode(node.next, childPrefix, true, sb)
	}
}
