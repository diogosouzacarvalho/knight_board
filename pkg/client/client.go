package client

import "github.com/diogosouzacarvalho/knight_board/internal/models"

type GameClient interface {
	GetBoard() (models.Board, error)
	GetCommands() (models.Commands, error)
}
