package main

import (
	"context"

	"github.com/tu6ge/RefineGPT/engine"
)

type ExampleValidator struct{}

func (v *ExampleValidator) Validate(
	ctx context.Context,
	state engine.State,
	c engine.Candidate,
) ([]engine.Feedback, error) {

	var obj map[string]any
	if err := c.As(&obj); err != nil {
		return []engine.Feedback{
			{
				Code:     "INVALID_JSON",
				Message:  err.Error(),
				Severity: engine.SeverityFatal,
			},
		}, nil
	}

	// 规则 1：必须有 action 字段
	if _, ok := obj["action"]; !ok {
		return []engine.Feedback{
			{
				Code:     "MISSING_FIELD",
				Target:   "$.action",
				Message:  "field `action` is required",
				Severity: engine.SeverityFixable,
			},
		}, nil
	}

	// 规则 2：action 只能是 allow / deny
	action, _ := obj["action"].(string)
	if action != "allow" && action != "deny" {
		return []engine.Feedback{
			{
				Code:     "INVALID_ACTION",
				Target:   "$.action",
				Message:  "action must be allow or deny",
				Severity: engine.SeverityFixable,
			},
		}, nil
	}

	return nil, nil
}
