package cmd

import (
	"regexp"
	"testing"

	"github.com/Diogenesoftoronto/tichu/deck"
	"github.com/google/uuid"
)

func TestDeckCmd(t *testing.T) {
	// Create and get an deck and idf
	// then use the Deckcmd to retrieve that deck from the db
	cmd := DeckCmd()
	cmd.Flags().Bool("json", false, "Send output as json")
	// // Call the base deck command to view the id
	c, got, err := ExecuteCommandC(cmd, "--json", "b97d90e2-6a1d-4a14-872e-6270aea1be49")
	if err != nil {
		t.Errorf("Command %v failed to execute, got %v,", c, got)
	}
	// Make sure that what i got is a deck
	if _, err := deck.ParseJson(got); err != nil {
		t.Errorf("Parsing the result of %v failed. with %v error.", got, err)
	}
}

func TestDeckCreateCmd(t *testing.T) {
	// call the command
	cmd := DeckCreateCmd()
	// this is necessary because you can only check for a uuid as a type not the actual uuid!
	// I will use the MustParse function in the uuid library to check whether something is a uuid.
	// first I must use regular expressions to exclude everything except for the uuid.
	cmd.Flags().IntP("amount", "a", 1, "The amount of decks to create")
	cmd, got, err := ExecuteCommandC(cmd, "--amount", "2", "1")
	if err != nil {
		t.Errorf("got %v from command %v with error: %v", got, cmd, err)
	}
	// we need to parse got to exclude the parts we don't need.
	r, err := regexp.Compile(`.+-.+`)
	if err != nil {
		t.Errorf("failed to compile regular expression")
	}
	// found is the string we retrieve using the regular expression
	found := r.FindAllString(got, -1)
	defer func() {
		// recover returns an error if panic occurs so we handle the err in the case it is not nil
		if err := recover(); err != nil {
			t.Errorf("Found could not be parsed as a UUID, %s, the error is %v, got is %v", found, err, got)
		}
	}()
	// this function panics if this is wrong so we need to handle the recovery
	uuid.MustParse(found[0])
	uuid.MustParse(found[1])
	// if there is no panic then the test passes 👼
}
