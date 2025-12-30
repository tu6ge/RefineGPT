package main

import (
	"context"
	"strings"

	"github.com/tu6ge/RefineGPT/llm"
)

type MockLLM struct{}

func (m *MockLLM) Complete(
	ctx context.Context,
	messages []llm.Message,
) (string, error) {

	// 简单模拟“根据反馈修正”
	joined := ""
	for _, msg := range messages {
		joined += msg.Content
	}

	if strings.Contains(joined, "INVALID_ACTION") {
		return `{"action":"allow"}`, nil
	}

	if strings.Contains(joined, "MISSING_FIELD") {
		return `{"action":"deny"}`, nil
	}

	// 第一次，返回一个明显错误的
	return `{"foo":"bar"}`, nil
}
