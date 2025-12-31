package main

type ShipBerth struct {
	ships  map[string]Ship
	berths map[string]Berth
}

func NewShipBerth(ships []Ship, berths []Berth) ShipBerth {
	shipMap := make(map[string]Ship)
	berthMap := make(map[string]Berth)

	for _, ship := range ships {
		shipMap[ship.ID] = ship
	}
	for _, berth := range berths {
		berthMap[berth.ID] = berth
	}
	return ShipBerth{
		ships:  shipMap,
		berths: berthMap,
	}
}

func (sb ShipBerth) Value() any {
	return sb
}
