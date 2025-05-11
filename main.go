package main

import (
	"os"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/internal/status"
	"github.com/diogosouzacarvalho/knight_board/internal/utils"
	"github.com/diogosouzacarvalho/knight_board/pkg/client"
	"github.com/diogosouzacarvalho/knight_board/pkg/game"
)

func main() {
	client := client.NewStorageClient(os.Getenv(models.BOARD_API), os.Getenv(models.COMMANDS_API))
	g := game.NewGame(client)

	if err := g.Init(); err != nil {
		out := models.Output{Status: err.Error()}
		utils.OutputJSON(out)
		return
	}

	if err := g.Exec(); err != nil {
		out := models.Output{Status: err.Error()}
		utils.OutputJSON(out)
		return
	}

	out := models.Output{
		Position: models.OutputPosition{
			X:         g.GetCurrentPosition().X,
			Y:         g.GetCurrentPosition().Y,
			Direction: string(g.GetFacingDirection()),
		},
		Status: status.StatusSuccess,
	}

	utils.OutputJSON(out)
}
