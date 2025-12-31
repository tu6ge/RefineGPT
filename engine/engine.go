package engine

import "context"

// Engine 是整个框架的核心
type Engine struct {
	Generator Generator
	Validator Validator
	Policy    LoopPolicy
}

func (e *Engine) Run(ctx context.Context, state State) (Candidate, []Feedback, error) {
	var history []Feedback

	for i := 0; i < e.Policy.MaxIteration; i++ {
		candidate, err := e.Generator.Generate(ctx, GenerateInput{
			State:    state,
			Feedback: history,
		})
		if err != nil {
			return nil, history, err
		}

		feedbacks, err := e.Validator.Validate(ctx, state, candidate)
		if err != nil {
			return nil, history, err
		}

		if len(feedbacks) == 0 {
			return candidate, history, nil
		}

		history = append(history, feedbacks...)

		if e.Policy.StopOnFatal && HasFatal(feedbacks) {
			return candidate, history, ErrFatalFeedback
		}
	}

	return nil, history, ErrMaxIteration
}
