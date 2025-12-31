package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tu6ge/RefineGPT/engine"
	"github.com/tu6ge/RefineGPT/generator"
	"github.com/tu6ge/RefineGPT/validator"
)

func main() {
	ctx := context.Background()

	base := time.Date(2025, 1, 10, 8, 0, 0, 0, time.UTC)

	ships := []Ship{
		{
			ID:          "ship1",
			Name:        "集装箱船A",
			ArrivalTime: base.Add(2 * time.Hour), // 10:00
			Draft:       12.5,
			CargoType:   "container",
			Length:      220,
		},
		{
			ID:          "ship2",
			Name:        "集装箱船B",
			ArrivalTime: base.Add(3 * time.Hour), // 11:00
			Draft:       11.8,
			CargoType:   "container",
			Length:      210,
		},
		{
			ID:          "ship3",
			Name:        "油轮C",
			ArrivalTime: base.Add(6 * time.Hour), // 14:00
			Draft:       15.0,
			CargoType:   "oil",
			Length:      280,
		},
		{
			ID:          "ship4",
			Name:        "油轮D",
			ArrivalTime: base.Add(7 * time.Hour), // 15:00
			Draft:       14.5,
			CargoType:   "oil",
			Length:      270,
		},
		{
			ID:          "ship5",
			Name:        "散货船E",
			ArrivalTime: base.Add(8 * time.Hour), // 16:00
			Draft:       9.5,
			CargoType:   "bulk",
			Length:      180,
		},
		{
			ID:          "ship6",
			Name:        "集装箱船F",
			ArrivalTime: base.Add(9 * time.Hour), // 17:00
			Draft:       10.5,
			CargoType:   "container",
			Length:      200,
		},
		{
			ID:          "ship7",
			Name:        "散货船G",
			ArrivalTime: base.Add(10 * time.Hour), // 18:00
			Draft:       8.5,
			CargoType:   "bulk",
			Length:      170,
		},
	}

	berths := []Berth{
		{
			ID:         "berth1",
			Name:       "集装箱泊位（潮汐）",
			MaxDraft:   12.0,
			Length:     250,
			CargoTypes: []string{"container"},
			Availability: []TimeWindow{
				{
					Start:    base.Add(0 * time.Hour), // 08:00
					End:      base.Add(6 * time.Hour), // 14:00
					MaxDraft: 13.0,                    // 放宽
				},
				{
					Start: base.Add(6 * time.Hour),  // 14:00
					End:   base.Add(12 * time.Hour), // 20:00
					// 默认 12.0
				},
			},
		},
		{
			ID:         "berth2",
			Name:       "油轮泊位（施工期）",
			MaxDraft:   18.0,
			Length:     300,
			CargoTypes: []string{"oil"},
			Availability: []TimeWindow{
				{
					Start:    base.Add(6 * time.Hour),  // 14:00
					End:      base.Add(10 * time.Hour), // 18:00
					MaxDraft: 14.0,                     // 收紧
				},
				{
					Start: base.Add(10 * time.Hour), // 18:00
					End:   base.Add(20 * time.Hour), // 次日 04:00
					// 回到 18.0
				},
			},
		},
		{
			ID:         "berth3",
			Name:       "通用泊位",
			MaxDraft:   10.0,
			Length:     200,
			CargoTypes: []string{"container", "bulk"},
			Availability: []TimeWindow{
				{
					Start: base.Add(0 * time.Hour), // 08:00
					End:   base.Add(8 * time.Hour), // 16:00
					// 使用默认 10.0
				},
				{
					Start:    base.Add(8 * time.Hour),  // 16:00
					End:      base.Add(14 * time.Hour), // 22:00
					MaxDraft: 9.0,                      // 更严格
				},
			},
		},
	}

	// 1️⃣ State
	state := NewShipBerth(ships, berths)

	// 2️⃣ Validator（支持以后扩展多个）
	v := validator.NewComposite(
		[]engine.Validator{
			&LengthRule{},
		},
		validator.DefaultPolicy(),
	)

	// 3️⃣ Generator
	gen := &generator.LLMGenerator{
		Client:  &MockLLM{},
		Adapter: &DispathPromptAdapter{},
		Parser:  &CandidateFactory{},
		Schema:  ``,
	}

	// 4️⃣ Engine
	e := &engine.Engine{
		Generator: gen,
		Validator: v,
		Policy: engine.LoopPolicy{
			MaxIteration: 5,
			StopOnFatal:  true,
		},
	}

	// 5️⃣ Run
	result, feedbacks, err := e.Run(ctx, state)
	if err != nil {
		panic(err)
	}

	fmt.Println("Final Candidate:", string(result.Raw()))
	fmt.Println("Feedback History:", feedbacks)
}
