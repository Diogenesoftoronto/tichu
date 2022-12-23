package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd() (t *testing.T) {
	cmd := RootCmd()
	assert.Equal(t, "tichu - multiplayer card game", cmd.Use)
	assert.Equal(t, "A card game for 4 players.", cmd.Short)
	assert.Equal(t, `A card game for 4 players.
	tichu play begins a tichu game with 4 players.

	tichu cheer cheers you on to victory.

	tichu help lists all commands.

	tichu end ends the game.

	tichu leaderboard lists the players and the score.

	tichu message sends a message to the other players.`, cmd.Long)
	assert.Equal(t, "Welcome to Tichu!\n", cmd.RunE)
	return nil
}
