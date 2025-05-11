package game

import (
	"strconv"
	"strings"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
)

func (g *Game) SetStartingPosition(command string) error {
	if !strings.HasPrefix(command, string(models.CommandTypeStart)) {
		return status.ErrGeneric
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

	if position.CollidesWithObstacle(g.board) {
		return status.ErrInvalidStartPosition
	}

	if !position.IsValid(g.board) {
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

		if nextPosition.CollidesWithObstacle(g.board) {
			break
		}

		if !nextPosition.IsValid(g.board) {
			return status.ErrOutOfBoard
		}

		g.currentPosition = nextPosition
	}

	return nil
}

func (g *Game) DoRotate(partialCommand string) error {
	direction := models.Direction(partialCommand)

	g.facingDirection = direction

	return nil
}
