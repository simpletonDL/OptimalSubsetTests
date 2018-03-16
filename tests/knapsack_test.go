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