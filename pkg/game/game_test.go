package game_test

import (
	"errors"
	"testing"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
	"github.com/diogosouzacarvalho/knight_board/pkg/game"
)

type MockClient struct {
	board    models.Board
	commands models.Commands
}

func (c *MockClient) GetBoard() (models.Board, error) {
	return c.board, nil
}

func (c *MockClient) GetCommands() (models.Commands, error) {
	return c.commands, nil
}

func (c *MockClient) SetBoard(board models.Board) {
	c.board = board
}

func (c *MockClient) SetCommands(commands []string) {
	c.commands.Commands = commands
}

func TestGameInit(t *testing.T) {
	client := &MockClient{}
	board := models.Board{
		Width:  10,
		Height: 10,
		Obstacles: []models.Coordinate{
			{X: 1, Y: 1},
			{X: 5, Y: 7},
		},
	}
	client.SetBoard(board)
	commands := []string{
		"START 1,0,NORTH",
		"ROTATE SOUTH",
		"MOVE 3",
		"ROTATE EAST",
		"MOVE 5",
	}
	client.SetCommands(commands)
	game := game.NewGame(client)

	if err := game.Init(); err != nil {
		t.Fatalf("got error on init: %s", err)
	}
}

func TestGameInitInvalidBoard(t *testing.T) {
	client := &MockClient{}
	board := models.Board{
		Width:     0,
		Height:    0,
		Obstacles: nil,
	}
	client.SetBoard(board)
	commands := []string{
		"START 1,0,NORTH",
	}
	client.SetCommands(commands)
	game := game.NewGame(client)
	err := game.Init()

	if !errors.Is(status.ErrGeneric, err) {
		t.Fatalf("expected generic error, got: %s", err)
	}
}

func TestGameInitInvalidCommands(t *testing.T) {
	client := &MockClient{}
	board := models.Board{
		Width:     10,
		Height:    10,
		Obstacles: nil,
	}
	client.SetBoard(board)
	commands := []string{}
	client.SetCommands(commands)
	g := game.NewGame(client)
	err := g.Init()

	if !errors.Is(status.ErrGeneric, err) {
		t.Fatalf("expected generic error, got: %s", err)
	}

	commands = []string{
		"MOVE 3,NORTH",
	}
	client.SetCommands(commands)
	g = game.NewGame(client)
	err = g.Init()

	if !errors.Is(status.ErrGeneric, err) {
		t.Fatalf("expected generic error, got: %s", err)
	}
}
