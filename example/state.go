package main

// ExampleState 是一个简单的任务描述
type ExampleState struct {
	Task string `json:"task"`
}

func (s ExampleState) Value() any {
	return s
}
