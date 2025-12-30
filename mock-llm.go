package main

type MockLLM struct {
	step int
}

func (m *MockLLM) GenerateDecision(feedback []RuleViolation) Decision {
	m.step++

	// 第一次：明显错误
	if m.step == 1 {
		return Decision{
			Assignments: []Assignment{
				{Agent: "", Task: "A", Time: "T1"},
				{Agent: "Bob", Task: "B", Time: "T1"},
				{Agent: "Bob", Task: "C", Time: "T1"},
			},
		}
	}

	// 第二次：修复错误
	return Decision{
		Assignments: []Assignment{
			{Agent: "Alice", Task: "A", Time: "T1"},
			{Agent: "Bob", Task: "B", Time: "T1"},
			{Agent: "Bob", Task: "C", Time: "T2"},
		},
	}
}
