package validator

import (
	"context"
	"sync"

	"github.com/tu6ge/RefineGPT/engine"
)

type CompositeValidator struct {
	validators []engine.Validator
	policy     Policy
}

func NewComposite(
	validators []engine.Validator,
	policy Policy,
) *CompositeValidator {
	if policy.Mode == "" {
		policy = DefaultPolicy()
	}
	return &CompositeValidator{
		validators: validators,
		policy:     policy,
	}
}

func (c *CompositeValidator) Validate(
	ctx context.Context,
	candidate engine.Candidate,
) ([]engine.Feedback, error) {

	switch c.policy.Mode {
	case ModeParallel:
		return c.validateParallel(ctx, candidate)
	default:
		return c.validateSequential(ctx, candidate)
	}
}

// 串行模式
func (c *CompositeValidator) validateSequential(
	ctx context.Context,
	candidate engine.Candidate,
) ([]engine.Feedback, error) {

	var all []engine.Feedback

	for _, v := range c.validators {
		feedbacks, err := v.Validate(ctx, candidate)
		if err != nil {
			return all, err
		}

		all = append(all, feedbacks...)

		if c.policy.StopOnFatal && engine.HasFatal(feedbacks) {
			break
		}

		if c.policy.MaxFeedbackNum > 0 &&
			len(all) >= c.policy.MaxFeedbackNum {
			break
		}
	}

	return all, nil
}

// 并行模式
func (c *CompositeValidator) validateParallel(
	ctx context.Context,
	candidate engine.Candidate,
) ([]engine.Feedback, error) {

	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		all  []engine.Feedback
		errs []error
	)

	for _, v := range c.validators {
		wg.Add(1)

		go func(v engine.Validator) {
			defer wg.Done()

			feedbacks, err := v.Validate(ctx, candidate)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errs = append(errs, err)
				return
			}

			all = append(all, feedbacks...)
		}(v)
	}

	wg.Wait()

	if len(errs) > 0 {
		return all, errs[0]
	}

	if c.policy.MaxFeedbackNum > 0 &&
		len(all) > c.policy.MaxFeedbackNum {
		all = all[:c.policy.MaxFeedbackNum]
	}

	return all, nil
}
