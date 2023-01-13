package deck

import (
	"encoding/json"
	"errors"
	// "fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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

func Parse(s string) (*Deck, error) {
	var result *Deck
	// split the string into many string
	str_array := strings.Split(s, " ")
	// each split is an element of the struct
	// parse each element into their respective types
	id, err := uuid.Parse(str_array[0])
	if err != nil {
		return nil, err
	}
	var numsconv []int
	for _, str := range str_array {
		number, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		numsconv = append(numsconv, number)
	}
	if err != nil {
		return nil, err
	}
	Type, err := ParseType(numsconv[0])
	if err != nil {
		return nil, err
	}

	value, err := ParseValue(numsconv[2])
	if err != nil {
		return nil, err
	}
	face, err := ParseFace(numsconv[3])
	if err != nil {
		return nil, err
	}
	suit, err := ParseSuit(numsconv[4])
	if err != nil {
		return nil, err
	}

	result = &Deck{
		Id: id,
		Cards: []Deck_item{
			{Card{
				Value: value,
				Face:  face,
				Suit:  suit,
			},
			},
		},
		Type: Type,
	}

	return result, nil
}

// type ServiceID struct {
//     UUID uuid.UUID
// }

// type Meta struct {
//     Name    string `json:"name"`
//     Version string `json:"version"`
//     SID     *ServiceID `json:"UUID"`
// }

// func (self *ServiceID) UnmarshalJSON(b []byte) error {
    // s := strings.Trim(string(b), "\"")
//     self.UUID = uuid.Parse(s)
//     if self.UUID == nil {
//             return errors.New("Could not parse UUID")
//     }
//     return nil
// }


func ParseJson(s string) (*Deck, error) {
	var result *Deck
	// unmarshall a simple struct

	// type simpDeck struct {
	// 	id string `json:"id"`
	// 	Cards []Deck_item 	`json:"cards"`
	// 	Type Deck_type	`json:"type"`
	// }
	// take the id of that and convert it to a uuid

	// take that 
	// unmarshall json from an array of bytes into struct!
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		return nil, err
	}
	return result, nil
}
