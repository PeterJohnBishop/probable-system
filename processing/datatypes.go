package processing

type Trip struct {
	RouteID      string `json:"route_id"`
	ServiceID    string `json:"service_id"`
	TripID       string `json:"trip_id"`
	TripHeadsign string `json:"trip_headsign"`
	DirectionID  int    `json:"direction_id"`
	BlockID      string `json:"block_id"`
	ShapeID      string `json:"shape_id"`
}
