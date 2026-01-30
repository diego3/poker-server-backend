package core

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	Naipe string `json:"naipe"`
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type Deck struct {
	stack Stack
	cards []Card
}

type Player struct {
	Uuid   string
	Name   string
	Wins   int
	Money  float64
	Active bool
	Raised bool
	Folded bool
	Called bool
}

func (p *Player) call() {
	em := EventManager{}
	em.trigger(EVENT_CALL, p)
}

const EVENT_GAME_INIT string = "GAME_INIT"
const EVENT_FOLD string = "PLAYER_FOLD"
const EVENT_CALL string = "PLAYER_CALL"
const EVENT_RAISE string = "PLAYER_RAISE"
const EVENT_WINNER string = "PLAYER_WIN_A_HAND"

type Game struct {
	Players               []Player
	Deck                  Deck
	MaxNumberOfPlayers    int
	MaxSecondsDelayToPlay int
	EventManager          EventManager
}

func (g *Game) newGame() {

}

/*
	TODO: Create an object that will coordinate the match (the various poker hands)
	- must have at least two players
	- which player will have the Button?
	- which player will be the Small Blind and Big Blind
	- what is the initial bet value? (small and big blind values) 50/100, 100, 200 etc
	-

*/

func (pm *Deck) Init() {
	naipes := []string{"C", "D", "H", "S"} // Suit initials in English, TODO: research legends
	for _, v := range naipes {
		for i := 2; i <= 10; i++ {
			pm.cards = append(pm.cards, Card{v, strconv.Itoa(i), i})
		}
	}

	faceCards := []string{"J", "Q", "K", "A"}
	for _, v := range naipes {
		corVal := 10
		for _, faceCard := range faceCards {
			// J = 11, Q = 12, K = 13, A = 14
			corVal++
			pm.cards = append(pm.cards, Card{v, faceCard, corVal})
		}
	}
	fmt.Printf("Deck initialized len: %d\n", len(pm.cards))

	pm.Shuffle()
}

func (pm *Deck) Shuffle() {
	var numbers []int
	for i := 1; i <= 52; i++ {
		numbers = append(numbers, i)
	}
	// array shuffle https://programming.guide/go/shuffle-slice-array.html
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	// add shuffled array to stack
	for _, val := range numbers {
		pm.stack.Push(val)
	}
}

// get qty of cards from the cards
func (pm *Deck) GetCards(qty int) []Card {
	var cards []Card
	if pm.stack.Empty() {
		return cards
	}

	for i := 0; i < qty; i++ {
		if pm.stack.Empty() {
			break
		}
		position, _ := pm.stack.Pop()
		cards = append(cards, pm.cards[position-1])
	}
	return cards
}

func (pm *Deck) Print() {
	for _, v := range pm.cards {
		//fmt.Printf("{n:%s c:%s val:%d}\n", v.naipe, v.code, v.value)
		fmt.Println(v)
	}
}
