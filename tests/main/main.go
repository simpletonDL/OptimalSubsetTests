package main

import (
	"fmt"
	"sync"
	"github.com/simpletonDL/OptimalSubsetTests"
)

func main() {
	/*tree := Tree{Root:NewNode(0,0,0)}
	tree.Root.AddChild(NewNode(0,0,1))
	tree.Root.Children[0].AddChild(NewNode(0,0,2))
	tree.Root.Children[0].AddChild(NewNode(0,0,3))
	tree.Root.Children[0].AddChild(NewNode(0,0,4))
	tree.Root.Children[0].Children[2].AddChild(NewNode(0,0,7))
	tree.Root.Children[0].Children[0].AddChild(NewNode(0,0,8))
	tree.Root.Children[0].Children[0].AddChild(NewNode(0,0,9))
	tree.Root.AddChild(NewNode(0,0,5))
	tree.Root.Children[1].AddChild(NewNode(0,0,6))
	tree.UpdateSizes()
	knapsack.SortChildren(tree)
	tree.Print()*/

	/*tree := Tree{Root:NewNode(0,0,0)}
	tree.Root.AddChild(NewNode(0,0,1))
	tree.Root.AddChild(NewNode(0,0,2))
	tree.Root.Children[1].AddChild(NewNode(0,0,3))
	tree.Root.Children[1].AddChild(NewNode(0,0,4))
	tree.Root.Children[1].Children[1].AddChild(NewNode(0,0,5))
	tree.Root.Children[1].Children[1].AddChild(NewNode(0,0,6))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,7))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,8))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,9))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,10))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,11))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(234,0,12))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,13))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,14))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,15))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,16))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,17))
	tree.Root.Children[1].Children[1].Children[1].AddChild(NewNode(0,0,18))
	tree.UpdateSizes()*/

	/*tree1, tree2, split := SplitTree(tree)
	tree1.Print()
	fmt.Println()
	tree2.Print()
	fmt.Println()
	fmt.Println(split.ID)*/

	/*tree := Tree{Root:NewNode(0,0,0)}
	tree.Root.AddChild(NewNode(0,0,1))
	tree.Root.AddChild(NewNode(0,0,2))
	tree.UpdateSizes()
	tree1, tree2, split := SplitTree(tree)
	tree1.Print()
	fmt.Println()
	tree2.Print()
	fmt.Println()
	fmt.Println(split.ID)*/

	/*tree.Root.Children[1].SetRequired()
	answer := FindOptimalProbability(tree, 10)
	fmt.Println(answer)*/

	/*tree := tests.GenerateTree(4, 10, 10)
	fmt.Println(knapsack.FindOptimalProbability(tree, 10)[10])
	fmt.Print(knapsack.FindOptimalSubset(tree, 10))*/

	/*waitGroup := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		j := i

		waitGroup.Add(1)

		go func() {
			a(j, waitGroup)
		}()
	}

	waitGroup.Wait()

	fmt.Println("Hello, playground 1")*/

	/*solver := ost.NewSolver()
	solver.AddTest("10", "", 10, 6)
	solver.AddTest("20", "10", 20, 7)
	solver.AddTest("30", "", 30, 5)
	fmt.Print(solver.Solve(30, 5))*/

	solver := ost.NewSolver()
	solver.AddTest("1", "", 7, 7)
	solver.AddTest("2", "", 10, 10)
	solver.AddTest("3", "2", 4, 4)
	solver.AddTest("4", "2", 6, 6)
	solver.AddTest("5", "1", 3, 3)

	solution1 := solver.Solve(20, 1)
	fmt.Println(solution1.GetAnswers())
	fmt.Println(solution1.ComputeTests(16))
	fmt.Println(solution1.ComputeTests(20))

	solver.AddTest("6", "", 1, 100)
	solution2 := solver.Solve(20, 1)
	fmt.Println(solution2.GetAnswers())
	fmt.Print(solution2.ComputeTests(len(solution2.GetAnswers()) - 1))
}

//cntr alt l

func a(i int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println(i)
}