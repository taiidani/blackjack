package main

import (
	"math/rand"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	d := Deck{}
	d.Cards = []Card{}

	for s := 0; s < 4; s++ {
		for r := 0; r < 13; r++ {
			d.Cards = append(d.Cards, Card{
				Rank: Rank(r),
				Suit: Suit(s),
			})
		}
	}

	return &d
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
