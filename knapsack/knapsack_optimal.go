package knapsack

import (
	"time"
	"math"
)

import (
	. "project/OptimalSubsetTests/tries"
)

/**
Глобальные переменные для тестирования времени.
Имеют смысл при тестировании findOptimalAnswerAndSubsetHelper
 */
var FOP_time float64 = 0.0 // Полное время выполнения FindOptimalProbability
var FOP_deletes_time = 0.0 // Время выполнения удалений ненужной динамики в FindOptimalProbability
var FOP_count int = 0 // Суммарное колличество записей в динамику
var FOP_dp_time = 0.0 // Время выполнения динамики


/*	Возвращает массив [0..bound] ответов: для каждого
	ограничения по времени - свой ответ. Можно вообще
	использовать отдельно, если восстанвление
	ответа не нужно.
	Принимает дерево и ограничение, возвращает массив
	ответов.
*/

func FindOptimalProbability(tree Tree, bound int) ([]float64) {
	n := tree.UpdateSizes() // Можно с оптиммизировать, но все равно залазит в O(nW)

	FOP_count += tree.GetSize() * bound
	timeNow := time.Now()

	dp := make([][]float64, n + 1)
	dp[0] = make([]float64, bound + 1) // [0..bound], initial [0..0]
	for w := range dp[0] {
		dp[0][w] = 0
	}

	current := 0
	var dfs func(*Node)
	dfs = func(node *Node) {
		leftBrother := 0
		for i, child := range node.Children {
			dfs(child)
			if i != 0 {
				timeN := time.Now()
				dp[leftBrother] = nil
				FOP_deletes_time += time.Now().Sub(timeN).Seconds()
			}
			leftBrother = current
		}

		timeN := time.Now()
		current++
		dp[current] = make([]float64, bound + 1)

		for w := 0; w <= bound; w++ {
			if node.IsRequired {
				if w - node.Weight < 0 {
					dp[current][w] = math.Inf(-1)
				} else {
					dp[current][w] = dp[current-1][w-node.Weight] + node.Profit
				}
			} else {
				dp[current][w] = dp[current-node.Size][w]
				if w - node.Weight >= 0 {
					dp[current][w] = math.Max(dp[current][w], dp[current-1][w-node.Weight] + node.Profit) // Максимизируем сложение
				}
			}
		}
		FOP_dp_time += time.Now().Sub(timeN).Seconds()

		if len(node.Children) != 0 {
			timeN := time.Now()
			dp[current - 1] = nil
			FOP_deletes_time += time.Now().Sub(timeN).Seconds()
		}
	}
	dfs(tree.Root)

	FOP_time += time.Now().Sub(timeNow).Seconds()
	return dp[n]
}

/**
Собственно вот и герой нашего дня - алгоритм с высотно-линейной
памятью и восстановлением ответа. Создал helper, потому что helper
меняет дерево, так что сначала лучше бы его скопировать.
Принимает дерево и ограничение по времени, возвращает ответ и массив
тестов.
 */
func FindOptimalAnswerAndSubset(tree Tree, W int) (float64, []*Node) {
	copyTree := tree.Copy()
	return findOptimalAnswerAndSubsetHelper(copyTree, W)
}

/**
Краткое описание алгоритма: делим дерево, получаем вершину split,
которая родитель одной части. Рассматриваем два случая берем/не берем
split, в зависимости от этого запускаемся рекурсивно от двух или
одной части.
 */
func findOptimalAnswerAndSubsetHelper(tree Tree, W int) (float64, []*Node) {
	tree.UpdateSizes()

	// База
	if tree.GetSize() == 1 {
		root := tree.Root
		if W - root.Weight >= 0 && root.ID != -1 {
			return root.Profit, []*Node{root}
		} else {
			return 0, []*Node{}
		}
	}

	// Делим дерево
	treeUp, treeDown, nodeSplit := SplitTree(tree)
	dpDown := FindOptimalProbability(treeDown, W)

	copyTreeUp := treeUp.Copy()
	copyNodeSplit := copyTreeUp.FindById(nodeSplit.ID)

	//case 1: берём nodeSplit, nodeSplit -> treeUp
	if !nodeSplit.IsRequired {
		nodeSplit.SetRequired()
	}
	dpUp := FindOptimalProbability(treeUp, W)

	ansWithSplit, upW, downW := math.Inf(-1), -1, -1
	for i := 0; i <= W; i++ {
		if ansWithSplit < dpUp[i] + dpDown[W-i] {
			ansWithSplit, upW, downW = dpUp[i]+dpDown[W-i], i, W-i
		}
	}

	//case 2: не берём copyNodeSplit -> copyTreeUp
	if !copyNodeSplit.IsRequired {
		ansWithoutSplit := math.Inf(-1)
		if copyNodeSplit.IsRoot() {
			if copyNodeSplit.IsRequired {
				ansWithoutSplit = math.Inf(-1)
			} else {
				ansWithoutSplit = 0
			}
		} else {
			parent := copyNodeSplit.Parent
			parent.Children = parent.Children[:len(parent.Children)-1]
			ansWithoutSplit = FindOptimalProbability(copyTreeUp, W)[W]
		}

		if ansWithSplit > ansWithoutSplit {
			ansUp, setUp := findOptimalAnswerAndSubsetHelper(treeUp, upW)
			ansDown, setDown := findOptimalAnswerAndSubsetHelper(treeDown, downW)
			return  ansUp + ansDown, append(setUp, setDown...)
		} else {
			return findOptimalAnswerAndSubsetHelper(copyTreeUp, W)
		}
	} else {
		ansUp, setUp := findOptimalAnswerAndSubsetHelper(treeUp, upW)
		ansDown, setDown := findOptimalAnswerAndSubsetHelper(treeDown, downW)
		return  ansUp + ansDown, append(setUp, setDown...)
	}

	return 0.0, []*Node{} // go глупый и ругался на то, что нет return :D
}
