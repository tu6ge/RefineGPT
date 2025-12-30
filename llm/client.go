package llm

import "context"

type Client interface {
	Complete(ctx context.Context, messages []Message) (string, error)
}
