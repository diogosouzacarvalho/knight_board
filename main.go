package main

import (
	"os"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/pkg/client"
	"github.com/diogosouzacarvalho/knight_board/pkg/game"
)

func main() {
	client := client.NewStorageClient(os.Getenv(models.BOARD_API), os.Getenv(models.COMMANDS_API))
	game := game.NewGame(client)

	if err := game.Init(); err != nil {
		panic(err)
	}
}
