package main

import (
	"fmt"
	"os"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/pkg/config"
)

func main() {
	board, err := config.GetBoard(os.Getenv(models.BOARD_API))
	if err != nil {
		panic(err)
	}
	fmt.Println(board)

	commands, err := config.GetCommands(os.Getenv(models.COMMANDS_API))
	if err != nil {
		panic(err)
	}

	fmt.Println(commands)
}
