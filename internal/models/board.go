package models

type Board struct {
	Width     int          `json:"width"`
	Height    int          `json:"height"`
	Obstacles []Coordinate `json:"obstacles"`
}
