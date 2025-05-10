package game_test

import (
	"testing"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
	"github.com/diogosouzacarvalho/knight_board/pkg/game"
)

func TestStartPosition(t *testing.T) {
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
	}
	client.SetCommands(commands)
	g := game.NewGame(client)

	if err := g.Init(); err != nil {
		t.Fatalf("got error on init: %s", err)
	}
	if err := g.SetStartingPosition(commands[0]); err != nil {
		t.Fatalf("got error on setting start position: %s", err)
	}

	client.SetCommands([]string{
		"MOVE 1",
	})
	g = game.NewGame(client)

	if err := g.Init(); err == nil {
		t.Fatalf("got nil error on init, expected %s", status.ErrGeneric)
	}
	if err := g.SetStartingPosition("MOVE 1"); err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.SetCommands([]string{
		"START -1,0,NORTH",
	})
	g = game.NewGame(client)

	if err := g.Init(); err != nil {
		t.Fatalf("got error on init: %s", err)
	}
	if err := g.SetStartingPosition("START -1,0,NORTH"); err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.SetBoard(models.Board{
		Width:  10,
		Height: 10,
		Obstacles: []models.Coordinate{
			{X: 1, Y: 1},
		},
	})
	client.SetCommands([]string{
		"START 1,1,NORTH",
	})
	g = game.NewGame(client)

	if err := g.Init(); err != nil {
		t.Fatalf("got error on init: %s", err)
	}
	if err := g.SetStartingPosition("START 1,1,NORTH"); err == nil {
		t.Fatalf("expected error, got nil")
	}
}
