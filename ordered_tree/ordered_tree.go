package ordered_tree

import (
	"github.com/MichaelFraser99/go-binary-tree/internal"
)

type OrderedNode struct {
	Contents C            // the value of the node
	Left     *OrderedNode // lesser node
	Right    *OrderedNode // greater node
}

func New(val any, impl C) *OrderedNode {
	newImpl := impl.Clone()
	newImpl.GetContents().SetValue(val)
	newImpl.GetContents().SetCount(1)
	return &OrderedNode{
		Contents: newImpl,
	}
}

type C interface {
	Clone() C
	GetContents() *Contents
	Compare(newValue any) bool // Should return true for greater than, and false for less than
	Equals(value any) bool
}

type Contents struct {
	Value any
	Count int
}

func (c *Contents) GetContents() *Contents {
	return c
}

func (c *Contents) IncrementCount() {
	c.Count += 1
}

func (c *Contents) SetValue(val any) {
	c.Value = val
}

func (c *Contents) SetCount(val int) {
	c.Count = val
}

func (n *OrderedNode) Value() any {
	return n.Contents.GetContents().Value
}

func (n *OrderedNode) Count() int {
	return n.Contents.GetContents().Count
}

func (n *OrderedNode) LeftNode() *internal.Node {
	if n.Left == nil {
		return nil
	} else {
		return internal.IndirectNode(n.Left)
	}
}

func (n *OrderedNode) SetLeft(node *OrderedNode) {
	n.Left = node
}

func (n *OrderedNode) RightNode() *internal.Node {
	if n.Right == nil {
		return nil
	} else {
		return internal.IndirectNode(n.Right)
	}
}

func (n *OrderedNode) SetRight(node *OrderedNode) {
	n.Right = node
}

func (n *OrderedNode) Add(value any) error {
	// Check type of value added matches that of the existing node
	if err := internal.CheckType(n.Contents.GetContents().Value, value); err != nil {
		return err
	}

	targetNode := n
	contentsImplementation := n.Contents
	run := true

	for run {

		if targetNode.Contents.Equals(value) {
			targetNode.Contents.GetContents().IncrementCount()
			run = false
		} else if targetNode.Contents.Compare(value) { // if value is greater than
			if targetNode.Right == nil {
				targetNode.SetRight(New(value, contentsImplementation))
				run = false
			}
			targetNode = targetNode.Right
		} else {
			if targetNode.Left == nil {
				targetNode.SetLeft(New(value, contentsImplementation))
				run = false
			}
			targetNode = targetNode.Left
		}
	}

	return nil
}

func (n *OrderedNode) AscList() []any {
	var list []any
	return internal.ListNodesAsc(n, list)
}

func (n *OrderedNode) DescList() []any {
	var list []any
	return internal.ListNodesDesc(n, list)
}
