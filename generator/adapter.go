package generator

import (
	"github.com/tu6ge/RefineGPT/llm"
)

type PromptAdapter interface {
	BuildMessages(input GenerateContext) ([]llm.Message, error)
}
