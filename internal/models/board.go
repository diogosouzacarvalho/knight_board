package models

type Board struct {
	Width     int          `json:"width"`
	Height    int          `json:"height"`
	Obstacles []Coordinate `json:"obstacles"`
}

func (b *Board) IsValid() bool {
	return b.Height > 1 && b.Width > 1
}
