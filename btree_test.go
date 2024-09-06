package gollections

import (
	"cmp"
	"fmt"
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	bt := NewBinaryTree[int]()

	bt.Insert(99)

	if v, b := bt.head.Get(); v != 99 || !b {
		t.Errorf("head should be %v, want %v", 99, v)
	}

	bt.Insert(77)
	if v, b := bt.head.prev.Get(); v != 77 || !b {
		t.Errorf("head left should be %d got %v", 77, v)
	}

	bt.Insert(101)
	if v, b := bt.head.next.Get(); v != 101 || !b {
		t.Errorf("head right should be %d got %v", 101, v)
	}

	bt.Insert(1)
	if v, b := bt.head.prev.prev.Get(); v != 1 || !b {
		t.Errorf("head left.left should be %d got %v", 1, v)
	}

}

func treeWithInput[T cmp.Ordered](input []T) *BinaryTree[T] {
	bt := NewBinaryTree[T]()

	for _, v := range input {
		bt.Insert(v)
	}

	return &bt
}

func TestPreorder(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "preorder-1", input: []int{99, 77, 33, 101, 90}, expected: []int{99, 77, 33, 90, 101}},
		{name: "preorder-2", input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{name: "preorder-3", input: []int{1, 99, 101, 1}, expected: []int{1, 99, 1, 101}},
		{name: "preorder-4", input: []int{5, 3, 7, 2, 4, 6, 8}, expected: []int{5, 3, 2, 4, 7, 6, 8}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bt := treeWithInput(test.input)
			fmt.Println("tree representation: ")
			fmt.Print(bt.String())

			var rs []int
			for _, v := range bt.Preorder {
				rs = append(rs, v)
			}

			fmt.Printf("travelsal order was:\n%v\n", rs)

			if !reflect.DeepEqual(rs, test.expected) {
				t.Errorf("got %v, want %v", rs, test.expected)
			}

		})
	}
}

func TestInorder(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "inorder-1", input: []int{99, 77, 33, 101, 90}, expected: []int{33, 77, 90, 99, 101}},
		{name: "inorder-2", input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{name: "inorder-3", input: []int{1, 99, 101, 1}, expected: []int{1, 1, 99, 101}},
		{name: "inorder-4", input: []int{5, 3, 7, 2, 4, 6, 8}, expected: []int{2, 3, 4, 5, 6, 7, 8}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bt := treeWithInput(test.input)

			fmt.Println("tree representation: ")
			fmt.Print(bt.String())

			var rs []int
			for _, v := range bt.Inorder {
				rs = append(rs, v)
			}

			fmt.Printf("travelsal order was:\n%v\n", rs)

			if !reflect.DeepEqual(rs, test.expected) {
				t.Errorf("got %v, want %v", rs, test.expected)
			}

		})
	}
}
