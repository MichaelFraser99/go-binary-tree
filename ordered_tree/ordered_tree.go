package ordered_tree

import (
	"errors"
	"github.com/MichaelFraser99/go-binary-tree/internal"
)

type OrderedNode struct {
	Contents C            // the value of the node
	Left     *OrderedNode // lesser node
	Right    *OrderedNode // greater node
}

func New(val any, impl C) *OrderedNode {
	newImpl := impl.New()
	newImpl.GetContents().SetValue(val)
	newImpl.GetContents().SetCount(1)
	return &OrderedNode{
		Contents: newImpl,
	}
}

type C interface {
	New() C // Creates a new instance of the defined implementation
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

func (n *OrderedNode) Find(value any) any {
	targetNode := n

	for true {
		if targetNode.Contents.Equals(value) {
			return targetNode.Value()
		}

		if targetNode.Contents.Compare(value) {
			if targetNode.Right == nil {
				return nil
			}
			targetNode = targetNode.Right
		} else {
			if targetNode.Left == nil {
				return nil
			}
			targetNode = targetNode.Left
		}
	}
	return nil
}

// Remove - Removes a value from the tree
/*
If multiple of target value, the count is decremented by 1

If only one of a value, the node is removed

If the removal results in an empty tree, an error is returned
*/
func (n *OrderedNode) Remove(value any) error {
	targetNode := n
	var parent *OrderedNode
	var tempHolder *OrderedNode

	for true {
		if targetNode.Contents.Equals(value) {
			if targetNode.Count() > 1 {
				targetNode.Contents.GetContents().SetCount(targetNode.Count() - 1)
				return nil
			} else {
				if parent == nil {
					if targetNode.Right == nil && targetNode.Left == nil {
						return errors.New("cannot remove last value in tree")
					}

					if targetNode.Right == nil {
						n.Contents.GetContents().SetValue(targetNode.Left.Value())
						n.Contents.GetContents().SetCount(targetNode.Left.Count())
						tempHolder = targetNode.Left
						n.SetLeft(tempHolder.Left)
						n.SetRight(tempHolder.Right)
						return nil
					} else {
						tempHolder = targetNode.Left

						if targetNode.Right.Left == nil {
							targetNode.Right.Left = tempHolder

							n.Contents.GetContents().SetValue(targetNode.Right.Value())
							n.Contents.GetContents().SetCount(targetNode.Right.Count())
							n.SetLeft(targetNode.Right.Left)
							n.SetRight(targetNode.Right.Right)
							return nil
						} else {
							nestedTarget := targetNode.Right.Left

							for true {
								if nestedTarget.Left == nil {
									nestedTarget.Left = tempHolder
									break
								} else {
									nestedTarget = nestedTarget.Left
								}
							}

							n.Contents.GetContents().SetValue(targetNode.Right.Value())
							n.Contents.GetContents().SetCount(targetNode.Right.Count())
							n.SetLeft(targetNode.Right.Left)
							n.SetRight(targetNode.Right.Right)
							return nil
						}
					}
				} else {
					if parent.Right != nil && parent.Right.Contents.Equals(targetNode.Value()) {
						if targetNode.Right == nil {
							parent.Right = targetNode.Left
						} else {
							tempHolder = targetNode.Left
							parent.Right = targetNode.Right

							if targetNode.Right.Left == nil {
								targetNode.Right.Left = tempHolder
							} else {
								nestedTarget := targetNode.Right.Left

								for true {
									if nestedTarget.Left == nil {
										nestedTarget.Left = tempHolder
										break
									} else {
										nestedTarget = nestedTarget.Left
									}
								}
							}
						}
					} else if parent.Left != nil && parent.Left.Contents.Equals(targetNode.Value()) {
						if targetNode.Left == nil {
							parent.Left = targetNode.Right
						} else {
							tempHolder = targetNode.Right
							parent.Left = targetNode.Left

							if targetNode.Left.Right == nil {
								targetNode.Left.Right = tempHolder
							} else {
								nestedTarget := targetNode.Left.Right

								for true {
									if nestedTarget.Right == nil {
										nestedTarget.Right = tempHolder
										break
									} else {
										nestedTarget = nestedTarget.Right
									}
								}
							}
						}
					}
				}
				return nil
			}
			return nil
		} else {
			if targetNode.Contents.Compare(value) {
				if targetNode.Right == nil {
					return nil
				}
				parent = targetNode
				targetNode = targetNode.Right
			} else {
				if targetNode.Left == nil {
					return nil
				}
				parent = targetNode
				targetNode = targetNode.Left
			}
		}
	}
	return nil
}
