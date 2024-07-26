package models

type Result struct {
	ID         string `json:"id"`
	RunnerID   string `json:"runnerId"`
	RaceResult string `json:"raceResult"`
	Location   string `json:"location"`
	Position   int    `json:"position,omitempty"`
	Year       int    `json:"year"`
}
