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
	facingDirection string
}

func NewGame(client client.GameClient) *Game {
	return &Game{client: client}
}

func (g *Game) Init() error {
	board, err := g.client.GetBoard()
	if err != nil {
		return err
	}
	if board.Height < 1 || board.Width < 1 {
		fmt.Println("invalid board")
		return status.ErrGeneric
	}
	g.board = board

	commands, err := g.client.GetCommands()
	if err != nil {
		return err
	}
	if len(commands.Commands) < 1 {
		fmt.Println("invalid commands")
		return status.ErrGeneric
	}
	if !strings.HasPrefix(commands.Commands[0], string(models.CommandTypeStart)) {
		fmt.Println("first command should be start")
		return status.ErrGeneric
	}
	g.baseCommands = commands

	return nil
}

func (g *Game) Exec() error {
	commands := g.baseCommands.Commands

	startCommand := commands[0]
	if err := g.SetStartingPosition(startCommand); err != nil {
		fmt.Printf("invalid start position: %s", err)
		if !errors.Is(err, status.ErrInvalidStartPosition) {
			err = status.ErrGeneric
		}
		return err
	}

	return nil
}
