package models

type Result struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

type Status struct {
	Status bool `json:"status"`
}
