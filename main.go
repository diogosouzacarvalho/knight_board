package main

import (
	"fmt"
	"os"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/pkg/client"
)

func main() {
	var (
		board    models.Board
		commands models.Commands
		err      error
	)
	client := client.NewStorageClient(os.Getenv(models.BOARD_API), os.Getenv(models.COMMANDS_API))

	if board, err = client.GetBoard(); err != nil {
		panic(err)
	}
	if commands, err = client.GetCommands(); err != nil {
		panic(err)
	}

	fmt.Println(board)
	fmt.Println(commands)
}
