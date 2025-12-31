package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/tu6ge/RefineGPT/engine"
	"github.com/tu6ge/RefineGPT/generator"
	"github.com/tu6ge/RefineGPT/llm"
)

const systemPrompt = `
你是一个港口停泊调度系统，正在解决一个【约束满足问题（Constraint Satisfaction Problem）】。
你的目标不是生成看起来合理的方案，而是生成【严格满足所有硬约束】的停泊计划。
只要违反任意一条硬约束，该方案即视为失败。


【硬约束规则（必须全部满足）】
1. 船舶吃水 ≤ 泊位最大吃水
2. 船舶长度 ≤ 泊位长度
3. 船舶货物类型 ∈ 泊位允许的货物类型
4. 每艘船舶最多只能分配一个泊位
5. 开始停泊时间 ≥ 船舶到港时间
6. 同一泊位在同一时间段内只能停靠一艘船
7. 停泊时间 ∈ 允许停泊的时间范围
`

func formatFeedback(feedback []engine.Feedback) string {
	if len(feedback) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("【必须修正的错误（以下问题在本轮必须解决）】\n")
	for _, err := range feedback {
		b.WriteString(fmt.Sprintf("- ShipID: %s\n", err.Target))
		for _, msg := range err.Message {
			b.WriteString(fmt.Sprintf("  - Violation: %s\n", msg))
		}
	}
	b.WriteString(
		"\n注意：如果某艘船未出现在以上错误列表中，说明其分配是正确的，不允许修改。\n\n",
	)
	b.WriteString("【修正原则（必须遵守）】\n")
	b.WriteString("- 仅修改存在错误的船舶分配\n")
	b.WriteString("- 优先只调整 berth_id\n")
	b.WriteString("- 只有在无法满足约束时，才允许调整 start_time 或 hours\n")
	b.WriteString("- 不得引入新的船舶或泊位\n")
	b.WriteString("- 不得删除任何船舶\n\n")
	return b.String()
}

func formatState(state engine.State) string {
	if state == nil {
		return ""
	}

	request := state.(ShipBerth)
	var b strings.Builder

	b.WriteString("【船舶信息】\n")
	for _, ship := range request.ships {
		b.WriteString(fmt.Sprintf(
			"- ShipID: %s\n  Name: %s\n  ArrivalTime: %s\n  Draft: %.2f\n  Length: %.2f\n  CargoType: %s\n",
			ship.ID,
			ship.Name,
			ship.ArrivalTime.Format(time.RFC3339),
			ship.Draft,
			ship.Length,
			ship.CargoType,
		))
	}
	b.WriteString("\n")

	// ===== 4. 泊位信息 =====
	b.WriteString("【泊位信息】\n")
	for _, berth := range request.berths {

		var times strings.Builder
		for i, t := range berth.Availability {
			times.WriteString(fmt.Sprintf("    - Availability: %s ~ %s\n",
				t.Start.Format(time.RFC3339),
				t.End.Format(time.RFC3339)))
			times.WriteString(fmt.Sprintf("    - MaxDraft: %.2f\n", berth.getMaxDraft(i)))
		}

		b.WriteString(fmt.Sprintf(
			"- BerthID: %s\n  Name: %s\n  Length: %.2f\n  AllowedCargoTypes: [%s]\n  AvailabilityTimeWindow:\n%s",
			berth.ID,
			berth.Name,
			berth.Length,
			strings.Join(berth.CargoTypes, ", "),
			times.String(),
		))
	}
	b.WriteString("\n")
	return b.String()
}

type DispathPromptAdapter struct{}

func NewDispathPromptAdapter() *DispathPromptAdapter {
	return &DispathPromptAdapter{}
}

func (a *DispathPromptAdapter) BuildMessages(
	ctx generator.GenerateContext,
) ([]llm.Message, error) {

	var userParts []string

	if ctx.State != nil {
		userParts = append(userParts,
			formatState(ctx.State))
	}

	if ctx.Schema != "" {
		userParts = append(userParts,
			"输出必须符合以下 JSON Schema：\n"+ctx.Schema)
	}

	if len(ctx.Feedback) > 0 {
		userParts = append(userParts,
			"上一轮校验反馈（请修正）：\n"+formatFeedback(ctx.Feedback))
	}

	// ===== 8. 输出格式（机器级严格）=====
	userParts = append(userParts,
		"【输出要求】\n"+
			"- 只允许输出 JSON\n"+
			"- 不得包含任何解释性文本\n"+
			"- 必须是一个数组\n\n",
	)

	userParts = append(userParts,
		`[
  {
    "ship_id": "ShipID",
    "berth_id": "BerthID 或 null",
    "start_time": "RFC3339 格式时间",
    "hours": 24
  }
]`,
	)

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
