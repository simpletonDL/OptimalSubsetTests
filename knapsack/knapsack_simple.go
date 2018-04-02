package knapsack

import (
	. "project/OptimalSubsetTests/tries"
)

/**
	Алгоритм с квадратичной памятью, принмает дерево и ограничение,
	возвращает ответ и набор тестов. Не в коем случае не использовать
	при размере таблички >= 10^9, если, конечно, нет гигов 20 оператики.
 */
func SimpleKnapsack(tree Tree, bound int) (float64, []*Node) {
	n := tree.UpdateSizes()
	dp := make([][]float64, n+1)
	pred := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, bound+1)
		pred[i] = make([]bool, bound+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = 0
		}
	}

	i := 1
	nodesInOrder := []*Node{nil}
	var buildDP func(*Node)
	buildDP = func(node *Node) {
		for _, child := range node.Children {
			buildDP(child)
		}

		for w := 0; w <= bound; w++ {
			if w - node.Weight >= 0 && node.Profit + dp[i-1][w-node.Weight] > dp[i-node.Size][w] {
				pred[i][w] = true
				dp[i][w] = node.Profit + dp[i-1][w-node.Weight]
			} else {
				pred[i][w] = false
				dp[i][w] = dp[i-node.Size][w]
			}
		}

		nodesInOrder = append(nodesInOrder, node)
		i++
	}
	buildDP(tree.Root)

	optimalSubset := []*Node{}
	i = n
	w := bound
	for i > 0 {
		currentNode := nodesInOrder[i]
		if pred[i][w] == true {
			optimalSubset = append(optimalSubset, currentNode)
			w -= currentNode.Weight
			i--
		} else {
			i -= currentNode.Size
		}
	}

	return dp[n][bound], optimalSubset
}