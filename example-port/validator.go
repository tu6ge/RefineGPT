package main

import (
	"context"
	"fmt"

	"github.com/tu6ge/RefineGPT/engine"
)

type LengthRule struct{}

func (dr *LengthRule) Validate(
	ctx context.Context,
	data engine.State,
	c engine.Candidate,
) ([]engine.Feedback, error) {

	d := data.Value().(ShipBerth)
	list := c.(*AssignmentList)
	var result []engine.Feedback

	for _, item := range list.list {
		ship := d.ships[item.ShipID]
		berth := d.berths[item.BerthID]
		if ship.Length > berth.Length {
			result = append(result, engine.Feedback{
				Code:   "LENGTH_EXCEEDS_BERTH",
				Target: ship.ID,
				Message: fmt.Sprintf("船舶长度%.2f米超过泊位长度%.2f米",
					ship.Length, berth.Length),
				Severity: engine.SeverityFixable,
			})
		}
	}
	if len(result) > 0 {
		return result, nil
	}
	return nil, nil
}
