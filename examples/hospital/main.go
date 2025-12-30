package main

import (
	"fmt"

	"github.com/tu6ge/RefineGPT/framework/core"
	"github.com/tu6ge/RefineGPT/framework/llm"
)

func main() {
	problem := &HospitalProblem{}

	engine := &core.RuleEngine{
		Constraints: problem.GetConstraints(),
	}

	solver := &core.Solver{
		Planner:      &llm.MockPlanner{},
		RuleEngine:   engine,
		MaxIteration: 5,
	}

	result, err := solver.Solve(problem)
	if err != nil {
		panic(err)
	}

	fmt.Println("Final Decision:", result)
}
