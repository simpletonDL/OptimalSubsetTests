package tests

import (
	"testing"
	"fmt"
	"project/OptimalSubsetTests/knapsack"
	"math"
	"sort"
	"time"
)

func sum(args ...int) {
	for x := range(args) {
		fmt.Print(x)
	}
}

func TestMulti(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("Test:", i)
		if (!testSingle()) {
			t.Error("Somethin wrong")
		}
	}
}

func TestTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		tree := GenerateTree(10000, 1000, 1.0)
		timeNow := time.Now()
		knapsack.FindMinOptimalSubset(tree, 100000)
		time := time.Now().Sub(timeNow).Seconds()
		fmt.Println(time)
	}
}

func TestNewKnapsack(t *testing.T) {
	Max := -1;
	for i := 0; i <= 1000; i++ {
		tree := GenerateTree(20, 100, 1.0)
		optimalAnswer, _ := FindMinOptimalSubset(tree, 500)
		knapsackAnswer, MaxSize := knapsack.FindProbability(tree, 500)
		if Max < MaxSize {
			Max = MaxSize
		}
		fmt.Println(optimalAnswer, knapsackAnswer, MaxSize, Max, (optimalAnswer - knapsackAnswer) < 0.000000001)
	}
}

func TestSplit(t *testing.T) {
	N := 1000
	low := N /3
	up := (2 * N) / 3
	for i := 0; i < 200000; i++ {
		tree := GenerateTree(N, 10, 1.0)
		tree.UpdateSizes()
		tree1, tree2 , _:= knapsack.SplitTree(tree)
		tree1.UpdateSizes()
		tree2.UpdateSizes()
		s1 := tree1.GetSize()
		s2 := tree2.GetSize()
		if low <= s1 && s1 <= up || low <= s2 && s2 <= up + 1 {
			//fmt.Println(s1, s2, low, up)
		} else {
			t.Error(s1, s2, "is not in range [", low, ",", up,"]")
		}
	}
}

func testSingle() bool {
	countVertex := 20
	maxWeight := 30
	maxProfit := 1.0

	tree := GenerateTree(countVertex, maxWeight, maxProfit)
	for targetBound := 0; targetBound <= countVertex * maxWeight; targetBound+=10 {
		optimalAnswer, optimalSubset := knapsack.FindMinOptimalSubset(tree, targetBound)
		bruteforceAnswer, bruteforceSubset := FindMinOptimalSubset(tree, targetBound)

		optimalSubsetID := []int{}
		bruteforceSubsetID := []int{}
		for _, node := range(optimalSubset) {
			optimalSubsetID = append(optimalSubsetID, node.ID)
		}
		for _, node := range(bruteforceSubset) {
			bruteforceSubsetID = append(bruteforceSubsetID, node.ID)
		}
		sort.Slice(optimalSubsetID, func(i, j int) bool {
			return optimalSubsetID[i] < optimalSubsetID[j]
			})

		sort.Slice(bruteforceSubsetID, func(i, j int) bool {
			return bruteforceSubsetID[i] < bruteforceSubsetID[j]
			})

		status := ""
		if math.Abs(bruteforceAnswer-optimalAnswer) < 0.0000000001 {
			fmt.Println("Bound: ", targetBound,
				"Optimal subset:", optimalSubsetID,
				"Bruteforce subset:", bruteforceSubsetID,
				"Bruteforce answer:", bruteforceAnswer,
				"Optimal answer:", optimalAnswer)
		} else {
			fmt.Print("Wrong answer!!!",
				"Bound: ", targetBound,
				"Status:", status,
				"Bruteforce answer:", bruteforceAnswer,
				"Optimal answer:", optimalAnswer,
				"Optimal subset:", optimalSubsetID,
				"Bruteforce subset:", bruteforceSubsetID)
			return false
		}

	}

	return true
}