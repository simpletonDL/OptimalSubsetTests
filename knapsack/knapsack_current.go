package knapsack

import (
	. "project/OptimalSubsetTests/tries"
	"math"
)

//Возвращает массив [0..bound] ответов
func FindProbability(tree Tree, bound int) (float64, int) {
	maxSize := 0
	n := tree.UpdateSizes() // Можно с оптиммизировать, но все равно залазит в O(nW)

	dp := make(map[int] []float64)
	dp[0] = make([]float64, bound + 1) // [0..bound], initial [0..0]
	for w := range dp[0] {
		dp[0][w] = 1
	}

	current := 0
	var dfs func(*Node)
	dfs = func(node *Node) {
		leftBrother := 0
		for i, child := range node.Children {
			dfs(child)
			if i != 0 {
				delete(dp, leftBrother)
				//fmt.Println("Delete", leftBrother)
			}
			leftBrother = current
		}

		// После всех детей мы считаем dp
		current++
		dp[current] = make([]float64, bound + 1)

		if maxSize < len(dp) {
			maxSize = len(dp)
		}

		for w := 0; w <= bound; w++ {
			dp[current][w] = dp[current-node.Size][w] // не берём
			if w-node.Weight >= 0 { // можем взять
				dp[current][w] = math.Min(dp[current][w], dp[current-1][w-node.Weight] * node.Profit) // Минимизируем вероятность
			}
		}

		if len(node.Children) != 0 {
			delete(dp, current - 1)
			//fmt.Println("Delete", current - 1)
		}
		//fmt.Println(dp, "afrter current", current)
	}
	dfs(tree.Root)
	return dp[n][bound], maxSize
}