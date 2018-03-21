package knapsack

import (
	. "project/OptimalSubsetTests/tries"
	"math"
)

/*	Возвращает массив [0..bound] ответов по
	максимизации сложения с учетом необходимых вершин
 */
func FindOptimalProbability(tree Tree, bound int) ([]float64) {
	n := tree.UpdateSizes() // Можно с оптиммизировать, но все равно залазит в O(nW)

	dp := make(map[int] []float64)
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
				delete(dp, leftBrother)
				//fmt.Println("Delete", leftBrother)
			}
			leftBrother = current
		}

		// После всех детей мы считаем dp
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
				dp[current][w] = dp[current-node.Size][w] // не берём
				if w - node.Weight >= 0 { // можем взять
					dp[current][w] = math.Max(dp[current][w], dp[current-1][w-node.Weight] + node.Profit) // Максимизируем сложение
				}
			}
		}

		if len(node.Children) != 0 {
			delete(dp, current - 1)
			//fmt.Println("Delete", current - 1)
		}
		//fmt.Println(dp, "afrter current", current)
	}
	dfs(tree.Root)
	//dp[0][i] = 0 for any i, so it is not correct
	return dp[n]
}

func FindOptimalSubset(tree Tree, W int) (float64, []*Node) {
	tree.UpdateSizes()

	if tree.GetSize() == 1 {
		root := tree.Root
		if W - root.Weight >= 0 && root.ID != -1 {
			return root.Profit, []*Node{root}
		} else {
			return 0, []*Node{}
		}
	}

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
			/*copyTreeUp.Print()
			fmt.Println()
			fmt.Println(copyNodeSplit)*/
			parent := copyNodeSplit.Parent
			parent.Children = parent.Children[:len(parent.Children)-1] // Удаляем все поддерево copyNodeSplit
			ansWithoutSplit = FindOptimalProbability(copyTreeUp, W)[W]
		}

		if ansWithSplit > ansWithoutSplit {
			//fmt.Print("case: 1")
			ansUp, setUp := FindOptimalSubset(treeUp, upW)
			ansDown, setDown := FindOptimalSubset(treeDown, downW)
			return  ansUp + ansDown, append(setUp, setDown...)
		} else {
			//fmt.Print("case: 2")
			return FindOptimalSubset(copyTreeUp, W)
		}
	} else {
		//fmt.Print("case: 3")
		ansUp, setUp := FindOptimalSubset(treeUp, upW)
		ansDown, setDown := FindOptimalSubset(treeDown, downW)
		return  ansUp + ansDown, append(setUp, setDown...)
	}

	return 0.0, []*Node{}
}
