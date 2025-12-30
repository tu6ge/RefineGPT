package llm

type Role string

const (
	RoleSystem Role = "system"
	RoleUser   Role = "user"
	RoleAssist Role = "assistant"
)

type Message struct {
	Role    Role
	Content string
}
