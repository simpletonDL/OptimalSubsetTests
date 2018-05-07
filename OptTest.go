package ost

import (
	"time"
	"github.com/simpletonDL/OptimalSubsetTests/knapsack"
	. "github.com/simpletonDL/OptimalSubsetTests/trees"
)

type Solver interface {
    AddTest(id, parentId string, duration time.Duration, failureProbability float64)
	Solve(boundDuration, step time.Duration) Solution
}

type Solution interface {
	GetStep() time.Duration
	GetAnswers() []float64
	ComputeTests(index int) []string // todo return map[string]bool
}

/*
В дереве timeByStep
 */
type solution struct {
	step time.Duration
	answers []float64
	tree *Tree
}

func (slt *solution) GetStep() time.Duration {
	return slt.step
}

func (slt *solution) GetAnswers() []float64 {
	return slt.answers
}

/*
В дереве Weight уже установлен от step корректно
 */
func (slt *solution) ComputeTests(index int) []string {
	_, nodes := knapsack.FindOptimalAnswerAndSubset(slt.tree, int64(index))
	return NodeToID(nodes)
}

type solver struct {
	tree  *Tree
	tests map[string]*Node
}

/*
Создаём дерево с фиктивным корнем.
 */
func NewSolver() Solver {
	return &solver{
		tree: &Tree{Root:NewNode(0, 0,"")},
		tests: make(map[string]*Node),
	}
}

func (slv *solver) AddTest(id, parentId string, duration time.Duration, failureProbability float64) {
	slv.tests[id] = NewNode(0, failureProbability, id)
	slv.tests[id].TimeDuration = duration
	if parentId != "" {
		slv.tests[parentId].AddChild(slv.tests[id])
	}
}

/*
Подвешиваем все отдельные деревья к slv.tree.Root
и устанавливает node.Weight относительно step.
 */
func (slv *solver) build(step time.Duration) {
	for _, node := range slv.tests {
		node.Weight = (node.TimeDuration / step).Nanoseconds()
		if node.Parent == nil {
			slv.tree.Root.AddChild(node)
		}
	}
}

func (slv *solver) Solve(boundDuration, step time.Duration) Solution {
	slv.build(step)

	var countSteps int64 = (boundDuration / step).Nanoseconds()
	answers := knapsack.FindOptimalProbability(slv.tree, countSteps)

	return &solution{step:step, answers:answers, tree: slv.tree.Copy()}
}

/*var subsetId []string
	for _, node := range subset {
		subsetId = append(subsetId, node.ID)
	} */