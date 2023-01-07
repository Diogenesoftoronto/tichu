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
	V_DOG     Value = 0
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

/* PARSE FUNCTIONS

These functions exist to turn normal integers into enums, limited types.
This provides type guarantees.*/

func ParseType(deck_int int) (Deck_type, error) {
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

func ParseSuit(suit_int int) (Suit, error) {
	switch suit_int {
	case 0:
		return CHARACTERS, nil
	case 1:
		return CIRCLES, nil
	case 2:
		return BAMBOO, nil
	case 3:
		return WINDS, nil
	default:
		return -1, errors.New("invalid deck suit")
	}

}

func ParseValue(value_int int) (Value, error) {
	switch value_int {
	case 0:
		return V_DOG, nil
	case 5:
		return V_FIVE, nil
	case 10:
		return V_TEN, nil
	case 25:
		return V_DRAGON, nil
	case -25:
		return V_PHEONIX, nil
	default:
		return -1, errors.New("invalid deck value")
	}

}

func ParseFace(face_int int) (Face, error) {
	switch face_int {
	case 0:
		return DOG, nil
	case 1:
		return MAHJONG, nil
	case 2:
		return ACE, nil
	case 3:
		return TWO, nil
	case 4:
		return THREE, nil
	case 5:
		return FOUR, nil
	case 6:
		return FIVE, nil
	case 7:
		return SIX, nil
	case 8:
		return SEVEN, nil
	case 9:
		return EIGHT, nil
	case 10:
		return NINE, nil
	case 11:
		return TEN, nil
	case 12:
		return JACK, nil
	case 13:
		return QUEEN, nil
	case 14:
		return KING, nil
	case 15:
		return PHEONIX, nil
	case 16:
		return DRAGON, nil
	default:
		return -1, errors.New("invalid deck face")
	}

}
