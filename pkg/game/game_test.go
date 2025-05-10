package game_test

import (
	"testing"

	"github.com/diogosouzacarvalho/knight_board/pkg/client"
	"github.com/diogosouzacarvalho/knight_board/pkg/game"
)

func TestGameInit(t *testing.T) {
	client := &client.MockClient{}
	game := game.NewGame(client)

	if err := game.Init(); err != nil {
		t.Fatalf("got error on init: %s", err)
	}
}
