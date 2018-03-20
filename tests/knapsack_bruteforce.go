package tests

import (
	. "project/OptimalSubsetTests/tries"
	"math"
)

func FindMinOptimalSubset(tree Tree, bound int) (float64, []*Node) {
	tree.UpdateSizes()
	treeSize := tree.Root.Size
	optimalProfit := 1.0
	optimalSubset := []*Node{}

	for i := 0; i < int(math.Pow(2, float64(treeSize))); i++ {
		var mask []int
		for j := uint(0); j < uint(treeSize); j++ {
			mask = append(mask, (i >> j) &1)
		}
		currentProfit, currentSubset := profitSubset(mask, tree, bound)
		if (currentProfit < optimalProfit) {
			optimalProfit = currentProfit
			optimalSubset = currentSubset
		}
	}
	return optimalProfit, optimalSubset
}


func profitSubset(mask []int, tree Tree, bound int) (float64, []*Node) {
	nodeMask := make(map[*Node]int)
	currentNodeID := 0

	var fillParentMask func(*Node)
	fillParentMask = func(node *Node) {
		for _, child := range(node.Children) {
			fillParentMask(child)
		}
		nodeMask[node] = mask[currentNodeID]
		currentNodeID++
	}
	fillParentMask(tree.Root)

	for node, _ := range(nodeMask) {
		if node.Parent != nil {
			if nodeMask[node] == 1 && nodeMask[node.Parent] == 0 {
				return 1.0, nil
			}
		}
	}

	weight := 0
	profit := 1.0
	subset := []*Node{}
	for node, _ := range(nodeMask) {
		if (nodeMask[node] == 1) {
			weight += node.Weight
			profit *= node.Profit
			subset = append(subset, node)
		}
	}

	if (weight <= bound) {
		return profit, subset
	} else {
		return 1, nil
	}
}