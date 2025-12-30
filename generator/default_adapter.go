package generator

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tu6ge/RefineGPT/engine"
	"github.com/tu6ge/RefineGPT/llm"
)

const systemPrompt = `
你是一个决策生成器。
你必须严格输出 JSON。
不要输出解释性文字。
如果收到错误反馈，请只修正相关字段。
`

func formatFeedback(feedback []engine.Feedback) string {
	if len(feedback) == 0 {
		return ""
	}

	raw, _ := json.MarshalIndent(feedback, "", "  ")
	return string(raw)
}

func formatState(state engine.State) string {
	if state == nil {
		return ""
	}

	b, err := json.MarshalIndent(state.Value(), "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", state.Value())
	}
	return string(b)
}

type DefaultPromptAdapter struct{}

func NewDefaultPromptAdapter() *DefaultPromptAdapter {
	return &DefaultPromptAdapter{}
}

func (a *DefaultPromptAdapter) BuildMessages(
	ctx GenerateContext,
) ([]llm.Message, error) {

	var userParts []string

	if ctx.State != nil {
		userParts = append(userParts,
			"当前状态：\n"+formatState(ctx.State))
	}

	if ctx.Schema != "" {
		userParts = append(userParts,
			"输出必须符合以下 JSON Schema：\n"+ctx.Schema)
	}

	if len(ctx.Feedback) > 0 {
		userParts = append(userParts,
			"上一轮校验反馈（请修正）：\n"+formatFeedback(ctx.Feedback))
	}

	userParts = append(userParts,
		"请输出新的 JSON 决策。")

	return []llm.Message{
		{
			Role:    llm.RoleSystem,
			Content: strings.TrimSpace(systemPrompt),
		},
		{
			Role:    llm.RoleUser,
			Content: strings.Join(userParts, "\n\n"),
		},
	}, nil
}
