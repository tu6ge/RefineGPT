package main

import (
	"context"
	"fmt"

	"github.com/tu6ge/RefineGPT/engine"
	"github.com/tu6ge/RefineGPT/generator"
)

func main() {
	ctx := context.Background()

	// 1️⃣ State
	state := ExampleState{
		Task: "decide whether to allow access",
	}

	// 3️⃣ Generator
	gen := &generator.LLMGenerator{
		Client:  &MockLLM{},
		Adapter: generator.NewDefaultPromptAdapter(),
		Parser:  &CandidateFactory{},
		Schema: `
{
  "type": "object",
  "properties": {
    "action": {
      "type": "string",
      "enum": ["allow", "deny"]
    }
  },
  "required": ["action"]
}
`,
	}

	// 4️⃣ Engine
	e := &engine.Engine{
		Generator: gen,
		Validator: &ExampleValidator{},
		Policy: engine.LoopPolicy{
			MaxIteration: 5,
			StopOnFatal:  true,
		},
	}

	// 5️⃣ Run
	result, feedbacks, err := e.Run(ctx, state)
	if err != nil {
		panic(err)
	}

	fmt.Println("Final Candidate:", string(result.Raw()))
	fmt.Println("Feedback History:", feedbacks)
}
