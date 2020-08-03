package main

import (
	"fmt"
	"strings"
)

type Hand struct {
	Name  string
	In    bool // If the player is still in the game
	Cards []Card
	Bank  []Card
}

func NewHand(name string) *Hand {
	return &Hand{
		Name:  name,
		In:    true,
		Cards: []Card{},
		Bank:  []Card{},
	}
}

func (h *Hand) String() string {
	return fmt.Sprintf("%s (%d): Cards: %s", h.Name, h.HighestValue(), h.StringVisible())
}

func (h *Hand) StringVisible() string {
	table := []string{}
	for _, card := range h.Cards {
		table = append(table, card.String())
	}

	return strings.Join(table, ", ")
}

func (h *Hand) Bust() bool {
	low := h.LowestValue()
	return low > 21
}

func (h *Hand) Blackjack() bool {
	high, low := h.TableValue()
	return low == 21 || high == 21
}

func (h *Hand) HighestValue() int {
	high, low := h.TableValue()
	if high > 21 {
		return low
	}
	return high
}

func (h *Hand) LowestValue() int {
	_, low := h.TableValue()
	return low
}

func (h *Hand) TableValue() (int, int) {
	high := 0
	low := 0
	for _, card := range h.Cards {
		low += card.Value(false)
		high += card.Value(true)
	}
	return high, low
}
