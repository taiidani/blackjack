package main

import "fmt"

type (
	Suit int
	Rank int

	Card struct {
		Suit Suit
		Rank Rank
	}
)

const (
	suitClub Suit = iota
	suitSpade
	suitHeart
	suitDiamond
)

const (
	rankAce Rank = iota
	rankTwo
	rankThree
	rankFour
	rankFive
	rankSix
	rankSeven
	rankEight
	rankNine
	rankTen
	rankJack
	rankQueen
	rankKing
)

func (s Suit) String() string {
	return [...]string{"Clubs", "Spades", "Hearts", "Diamonds"}[s]
}

func (r Rank) String() string {
	return [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}[r]
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

func (c Card) Value(aceHigh bool) int {
	switch c.Rank {
	case rankAce:
		if aceHigh {
			return 11
		}
		return 1
	case rankJack, rankQueen, rankKing:
		return 10
	default:
		return int(c.Rank) + 1
	}
}
