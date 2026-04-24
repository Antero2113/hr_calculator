package models

type Record struct {
	Position     string `json:"position"`
	Client       string `json:"client"`
	Operations   string `json:"operations"`
	Measure      string `json:"measure"`
	Min          int    `json:"min"`
	Max          int    `json:"max"`
	PeriodType   string `json:"period_type"`
	PeriodCount  int    `json:"period_count"`
}