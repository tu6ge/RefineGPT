package engine

type Severity string

const (
	SeverityFixable Severity = "fixable"
	SeverityFatal   Severity = "fatal"
	SeverityWarn    Severity = "warn"
)

// Feedback 是规则失败的结构化表达
type Feedback struct {
	Code     string         `json:"code"`
	Target   string         `json:"target"` // JSONPath / 字段名 / 语义位置
	Message  string         `json:"message"`
	Severity Severity       `json:"severity"`
	Meta     map[string]any `json:"meta,omitempty"`
}

func HasFatal(feedbacks []Feedback) bool {
	for _, f := range feedbacks {
		if f.Severity == SeverityFatal {
			return true
		}
	}
	return false
}
