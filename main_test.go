package main_test

import (
	"testing"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
)

type MockClient struct{}

func (c *MockClient) GetBoard() (models.Board, error) {
	board := models.Board{
		Width:  10,
		Height: 10,
		Obstacles: []models.Coordinate{
			{X: 1, Y: 1},
			{X: 5, Y: 7},
		},
	}

	return board, nil
}

func (c *MockClient) GetCommands() (models.Commands, error) {
	commands := models.Commands{
		Commands: []string{
			"START 1,0,NORTH",
			"ROTATE SOUTH",
			"MOVE 3",
			"ROTATE EAST",
			"MOVE 5",
		},
	}

	return commands, nil
}

func TestMain(t *testing.T) {
	var (
		err      error
	)
	client := &MockClient{}

	if _, err = client.GetBoard(); err != nil {
		t.Fatal("error getting board")
	}
	if _, err = client.GetCommands(); err != nil {
		t.Fatal("error getting commands")
	}
}
