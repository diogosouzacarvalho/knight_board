package models

import "slices"

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c *Coordinate) IsValid(board Board) bool {
	return (c.X >= 0 && c.Y >= 0) &&
		(c.X < board.Width && c.Y < board.Height)
}

func (c *Coordinate) CollidesWithObstacle(b Board) bool {
	return slices.Contains(b.Obstacles, *c)
}
