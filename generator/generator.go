package generator

import (
	"context"

	"github.com/tu6ge/RefineGPT/engine"
	"github.com/tu6ge/RefineGPT/llm"
)

type LLMGenerator struct {
	Client  llm.Client
	Adapter PromptAdapter
	Parser  engine.CandidateParser
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

	raw, err := g.Client.Complete(ctx, messages)
	if err != nil {
		return nil, err
	}

	return g.Parser.Parse(ctx, raw)
}
