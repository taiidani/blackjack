package main

import (
	"math/rand"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d.Cards) != 52 {
		t.Errorf("NewDeck() len = %d, want %d", len(d.Cards), 51)
	}
}
func TestDeck_Shuffle(t *testing.T) {
	rand.Seed(1)
	d := NewDeck()
	d.Shuffle()

	if len(d.Cards) != 52 {
		t.Errorf("Shuffle() len = %d, want %d", len(d.Cards), 52)
	} else if d.Cards[0].Value(false) != 10 {
		t.Errorf("Shuffle() [0] = %d, want %d", d.Cards[0].Value(false), 10)
	} else if d.Cards[10].Value(false) != 10 {
		t.Errorf("Shuffle() [10] = %d, want %d", d.Cards[10].Value(false), 10)
	} else if d.Cards[20].Value(false) != 1 {
		t.Errorf("Shuffle() [20] = %d, want %d", d.Cards[20].Value(false), 1)
	} else if d.Cards[30].Value(false) != 10 {
		t.Errorf("Shuffle() [30] = %d, want %d", d.Cards[30].Value(false), 10)
	} else if d.Cards[40].Value(false) != 10 {
		t.Errorf("Shuffle() [40] = %d, want %d", d.Cards[40].Value(false), 10)
	} else if d.Cards[50].Value(false) != 9 {
		t.Errorf("Shuffle() [50] = %d, want %d", d.Cards[50].Value(false), 9)
	}
}
