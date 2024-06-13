package schemas

// PerformanceSchema represents the detailed information of a performance
type PerformanceSchema struct {
	ItemSchema
	Performers []string `json:"performers"`
}

// PerformanceProductionSchema represents the detailed information of a performance production
type PerformanceProductionSchema struct {
	ItemSchema
	Producer  string `json:"producer"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
