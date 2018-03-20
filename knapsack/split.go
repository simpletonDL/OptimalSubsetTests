package knapsack

import (
	. "project/OptimalSubsetTests/tries"
	"sort"
)

type bySubtreeSize []*Node

func (s bySubtreeSize) Len() int {
	return len(s)
}
func (s bySubtreeSize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s bySubtreeSize) Less(i, j int) bool {
	return s[i].Size < s[j].Size
}

/*
Требует UpdateSize(), и смысла сплитить дерево из одной вершины нет
 */
func SplitTree(tree Tree) (Tree, Tree, *Node) {
	// Проверить случай с 0 и 1 вершиной
	n := tree.GetSize()
	if n < 2 {
		return tree, Tree{}, tree.Root
	}

	SortChildren(tree)
	low := n / 3
	up := (2 * n) / 3
	ptr := tree.Root

	for len(ptr.Children) != 0 && ptr.Children[len(ptr.Children) - 1].Size > up {
		//fmt.Println(ptr.Children[len(ptr.Children) - 1].Size)
		ptr = ptr.Children[len(ptr.Children) - 1]
	}
	// Первый раз, когда правый ребёнок <= up или нет детей
	size := 0
	for i := len(ptr.Children) - 1; i >= 0; i-- {
		size += ptr.Children[i].Size
		if size >= low {
			var newRoot *Node
			if i == len(ptr.Children) - 1 { // Если была только одна итерация
				newRoot = ptr.Children[i]
			} else {
				newRoot = NewNode(0, 1, -1) // profit = 1 для умножкния, для сложение нужен 0
				newRoot.Children = make([]*Node, len(ptr.Children)-i)
				copy(newRoot.Children, ptr.Children[i:])
			}
			ptr.Children = ptr.Children[:i]

			return tree, Tree{Root:newRoot}, ptr
		}
	}
	//Вот здесь бы exeption кидать
	return Tree{}, Tree{}, nil
}

func SortChildren(tree Tree)  {
	var dfs func(*Node)
	dfs = func (node *Node) {
		for _, child := range node.Children {
			dfs(child)
		}
		sort.Sort(bySubtreeSize(node.Children))
	}
	dfs(tree.Root)
}