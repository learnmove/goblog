package models

type Page struct {
	LastPage    float64 `json:"last_page"`
	Total       float64 `json:"total"`
	CurrentPage float64 `json:"current_page"`
	PerPage     float64 `json:"per_page"`
	From        float64 `json:"from"`
	To          float64 `json:"to"`
}
