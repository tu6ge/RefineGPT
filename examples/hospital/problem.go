package main

import (
	"github.com/tu6ge/RefineGPT/framework/core"
	"github.com/tu6ge/RefineGPT/framework/schema"
)

type HospitalProblem struct{}

func (h *HospitalProblem) GetName() string {
	return "hospital-scheduling"
}

func (h *HospitalProblem) GetSchema() schema.JSONSchema {
	return schema.JSONSchema{
		"type": "object",
	}
}

func (h *HospitalProblem) GetConstraints() []core.Constraint {
	return []core.Constraint{
		&NoDuplicateShift{},
	}
}

func (h *HospitalProblem) GetContext() map[string]any {
	return map[string]any{
		"people": []string{"Alice", "Bob"},
		"shifts": []string{"Morning"},
	}
}
