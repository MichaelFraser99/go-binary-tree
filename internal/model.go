package internal

type Node interface {
	Value() any
	Count() int
	LeftNode() *Node
	RightNode() *Node
}

type DecisionNode struct {
	Question  string        // the question to ask at this node
	LeftNode  *DecisionNode // the node to go to if the answer is yes
	RightNode *DecisionNode // the node to go to if the answer is no
}
