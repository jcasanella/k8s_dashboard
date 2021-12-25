package models

type Pod struct {
	Continue           string `json:"continue"`
	Name               string `json:"name"`
	RemainingItemCount int64  `json:"remainingitemcount"`
}
