package tests

import (
	. "github.com/simpletonDL/OptimalSubsetTests/trees"
	"math"
)

/**
	Брутфорсный алгоритм, принимает дерево и ограничение,
	возвращает ответ и массив тестов. Вообщем работает
	максимально глупо, за 2^n * n.
 */
func FindMinOptimalSubset(tree *Tree, bound int64) (float64, []*Node) {
	tree.UpdateSizes()
	treeSize := tree.Root.Size
	optimalProfit := 1.0
	var optimalSubset []*Node

	for i := 0; i < int(math.Pow(2, float64(treeSize))); i++ {
		var mask []int
		for j := uint(0); j < uint(treeSize); j++ {
			mask = append(mask, (i >> j) &1)
		}
		currentProfit, currentSubset := profitSubset(mask, tree, bound)
		if currentProfit > optimalProfit {
			optimalProfit = currentProfit
			optimalSubset = currentSubset
		}
	}
	return optimalProfit, optimalSubset
}


func profitSubset(mask []int, tree *Tree, bound int64) (float64, []*Node) {
	nodeMask := make(map[*Node]int)
	currentNodeID := 0

	var fillParentMask func(*Node)
	fillParentMask = func(node *Node) {
		for _, child := range node.Children {
			fillParentMask(child)
		}
		nodeMask[node] = mask[currentNodeID]
		currentNodeID++
	}
	fillParentMask(tree.Root)

	for node := range nodeMask {
		if node.Parent != nil {
			if nodeMask[node] == 1 && nodeMask[node.Parent] == 0 {
				return 1.0, nil
			}
		}
	}

	weight := int64(0)
	profit := 0.0
	var subset []*Node
	for node := range nodeMask {
		if nodeMask[node] == 1 {
			weight += node.Weight
			profit += node.Profit
			subset = append(subset, node)
		}
	}

	if weight <= bound {
		return profit, subset
	} else {
		return 1, nil
	}
}