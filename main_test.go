package main_test

import (
	"errors"
	"testing"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
	"github.com/diogosouzacarvalho/knight_board/pkg/game"
)

type MockClient struct {
	commands models.Commands
}

func (c *MockClient) GetBoard() (models.Board, error) {
	return models.Board{
		Width:     10,
		Height:    10,
		Obstacles: nil,
	}, nil
}
func (c *MockClient) GetCommands() (models.Commands, error) {
	return c.commands, nil
}
func (c *MockClient) SetCommands(commands models.Commands) {
	c.commands = commands
}

func TestMainOutOfBoard(t *testing.T) {
	client := &MockClient{}
	client.SetCommands(models.Commands{
		Commands: []string{
			"START 1,0,WEST",
			"MOVE 2",
		},
	})
	g := game.NewGame(client)

	if err := g.Init(); err != nil {
		t.Fatalf("expected nil error, got: %s", err)
	}

	err := g.Exec()
	if !errors.Is(status.ErrOutOfBoard, err) {
		t.Fatalf("expected %s, got: %s", status.ErrOutOfBoard, err)
	}
}

func TestMain(t *testing.T) {
	client := &MockClient{}
	client.SetCommands(models.Commands{
		Commands: []string{
			"START 1,0,EAST",
			"MOVE 2",
		},
	})
	g := game.NewGame(client)

	if err := g.Init(); err != nil {
		t.Fatalf("expected nil error, got: %s", err)
	}

	if err := g.Exec(); err != nil {
		t.Fatalf("expected nil error, got: %s", err)
	}

	finalPosition := g.GetCurrentPosition()
	if finalPosition.X != 3 {
		t.Fatalf("expected x: %d, got: %d", 3, finalPosition.X)
	}
	if finalPosition.Y != 0 {
		t.Fatalf("expected y: %d, got: %d", 0, finalPosition.Y)
	}
	finalDirection := g.GetFacingDirection()
	if finalDirection != "EAST" {
		t.Fatalf("expected direction: %s, got: %s", "EAST", finalDirection)
	}
}

func TestMainErrInvalidStartPosition(t *testing.T) {
	client := &MockClient{}
	client.SetCommands(models.Commands{
		Commands: []string{
			"START 11,0,EAST",
			"MOVE 2",
		},
	})
	g := game.NewGame(client)

	if err := g.Init(); err != nil {
		t.Fatalf("expected nil error, got: %s", err)
	}

	if err := g.Exec(); err == nil {
		t.Fatalf("expected %s error, got: nil", status.ErrInvalidStartPosition)
	}
}
