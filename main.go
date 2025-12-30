package main

import "fmt"

func main() {
	engine := NewRuleEngine(
		AgentNotEmptyRule{},
		NoOverlapRule{},
	)

	llm := &MockLLM{}

	loop := LoopManager{
		MaxIterations: 5,
		LLM:           llm,
		Engine:        engine,
	}

	decision, err := loop.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Final Decision:\n%+v\n", decision)
}
