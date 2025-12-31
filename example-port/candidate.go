package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/tu6ge/RefineGPT/engine"
)

type Assignment struct {
	ShipID    string    `json:"ship_id"`
	BerthID   string    `json:"berth_id"`
	StartTime time.Time `json:"start_time"`
	Hours     int       `json:"hours"`
}
type AssignmentList struct {
	list []Assignment
}

var _ engine.Candidate = (*AssignmentList)(nil)

func (c *AssignmentList) Raw() []byte {
	raw, _ := json.Marshal(c)
	return raw
}

// As 反序列化为指定结构
func (c *AssignmentList) As(v any) error {
	return json.Unmarshal(c.Raw(), v)
}

func (a *Assignment) EndTime() time.Time {
	return a.StartTime.Add(time.Duration(a.Hours) * time.Hour)
}

func (a AssignmentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.list)
}

func (a *AssignmentList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.list)
}

type CandidateFactory struct{}

func (cf *CandidateFactory) FromLLMOutput(ctx context.Context, raw string) (engine.Candidate, error) {
	var obj AssignmentList
	err := json.Unmarshal([]byte(raw), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
