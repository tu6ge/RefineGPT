package main

import (
	"context"

	"github.com/tu6ge/RefineGPT/llm"
)

type MockLLM struct{}

func (m *MockLLM) Complete(
	ctx context.Context,
	messages []llm.Message,
) (string, error) {

	return `[
  {
    "ship_id": "ship1",
    "berth_id": "berth1",
    "start_time": "2025-01-10T10:00:00Z",
    "hours": 4
  }]`, nil
}
