package main

import "fmt"

type Game struct {
	deck    *Deck
	house   *Hand
	players []*Hand
}

func NewGame(players int) *Game {
	g := &Game{}
	g.deck = NewDeck()
	g.deck.Shuffle()

	g.house = NewHand("House")
	for i := 0; i < players; i++ {
		g.players = append(g.players, NewHand(fmt.Sprintf("P%d", i+1)))
	}

	g.Deal()
	return g
}

func (g *Game) Deal() {
	all := append([]*Hand{g.house}, g.players...)

	// Deal everyone one card faceup
	for p, player := range all {
		player.Cards = append(player.Cards, g.deck.Cards[p])
	}
	g.deck.Cards = g.deck.Cards[len(all):]

	// Deal all players a second faceup card
	for p, player := range g.players {
		player.Cards = append(player.Cards, g.deck.Cards[p])
	}
	g.deck.Cards = g.deck.Cards[len(g.players):]

	// House's second card is facedown
	g.house.Bank = append(g.house.Bank, g.deck.Cards[0])
	g.deck.Cards = g.deck.Cards[1:]
}

func (g *Game) Hit(player *Hand) Card {
	newCard := g.deck.Cards[0]
	player.Cards = append(player.Cards, newCard)
	g.deck.Cards = g.deck.Cards[1:]
	return newCard
}

func (g *Game) Winners() []*Hand {
	winners := []*Hand{}
	toBeat := g.house.HighestValue()

	for _, player := range g.players {
		if player.Bust() {
			continue
		} else if g.house.Bust() ||
			player.Blackjack() ||
			player.HighestValue() > toBeat {
			winners = append(winners, player)
		}
	}

	return winners
}
