package game

import (
	"fmt"
	"slices"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
)

func (g *Game) SetStartingPoint(position models.Coordinate, direction string) error {
	if slices.Contains(g.board.Obstacles, position) {
		fmt.Println("Starting position overlaps with obstacle")
		return status.ErrInvalidStartPosition
	}

	if position.X < 0 ||
		position.Y < 0 ||
		position.X+1 > g.board.Width ||
		position.Y+1 > g.board.Height {
		return status.ErrInvalidStartPosition
	}

	g.currentPosition = position
	g.facingDirection = direction

	return nil
}
