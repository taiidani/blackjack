package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(3)
	play(g)
}

func play(g *Game) {
	r := 1
	for {
		fmt.Println("Round", r)
		printPlayers(append([]*Hand{g.house}, g.players...)...)

		hit := round(g)
		if !hit {
			break
		}
		r++
		time.Sleep(time.Second)
	}

	fmt.Println("House Round")
	// Flip the facedown card
	g.house.Cards = append(g.house.Cards, g.house.Bank...)
	g.house.Bank = []Card{}

	printPlayers(append([]*Hand{g.house}, g.players...)...)
	houseRound(g)
	time.Sleep(time.Second)

	// Check game over
	fmt.Println("Game over")
	printPlayers(append([]*Hand{g.house}, g.players...)...)
	winners := g.Winners()
	if len(winners) == 0 {
		fmt.Println("All players lost!")
	} else {
		fmt.Println("We've got winners!")
		for _, player := range winners {
			fmt.Println(player.Name)
		}
	}
}

func round(g *Game) (hit bool) {
	for _, player := range g.players {
		if !player.In {
			continue
		}

		fmt.Printf("%s...", player.Name)
		time.Sleep(time.Second)
		if player.Blackjack() {
			fmt.Println("Blackjack!")
			player.In = false
			continue
		} else if high := player.HighestValue(); high > 17 {
			fmt.Printf("STAY. ")
		} else {
			fmt.Print("HIT ")
			hit = true
			card := g.Hit(player)
			fmt.Print("for " + card.String() + ". ")
		}

		if player.Bust() {
			fmt.Printf("Bust with %d!\n", player.LowestValue())
			player.In = false
		} else if player.Blackjack() {
			fmt.Println("Blackjack!")
			player.In = false
		} else {
			fmt.Println()
		}
	}
	fmt.Println()
	return
}

func houseRound(g *Game) {
	// Keep hittin'
	for g.house.HighestValue() < 16 {
		fmt.Print("House...")
		time.Sleep(time.Second)
		fmt.Print("HIT ")
		card := g.Hit(g.house)
		fmt.Println("for " + card.String() + ". ")
	}

	if g.house.Bust() {
		fmt.Printf("Bust with %d!\n", g.house.LowestValue())
	} else if g.house.Blackjack() {
		fmt.Println("Blackjack!")
	} else {
		fmt.Println()
	}

	fmt.Println()
}

func printPlayers(players ...*Hand) {
	for _, p := range players {
		if !p.In {
			continue
		}
		fmt.Println(p)
	}
	fmt.Println("---")
}
