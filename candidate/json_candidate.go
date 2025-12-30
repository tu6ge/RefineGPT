package candidate

import (
	"encoding/json"
)

// JSONCandidate 是一个通用的 Candidate 实现
// 内部只存原始 JSON，不关心业务结构
type JSONCandidate struct {
	raw []byte
}

// NewJSONCandidateFromBytes 用原始 JSON 构造
func NewJSONCandidateFromBytes(raw []byte) (*JSONCandidate, error) {
	if !json.Valid(raw) {
		return nil, ErrInvalidJSON
	}
	return &JSONCandidate{raw: raw}, nil
}

// NewJSONCandidateFromAny 用任意结构体构造
func NewJSONCandidateFromAny(v any) (*JSONCandidate, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return &JSONCandidate{raw: raw}, nil
}

// Raw 返回原始 JSON
func (c *JSONCandidate) Raw() []byte {
	return c.raw
}

// As 反序列化为指定结构
func (c *JSONCandidate) As(v any) error {
	return json.Unmarshal(c.raw, v)
}
