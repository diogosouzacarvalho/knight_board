package game

import (
	"fmt"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/pkg/client"
)

type Game struct {
	client       client.GameClient
	board        models.Board
	baseCommands models.Commands
}

func NewGame(client client.GameClient) *Game {
	return &Game{client: client}
}

func (g *Game) Init() error {
	board, err := g.client.GetBoard()
	if err != nil {
		return err
	}
	g.board = board

	commands, err := g.client.GetCommands()
	if err != nil {
		return err
	}
	g.baseCommands = commands

	fmt.Printf("\nBoard: %+v", g.board)
	fmt.Printf("\nCommands: %+v", g.baseCommands)

	return nil
}
