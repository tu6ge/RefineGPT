package main

import (
	"time"
)

// 船舶信息
type Ship struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ArrivalTime time.Time `json:"arrival_time"`
	Draft       float64   `json:"draft"`
	CargoType   string    `json:"cargo_type"`
	Length      float64   `json:"length"`
}

// 泊位信息
type Berth struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	MaxDraft     float64      `json:"max_draft"`
	Length       float64      `json:"length"`
	CargoTypes   []string     `json:"cargo_types"`
	Availability []TimeWindow `json:"availability"`
}

func (b *Berth) getMaxDraft(index int) float64 {
	if index >= 0 && index < len(b.Availability) {
		if b.Availability[index].MaxDraft != 0 {
			return b.Availability[index].MaxDraft
		}
		return b.MaxDraft
	}
	return 0
}

// 时间窗口（左闭右开，推荐）
type TimeWindow struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	MaxDraft float64   `json:"max_draft"`
}
