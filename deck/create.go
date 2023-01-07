package deck

import (
	"github.com/google/uuid"
)

// create a new deck
func New(t Deck_type) *Deck {
	return &Deck{
		Id:    uuid.New(),
		Cards: []Deck_item{},
		Type:  t,
	}
}

// create a new deck item
func NewItem(f Face, s Suit, v Value) Deck_item {
	return Deck_item{
		Card{Face: f,
			Suit:  s,
			Value: v,
		},
	}
}
