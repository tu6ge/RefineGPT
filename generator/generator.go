package generator

import (
	"context"

	"github.com/tu6ge/RefineGPT/candidate"
	"github.com/tu6ge/RefineGPT/engine"
	"github.com/tu6ge/RefineGPT/llm"
)

type LLMGenerator struct {
	Client  llm.Client
	Adapter PromptAdapter
	Schema  string
}

func (g *LLMGenerator) Generate(
	ctx context.Context,
	input engine.GenerateInput,
) (engine.Candidate, error) {

	messages, err := g.Adapter.BuildMessages(GenerateContext{
		State:    input.State,
		Feedback: input.Feedback,
		Schema:   g.Schema,
	})
	if err != nil {
		return nil, err
	}

	output, err := g.Client.Complete(ctx, messages)
	if err != nil {
		return nil, err
	}

	return candidate.NewJSONCandidateFromBytes([]byte(output))
}
