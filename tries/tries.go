package tries

import (
	"fmt"
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

func (tree Tree) GetSize() int {
	return tree.Root.Size
}

func (tree Tree) FindById(ID int) *Node {
	var target *Node
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node.ID == ID {
			target = node
			return
		} else {
			for _, child := range node.Children {
				dfs(child)
			}
		}
	}
	dfs(tree.Root)
	return target
}

func (tree Tree) Copy() Tree {
	copyTree := Tree{Root:tree.Root.Copy()}

	var dfs func(*Node, *Node)
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
	return copyTree
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

	ID         int
	IsRequired bool
	Parent     *Node
	Children   []*Node
}

func NewNode(weight int, profit float64, id int) *Node {
	return &Node{weight, profit, 1, id, false, nil, []*Node{}}
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
	var dfs func(*Node)
	dfs = func (ptr *Node) {
		ptr.IsRequired = true
		if ptr.Parent != nil {
			dfs(ptr.Parent)
		}
	}
	dfs(node)
}

func (node *Node) Print() {
	fmt.Print("[", node, " ")
	for _, ptr := range node.Children {
		ptr.Print()
	}
	fmt.Print(" ", "]")
}

func (node *Node) Copy() *Node {
	return &Node{node.Weight, node.Profit, node.Size, node.ID, node.IsRequired, node.Parent, []*Node{}}
}

func (node *Node) IsRoot() bool {
	return node.Parent == nil
}