package game

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
)

func (g *Game) SetStartingPosition(command string) error {
	if !strings.HasPrefix(command, string(models.CommandTypeStart)) {
		return errors.New("command should be start")
	}

	temp := strings.Split(command, " ")
	values := strings.Split(temp[1], ",")
	x, err := strconv.ParseInt(values[0], 10, 16)
	if err != nil {
		return err
	}
	y, err := strconv.ParseInt(values[1], 10, 16)
	if err != nil {
		return err
	}
	position := models.Coordinate{
		X: int(x),
		Y: int(y),
	}
	direction := values[2]

	if slices.Contains(g.board.Obstacles, position) {
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
