package cmd

import (
	"testing"
)

func TestDeckCmd(t *testing.T) {
	cmd := DeckCmd()
	assert := func(want, got interface{}) {
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	}
	got := cmd.Use
	want := "deck [commands] [flags] [args]"
	assert(want, got)

	got = cmd.Short
	want = "Manage the cards in a deck"
	assert(want, got)

}

func TestDeckCreateCmd(t *testing.T) {
	cmd := DeckCreateCmd()
	assert := func(want, got interface{}) {
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	}
	got := cmd.Use
	want := "create [flags] ...<arg>"
	assert(want, got)

	got = cmd.Short
	want = "Create a new deck"
	assert(want, got)

	// if err := cmd.RunE(cmd, []string{"1"}); err != nil {
	// 	t.Errorf("running the command returned an error: %v", err)
	// }
}
