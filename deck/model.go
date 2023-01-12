package deck

import (
	"github.com/google/uuid"
)

type Face int8

const (
	// This is the faces of the cards
	DOG Face = iota - 1
	MAHJONG
	ACE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	PHEONIX
	DRAGON
)

type Suit int8

const (
	// This is the suits of the cards
	CHARACTERS Suit = iota - 1
	CIRCLES
	BAMBOO
	WINDS
)

type Value int8

const (
	// This is the value of the cards
	V_DOG     Value = 0
	V_FIVE    Value = 5
	V_TEN     Value = 10
	V_KING    Value = 10
	V_PHEONIX Value = -25
	V_DRAGON  Value = 25
)

// This is a struct that represents a deck of cards
type Deck struct {
	Id    uuid.UUID   `json:id`
	Cards []Deck_item `json:cards`
	Type  Deck_type   `json:type`
}

type Deck_type int

const (
	// This is the types of decks
	Unset Deck_type = iota - 1
	Standard
	Server
	User
	Game
	Message
)

// This is a struct that represents a deck item
type Deck_item struct {
	Card
}

// This is a struct that represents a card
type Card struct {
	Value Value
	Face  Face
	Suit  Suit
}

/* PARSE FUNCTIONS

These functions exist to turn normal integers into enums, limited types.
This provides type guarantees.*/
