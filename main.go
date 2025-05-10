package main

import (
	"fmt"
	"os"

	"github.com/diogosouzacarvalho/knight_board/internal/models"
	"github.com/diogosouzacarvalho/knight_board/pkg/client"
	"github.com/diogosouzacarvalho/knight_board/pkg/game"
)

type OutputPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
	Direction string `json:"direction"`
}

type output struct {
	Position OutputPosition `json:"position,omitempty"`
	Status string `json:"status"`
}

func main() {
	client := client.NewStorageClient(os.Getenv(models.BOARD_API), os.Getenv(models.COMMANDS_API))
	g := game.NewGame(client)

	var out output


	if err := g.Init(); err != nil {
		out = output{Status: err.Error()}
		fmt.Printf("%+v", out)
		return
	}

	if err := g.Exec(); err != nil {
		out = output{Status: err.Error()}
	} else {
		out = output{
			Position: OutputPosition{
				X: g.GetCurrentPosition().X,
				Y: g.GetCurrentPosition().Y,
				Direction: string(g.GetFacingDirection()),
			},
			Status: "SUCCESS",
		}
	}

	fmt.Printf("%+v", out)
}
