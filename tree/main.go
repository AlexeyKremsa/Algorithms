package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	values := []int{5, 4, 3, 6, 1}

	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " | ") })
	fmt.Println()

	fmt.Println("Is balanced:", tree.IsBalacned())
}

type Tree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (t *Tree) IsBalacned() bool {
	if t.Root == nil {
		return true
	}

	if t.Root.balancedHeight() > -1 {
		return true
	}

	return false
}

func (n *Node) balancedHeight() int {
	if n == nil {
		return 0
	}

	hLeft := n.Left.balancedHeight()
	hRight := n.Right.balancedHeight()

	if hLeft == -1 || hRight == -1 {
		return -1
	}

	if math.Abs(float64(hLeft-hRight)) > 1 {
		return -1
	}

	if hLeft > hRight {
		return hLeft + 1
	}

	return hRight + 1
}

func (t *Tree) Insert(value int) error {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return nil
	}

	return t.Root.Insert(value)
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func (t *Tree) Find(v int) (int, bool) {
	if t.Root == nil {
		return 0, false
	}
	return t.Root.Find(v)
}

func (t *Tree) Delete(v int) error {
	if t.Root == nil {
		return errors.New("Tree delete: nil")
	}

	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(v, fakeParent)
	if err != nil {
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

func (n *Node) Insert(value int) error {
	if n == nil {
		return errors.New("Can not insert into nil tree")
	}

	switch {
	case n.Value == value:
		return nil

	case n.Value < value:
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return nil
		}
		return n.Right.Insert(value)

	case n.Value > value:
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return nil
		}
		return n.Left.Insert(value)
	}

	return nil
}

func (n *Node) Find(v int) (int, bool) {
	if n == nil {
		return 0, false
	}

	switch {
	case v == n.Value:
		return n.Value, true
	case v < n.Value:
		return n.Left.Find(v)
	case v > n.Value:
		return n.Right.Find(v)
	}

	return 0, false
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}

	if n.Right == nil {
		return n, parent
	}

	return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode: nil")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}

	parent.Right = replacement
	return nil
}

func (n *Node) Delete(v int, parent *Node) error {
	if n == nil {
		return errors.New("Delete: nil")
	}

	switch {
	case v < n.Value:
		return n.Left.Delete(v, n)
	case v > n.Value:
		return n.Right.Delete(v, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		replacement, replParent := n.Left.findMax(n)
		n.Value = replacement.Value

		return replacement.Delete(replacement.Value, replParent)
	}
}
