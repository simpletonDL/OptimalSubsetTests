package tries

import (
	"fmt"
)

type Tree struct {
	Root *Node
}

func (tree Tree) UpdateSizes() int {
	return tree.Root.UpdateSizes()
}

func (tree Tree) Print() {
	tree.Root.Print()
}

type Node struct {
	Weight int
	Profit float64
	Size   int

	ID int
	Parent *Node
	Children []*Node
}

func NewNode(weight int, profit float64, id int) *Node {
	return &Node{weight, profit, 1, id, nil, []*Node{}}
}

func (node *Node) AddChild(child *Node) {
	node.Children = append(node.Children, child)
	child.Parent = node
}

func (node *Node) UpdateSizes() (int) {
	node.Size = 1
	for _, ptr := range node.Children {
		node.Size += ptr.UpdateSizes()
	}
	return node.Size
}

func (node Node) Print() {
	fmt.Print("[", node.Weight, " ")
	for _, ptr := range node.Children {
		ptr.Print()
	}
	fmt.Print(" ", node.Profit, "]")
}