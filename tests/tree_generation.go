package tests

import (
	. "project/OptimalSubsetTests/tries"
	"time"
	"math/rand"
)
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateNode(maxWeight int, maxProfit float64, id int) *Node {
	newNode := NewNode(1 + random.Intn(maxWeight), float64(random.Float64() * maxProfit), id)
	return newNode
}

func GenerateTree(n int, maxWeight int, maxProfit float64) Tree {
	genTree := Tree{Root: generateNode(maxWeight, maxProfit, 0)}
	nodes := []*Node{genTree.Root}

	for i := 1; i < n; i++ {
		parentID := random.Intn(len(nodes))
		newNode := generateNode(maxWeight, maxProfit, i)
		nodes[parentID].AddChild(newNode)
		if random.Intn(10) >= 9 {
			newNode.SetRequired()
		}
		nodes = append(nodes, newNode)
	}
	return genTree
}

//*/