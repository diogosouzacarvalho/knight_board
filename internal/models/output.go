package models

type OutputPosition struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Direction string `json:"direction"`
}

type Output struct {
	Position OutputPosition `json:"position,omitempty"`
	Status   string         `json:"status"`
}
