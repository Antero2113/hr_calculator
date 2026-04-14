package models

type Record struct {
	Name           string `json:"name"`
	DepartmentFull string `json:"department_full"`
	Position       string `json:"position"`
	Client         string `json:"client"`
	ProcessName    string `json:"process_name"`
	Operations     string `json:"operations"`
	Measure        string `json:"measure"`
	Min            int    `json:"min"`
	Max            int    `json:"max"`
	PeriodType     string `json:"period_type"`
	PeriodCount    int    `json:"period_count"`
}