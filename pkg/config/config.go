package config

import (
	"encoding/json"
	"net/http"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
)

var sampleBoard models.Board = models.Board{
	Width:     10,
	Height:    10,
	Obstacles: nil,
}

func GetBoard(endpoint string) (models.Board, error) {
	if endpoint == "" {
		return sampleBoard, nil
	}

	var board models.Board

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return board, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return board, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&board); err != nil {
		return board, err
	}

	return board, nil
}

var sampleCommand models.Commands = models.Commands{
	Commands: []string{
		"START 1,0,NORTH",
		"ROTATE SOUTH",
		"MOVE 3",
		"ROTATE EAST",
		"MOVE 5",
	},
}

func GetCommands(endpoint string) (models.Commands, error) {
	if endpoint == "" {
		return sampleCommand, nil
	}

	var commands models.Commands

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return commands, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return commands, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&commands); err != nil {
		return commands, err
	}

	return commands, nil
}
