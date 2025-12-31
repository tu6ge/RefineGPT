package main

import (
	"context"
	"encoding/json"

	"github.com/tu6ge/RefineGPT/engine"
)

type CandidateFactory struct{}

func (cf *CandidateFactory) FromLLMOutput(ctx context.Context, raw string) (engine.Candidate, error) {
	var obj MapCandidate
	json.Unmarshal([]byte(raw), &obj)
	return obj, nil
}

type MapCandidate map[string]any

func (m MapCandidate) Raw() []byte {
	b, _ := json.Marshal(m)
	return b
}

func (m MapCandidate) As(v any) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
