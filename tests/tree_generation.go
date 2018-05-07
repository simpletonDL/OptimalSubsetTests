package tests

import (
	. "github.com/simpletonDL/OptimalSubsetTests/trees"
	"time"
	"math/rand"
	"strconv"
)
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

/**
Генерит рандомные тест
 */
func generateNode(maxWeight int, maxProfit float64, id string) *Node {
	newNode := NewNode(1 + int64(random.Intn(maxWeight)), float64(random.Float64() * maxProfit), id)
	return newNode
}

/**
Генерит рандомное дерево из n вершин, и из тестов максимального
весв и максимального профита maxWeight и maxProfit соответсвенно.
 */
func GenerateTree(n int, maxWeight int, maxProfit float64) *Tree {
	genTree := &Tree{Root: generateNode(maxWeight, maxProfit, "0")}
	nodes := []*Node{genTree.Root}

	for i := 1; i < n; i++ {
		parentID := random.Intn(len(nodes))
		newNode := generateNode(maxWeight, maxProfit, strconv.Itoa(i))
		nodes[parentID].AddChild(newNode)
		/*if random.Intn(10) >= 9 {
			newNode.SetRequired()
		}*/
		nodes = append(nodes, newNode)
	}
	return genTree
}

//*/