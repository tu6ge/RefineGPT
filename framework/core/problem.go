package core

import "github.com/tu6ge/RefineGPT/framework/schema"

type Problem interface {
	GetName() string
	GetSchema() schema.JSONSchema
	GetConstraints() []Constraint
	GetContext() map[string]any
}
