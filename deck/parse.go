package deck

import (
	"errors"
	// "strings"

	// "github.com/google/uuid"
)

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

// TODO: Consider json unmarshalling instead
// func Parse(s string) (*Deck, error) {
// 	var result *Deck
// 	// split the string into many string
// 	str_array := strings.Split(s, " ")
// 	// each split is an element of the struct
// 	// parse each element into their respective types
// 	id, err :=  uuid.Parse(str_array[0]); if err != nil {
// 		return result, err
// 	}

	
// 	result = &Deck{
// 		Id: id,
// 		Deck_type: 0,
// 		Deck_item: Card{}, 
// 	}

// 	return result, nil
// }