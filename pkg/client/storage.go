package client

import (
	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/utils"
)

type StorageClient struct {
	boardEndpoint    string
	commandsEndpoint string
}

func NewStorageClient(boardEndpoint, commandsEndpoint string) *StorageClient {
	return &StorageClient{
		boardEndpoint:    boardEndpoint,
		commandsEndpoint: commandsEndpoint,
	}
}

func (sc *StorageClient) GetBoard() (models.Board, error) {
	var board models.Board

	return utils.GetData(sc.boardEndpoint, board)
}

func (sc *StorageClient) GetCommands() (models.Commands, error) {
	var commands models.Commands

	return utils.GetData(sc.commandsEndpoint, commands)
}
