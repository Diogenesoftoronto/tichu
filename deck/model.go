package deck

import (
	"errors"

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
	V_FIVE    Value = 5
	V_TEN     Value = 10
	V_KING    Value = 10
	V_PHEONIX Value = -25
	V_DRAGON  Value = 25
)

// This is a struct that represents a deck of cards
type Deck struct {
	Id    uuid.UUID
	Cards []Deck_item
	Type  Deck_type
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

func ParseDeck_type(deck_int int) (Deck_type, error) {
	switch deck_int {
	case 0:
		return Unset, nil
	case 1:
		return Standard, nil
	case 2:
		return Server, nil
	case 3:
		return User, nil
	case 4:
		return Game, nil
	case 5:
		return Message, nil

	default:
		return -1, errors.New("invalid deck type")
	}

}
