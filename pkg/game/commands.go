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
	direction := models.Direction(values[2])

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

func (g *Game) DoMove(partialCommand string) error {
	targetValue, err := strconv.ParseInt(partialCommand, 10, 16)
	if err != nil {
		return err
	}
	for i := 0; i < int(targetValue); i++ {
		nextPosition := g.currentPosition
		switch g.facingDirection {
		case models.DirectionEast:
			nextPosition.X++
		case models.DirectionWest:
			nextPosition.X--
		case models.DirectionNorth:
			nextPosition.Y++
		case models.DirectionSouth:
			nextPosition.Y--
		default:
			return status.ErrGeneric
		}

		if slices.Contains(g.board.Obstacles, nextPosition) {
			break
		}

		if nextPosition.X < 0 || nextPosition.Y < 0 || nextPosition.X+1 > g.board.Width ||
			nextPosition.Y+1 > g.board.Height {
			return status.ErrOutOfBoard
		}

		g.currentPosition = nextPosition
	}

	return nil
}
