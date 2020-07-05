// Package entity contains common data structures
package entity

// PortDetails stores port fields
type PortDetails struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Alias    []string  `json:"alias"`
	Coords   []float64 `json:"coordinates"`
	Province string    `json:"province"`
	Timezone string    `json:"timezone"`
	Unlocs   []string  `json:"unlocs"`
	Code     string    `json:"code"`
	Regions  []string  `json:"regions"`
}
