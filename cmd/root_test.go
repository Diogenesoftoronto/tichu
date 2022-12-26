package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	cmd := RootCmd()
	assert := assert.New(t)
	got := cmd.Use
	want := "tichu - multiplayer card game"
	assert.Equal(want, got)

	got = cmd.Short
	want = "A card game for 4 players."
	assert.Equal(want, got)
	
}
