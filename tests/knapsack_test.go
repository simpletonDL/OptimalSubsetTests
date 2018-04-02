package tests

import (
	"testing"
	"fmt"
	"project/OptimalSubsetTests/knapsack"
	"math"
	"sort"
	"time"
	"project/OptimalSubsetTests/tries"
)

func TestNewKnapsack(t *testing.T) {
	Max := -1;
	for i := 0; i <= 1000; i++ {
		tree := GenerateTree(20, 100, 1.0)
		optimalAnswer, _ := FindMinOptimalSubset(tree, 500)
		knapsackAnswer := knapsack.FindOptimalProbability(tree, 500)[500]
		fmt.Println(optimalAnswer, knapsackAnswer, Max, (optimalAnswer - knapsackAnswer) < 0.000000001)
	}
}

/**
Тест деления дерево в отношении 1 к 2.
Можно менять N. Генерится дерево из N
вершин, делится на две части  и проверяется
размеры полученных деревьев.
 */
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

func TestSingle(t *testing.T) {
	countVertex := 20
	maxWeight := 30
	maxProfit := 1.0

	tree := GenerateTree(countVertex, maxWeight, maxProfit)
	for targetBound := 0; targetBound <= countVertex * maxWeight; targetBound+=10 {
		actualAnswer, actualSubset := knapsack.FindOptimalAnswerAndSubset(tree, targetBound)
		expectedAnswer, expectedSubset := knapsack.SimpleKnapsack(tree, targetBound)

		optimalSubsetID := []int{}
		bruteforceSubsetID := []int{}
		for _, node := range(actualSubset) {
			optimalSubsetID = append(optimalSubsetID, node.ID)
		}
		for _, node := range(expectedSubset) {
			bruteforceSubsetID = append(bruteforceSubsetID, node.ID)
		}
		sort.Slice(optimalSubsetID, func(i, j int) bool {
			return optimalSubsetID[i] < optimalSubsetID[j]
			})

		sort.Slice(bruteforceSubsetID, func(i, j int) bool {
			return bruteforceSubsetID[i] < bruteforceSubsetID[j]
			})

		status := ""
		if math.Abs(expectedAnswer-actualAnswer) < 0.0000000001 {
			fmt.Println("Bound: ", targetBound,
				"Optimal subset:", optimalSubsetID,
				"Bruteforce subset:", bruteforceSubsetID,
				"Bruteforce answer:", expectedAnswer,
				"Optimal answer:", actualAnswer)
		} else {
			t.Error("Wrong answer!!!",
				"Bound: ", targetBound,
				"Status:", status,
				"Bruteforce answer:", expectedAnswer,
				"Optimal answer:", actualAnswer,
				"Optimal subset:", optimalSubsetID,
				"Bruteforce subset:", bruteforceSubsetID)
		}

	}
}

func TestFullIteration(t *testing.T) {
	N := 1000
	MaxWeight := 10
	W := 500
	MaxProfit := 10.0
	for i := 0; i < 100; i++ {
		fmt.Println("Test:", i)
		tree := GenerateTree(N, MaxWeight, MaxProfit)
		ansRequired, setRequired := knapsack.SimpleKnapsack(tree, W)
		ansGet, setGet := knapsack.FindOptimalAnswerAndSubset(tree, W)
		fmt.Println(ansRequired, ansGet)
		fmt.Println(NodeToID(setGet))
		fmt.Println(NodeToID(setRequired))
		if math.Abs(ansRequired - ansGet) > 0.00000001 {
			t.Error("You are looser")
		}
	}

}

/**
Сравнение время работы обычного рюкзака (квадратичная память)
и оптимального (линейновысотная память). Выводит массив отношений
времени работы второго к первому, а потом среднее арифметичемкое.
Стандартно выводит 3.5.
 */
func TestFOAS_time(t *testing.T) {
	N := 10000
	MAXWEIGHT := 10
	W := 10000
	MAXPROFIT := 10.0

	var divTime []float64
	//1.7
	COUNTTEST := 10
	for i := 0; i < COUNTTEST; i++ {
		tree := GenerateTree(N, MAXWEIGHT, MAXPROFIT)

		timeNow := time.Now()
		knapsack.SimpleKnapsack(tree, W)
		time1 := time.Now().Sub(timeNow).Seconds()

		timeNow = time.Now()
		knapsack.FindOptimalAnswerAndSubset(tree, W)
		time2 := time.Now().Sub(timeNow).Seconds()

		fmt.Println("Simple knapsack time:", time1, "Knapsack time", time2)

		divTime = append(divTime, time2/time1)
		//fmt.Println("FOP_time:", knapsack.FOP_time, "FOP_count / NW:", float64(knapsack.FOP_count)/float64(N*W), ", All time:", time2)
	}
	sum := 0.0
	for _, differ := range divTime {
		sum += differ
	}
	fmt.Println(divTime)
	fmt.Print(sum / float64(COUNTTEST))
}

/**
Сравнивает время работы обычного рюкзака и хорощего.
Так же выводит общее время работы последнего,
время записей в динамику и отношение кол-ва записей
в динамике к n*W.
Доказывает тем самым, что большинство (99%) тратится
именно на них.
Так как сравниваем с обычным рюкзаком, то лучше больше
10^8 не тестировать.
 */
func TestFOPcompareFOAStime(t *testing.T) {
	N := 10000
	MAXWEIGHT := 10
	W := 10000
	MAXPROFIT := 10.0

	tree := GenerateTree(N, MAXWEIGHT, MAXPROFIT)

	timeNow := time.Now()
	knapsack.SimpleKnapsack(tree, W)
	time1 := time.Now().Sub(timeNow).Seconds()

	timeNow = time.Now()
	knapsack.FindOptimalAnswerAndSubset(tree, W)
	time2 := time.Now().Sub(timeNow).Seconds()

	fmt.Println("Simple knapsack time:", time1)
	fmt.Println("Optimal knapsack time:", time2)
	fmt.Println("Optimal dynamic time:", knapsack.FOP_dp_time)
	fmt.Println("Optimal count records / nW:", float64(knapsack.FOP_count) / float64(N * W))
}

/**
Почему-то map медленная.
 */
func TestMap(t *testing.T) {
	mp := make(map[uint32]int)
	arr := make([]int, 5)

	timeNow := time.Now()
	for i := 0; i < 10000000; i++ {
		mp[5] = 1
	}
	timeMap := time.Now().Sub(timeNow).Seconds()

	timeNow = time.Now()
	for i := 0; i < 10000000; i++ {
		arr[2] = 2
	}
	timeArr := time.Now().Sub(timeNow).Seconds()

	fmt.Println(timeArr, timeMap)
}

func NodeToID(nodes []*tries.Node) []int {
	var xs []int
	for _, node := range nodes {
		xs = append(xs, node.ID)
	}
	sort.Ints(xs)
	return xs
}