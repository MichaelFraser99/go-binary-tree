package internal

import (
	"errors"
	"fmt"
	"reflect"
)

func CheckType(existingValue, newValue any) error {
	if reflect.TypeOf(existingValue) != reflect.TypeOf(newValue) {
		return errors.New(fmt.Sprintf("cannot add miss-matching types to ordered tree - expected %s, got %s", reflect.TypeOf(existingValue), reflect.TypeOf(newValue)))
	}
	return nil
}

func ListNodesAsc(node Node, list []any) []any {
	if node.LeftNode() != nil {
		list = append(ListNodesAsc(*node.LeftNode(), list))
	}

	for i := 0; i < node.Count(); i++ {
		list = append(list, node.Value())
	}

	if node.RightNode() != nil {
		list = append(ListNodesAsc(*node.RightNode(), list))
	}
	return list
}

func ListNodesDesc(node Node, list []any) []any {
	if node.RightNode() != nil {
		list = append(ListNodesDesc(*node.RightNode(), list))
	}

	for i := 0; i < node.Count(); i++ {
		list = append(list, node.Value())
	}

	if node.LeftNode() != nil {
		list = append(ListNodesDesc(*node.LeftNode(), list))
	}
	return list
}

func IndirectNode(v Node) *Node {
	return &v
}
