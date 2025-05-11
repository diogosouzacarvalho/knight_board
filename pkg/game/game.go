package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
	"github.com/diogosouzacarvalho/knight_board/pkg/client"
)

type Game struct {
	client          client.GameClient
	board           models.Board
	baseCommands    models.Commands
	currentPosition models.Coordinate
	facingDirection models.Direction
}

func NewGame(client client.GameClient) *Game {
	return &Game{client: client}
}

func (g *Game) Init() error {
	var (
		board models.Board
		commands models.Commands
		err error
	)
	if board, err = g.client.GetBoard(); err != nil {
		return status.ErrGeneric
	}
	if !board.IsValid() {
		return status.ErrGeneric
	}
	g.board = board

	if commands, err = g.client.GetCommands(); err != nil {
		return status.ErrGeneric
	}
	if !commands.IsValid() {
		return status.ErrGeneric
	}
	g.baseCommands = commands

	return nil
}

func (g *Game) Exec() error {
	var err error
	commands := g.baseCommands.Commands

	startCommand := commands[0]
	if err = g.SetStartingPosition(startCommand); err != nil {
		fmt.Printf("invalid start position: %s", err)
		if !errors.Is(err, status.ErrInvalidStartPosition) {
			err = status.ErrGeneric
		}
		return err
	}

	doCommands := commands[1:]

	for len(doCommands) > 0 {
		currentCommand := doCommands[0]
		temp := strings.Split(currentCommand, " ")
		commandType := models.CommandType(temp[0])

		switch commandType {
		case models.CommandTypeMove:
			err = g.DoMove(temp[1])
		case models.CommandTypeRotate:
			err = g.DoRotate(temp[1])
		default:
			err = status.ErrGeneric
		}
		if err != nil {
			return err
		}

		// if len(doCommands) == 1 {
		// 	break
		// }
		doCommands = doCommands[1:]
	}

	return nil
}

func (g *Game) GetCurrentPosition() models.Coordinate {
	return g.currentPosition
}

func (g *Game) GetFacingDirection() models.Direction {
	return g.facingDirection
}
