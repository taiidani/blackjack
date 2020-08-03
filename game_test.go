package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	rand.Seed(1)
	g := NewGame(1)
	expectHouse := &Hand{
		Name:  "House",
		In:    true,
		Cards: []Card{{Rank: rankKing, Suit: suitHeart}},
		Bank:  []Card{{Rank: rankKing, Suit: suitDiamond}},
	}

	expectP1 := &Hand{
		Name:  "P1",
		In:    true,
		Cards: []Card{{Rank: rankAce, Suit: suitHeart}, {Rank: rankFive, Suit: suitDiamond}},
		Bank:  []Card{},
	}

	if len(g.deck.Cards) != 48 {
		t.Errorf("NewGame() len = %d, want %d", len(g.deck.Cards), 48)
	} else if !reflect.DeepEqual(g.house, expectHouse) {
		t.Errorf("NewGame() house = %v, want %v", g.house, expectHouse)
	} else if !reflect.DeepEqual(g.players[0], expectP1) {
		t.Errorf("NewGame() p1 = %v, want %v", g.players[0], expectP1)
	}
}

func TestHit(t *testing.T) {
	rand.Seed(1)
	g := NewGame(1)
	expectP1 := &Hand{
		Name: "P1",
		In:   true,
		Cards: []Card{
			{Rank: rankAce, Suit: suitHeart},
			{Rank: rankFive, Suit: suitDiamond},
			{Rank: rankAce, Suit: suitSpade},
		},
		Bank: []Card{},
	}

	g.Hit(g.players[0])

	if len(g.deck.Cards) != 47 {
		t.Errorf("Hit() len = %d, want %d", len(g.deck.Cards), 47)
	} else if g.players[0].Blackjack() {
		t.Errorf("Hit() blackjack = %t, want %t", g.players[0].Blackjack(), false)
	} else if g.players[0].Bust() {
		t.Errorf("Hit() bust = %t, want %t", g.players[0].Bust(), false)
	} else if !reflect.DeepEqual(g.players[0], expectP1) {
		t.Errorf("Hit() p1 = %v, want %v", g.players[0], expectP1)
	}
}
