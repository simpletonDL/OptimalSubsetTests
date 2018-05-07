package trees

import (
	"fmt"
	"time"
)

/**
Тут описаны структуры для работы с деревьями,
собственно само дерево Tree и его узлы Node,
а так же всякие нужные функции, имя которых говорит
само за себя. Единственная по первости непонятная
функция, это SetRequired. Она описана ниже.
Ну и забавная функция Print, выводящее дерево
в виде скобочной последовательности.
 */

type Tree struct {
	Root *Node
}

func (tree *Tree) GetSize() int {
	return tree.Root.Size
}

func (tree *Tree) FindById(ID string) *Node {
	return findByIdInternal(tree.Root, ID)
}

func findByIdInternal(node *Node, ID string) *Node {
	if node.ID == ID {
		return node
	}

	for _, child := range node.Children {
		if result := findByIdInternal(child, ID); result != nil {
			return result
		}
	}

	return nil
}

func (tree *Tree) Copy() *Tree {
	return &Tree{Root:copyInternal(tree.Root)}

	/*var dfs func(*Node, *Node)
	dfs = func (node *Node, copyNode *Node) {
		copyChildren := make([]*Node, 0)
		for _, child := range node.Children {
			copyChild := child.Copy()
			copyChild.Parent = copyNode
			copyChildren = append(copyChildren, copyChild)
			dfs(child, copyChild)
		}
		copyNode.Children = copyChildren
	}

	dfs(tree.Root, copyTree.Root)
	return copyTree*/
}

func copyInternal(node *Node) *Node {
	nodeCopy := node.Copy()
	for _, child := range node.Children {
		nodeCopy.AddChild(copyInternal(child))
	}
	return nodeCopy
}

func (tree *Tree) UpdateSizes() int {
	return tree.Root.UpdateSizes()
}

func (tree *Tree) Print() {
	tree.Root.Print()
}

type Node struct {
	Weight       int64
	Profit       float64

	ID           string
	TimeDuration time.Duration

	Size         int
	IsRequired   bool
	Parent       *Node
	Children     []*Node
}

func NewNode(timeByStep int64, profit float64, id string) *Node {
	return &Node{
		Weight:   timeByStep,
		Profit:   profit,
		ID:       id,
		Size:     1,
		// Children: []*Node{},
	}
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

/**
Устанавливает значение IsRequired true,
это означает, что этот node должен быть
взят в его дереве (а значит и все его
предки).
 */
func (node *Node) SetRequired()  {
	var dfs func(*Node) // todo сделать for
	dfs = func (ptr *Node) {
		ptr.IsRequired = true
		if ptr.Parent != nil {
			dfs(ptr.Parent)
		}
	}
	dfs(node)
}

func (node *Node) Print() {
	fmt.Print("[", node.ID, " ", node.Weight, " ")
	for _, ptr := range node.Children {
		ptr.Print()
	}
	fmt.Print(" ", "]")
}

func (node *Node) Copy() *Node {
	return &Node{node.Weight, node.Profit, node.ID, node.TimeDuration, node.Size, node.IsRequired, node.Parent, []*Node{}} // todo разбить
}

func (node *Node) IsRoot() bool {
	return node.Parent == nil
}

func NodeToID(nodes []*Node) []string { // todo return map[string] bool
	var xs []string
	for _, node := range nodes {
		xs = append(xs, node.ID)
	}
	return xs
}