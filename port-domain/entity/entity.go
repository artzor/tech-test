// Package entity contains common data structures
package entity

// PortDetails stores port fields
type PortDetails struct {
	ID       string    `bson:"_id"`
	Name     string    `bson:"name"`
	City     string    `bson:"city"`
	Country  string    `bson:"country"`
	Alias    []string  `bson:"alias"`
	Coords   []float64 `bson:"coordinates"`
	Province string    `bson:"province"`
	Timezone string    `bson:"timezone"`
	Unlocs   []string  `bson:"unlocs"`
	Code     string    `bson:"code"`
	Regions  []string  `bson:"regions"`
}
