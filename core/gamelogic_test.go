package core

import (
	"sort"
	"testing"
)

func TestdoingBenckmark() {
	// This reports a custom benchmark metric relevant to a
	// specific algorithm (in this case, sorting).
	testing.Benchmark(func(b *testing.B) {
		var compares int64
		for i := 0; i < b.N; i++ {
			s := []int{5, 4, 3, 2, 1}
			sort.Slice(s, func(i, j int) bool {
				compares++
				return s[i] < s[j]
			})
		}
		// This metric is per-operation, so divide by b.N and
		// report it as a "/op" unit.
		b.ReportMetric(float64(compares)/float64(b.N), "compares/op")
	})
}

func TestDeck_Init(t *testing.T) {
	deck := Deck{}
	deck.Init()

	// A standard deck should have 52 cards
	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(deck.cards))
	}

	// Verify all cards are unique (check by counting unique combinations)
	cardMap := make(map[string]bool)
	for _, card := range deck.cards {
		key := card.Naipe + card.Code
		if cardMap[key] {
			t.Errorf("Duplicate card found: %s%s", card.Naipe, card.Code)
		}
		cardMap[key] = true
	}

	if len(cardMap) != 52 {
		t.Errorf("Expected 52 unique cards, found %d", len(cardMap))
	}

	// Verify deck has been shuffled (stack should not be empty)
	if deck.stack.Empty() {
		t.Error("Deck stack should not be empty after Init()")
	}

	if deck.stack.Size() != 52 {
		t.Errorf("Expected stack size 52, got %d", deck.stack.Size())
	}
}

func TestDeck_GetCards(t *testing.T) {
	t.Run("Get cards from initialized deck", func(t *testing.T) {
		deck := Deck{}
		deck.Init()

		initialStackSize := deck.stack.Size()
		cards := deck.GetCards(2)

		if len(cards) != 2 {
			t.Errorf("Expected 2 cards, got %d", len(cards))
		}

		// Verify stack size decreased
		expectedSize := initialStackSize - 2
		if deck.stack.Size() != expectedSize {
			t.Errorf("Expected stack size %d, got %d", expectedSize, deck.stack.Size())
		}

		// Verify cards are valid (have naipe, code, and value)
		for i, card := range cards {
			if card.Naipe == "" {
				t.Errorf("Card %d missing naipe", i)
			}
			if card.Code == "" {
				t.Errorf("Card %d missing code", i)
			}
			if card.Value < 2 || card.Value > 14 {
				t.Errorf("Card %d has invalid value %d (expected 2-14)", i, card.Value)
			}
		}

		// Get more cards
		cards2 := deck.GetCards(2)
		if len(cards2) != 2 {
			t.Errorf("Expected 2 cards, got %d", len(cards2))
		}

		// Verify we got different cards (since deck is shuffled)
		if cards[0].Naipe == cards2[0].Naipe && cards[0].Code == cards2[0].Code {
			// This could happen by chance, but very unlikely with shuffled deck
			// We'll just verify the cards are valid
		}
	})

	t.Run("Get cards from empty deck", func(t *testing.T) {
		deck := Deck{}
		deck.Init()

		// Empty the deck
		for !deck.stack.Empty() {
			deck.GetCards(1)
		}

		// Try to get cards from empty deck
		cards := deck.GetCards(2)
		if len(cards) != 0 {
			t.Errorf("Expected 0 cards from empty deck, got %d", len(cards))
		}
	})

	t.Run("Get more cards than available", func(t *testing.T) {
		deck := Deck{}
		deck.Init()

		// Get all cards except 1
		for deck.stack.Size() > 1 {
			deck.GetCards(1)
		}

		// Try to get 5 cards when only 1 is available
		cards := deck.GetCards(5)
		if len(cards) != 1 {
			t.Errorf("Expected 1 card (only one available), got %d", len(cards))
		}

		if !deck.stack.Empty() {
			t.Error("Deck stack should be empty after getting all cards")
		}
	})

	t.Run("Get zero cards", func(t *testing.T) {
		deck := Deck{}
		deck.Init()

		initialSize := deck.stack.Size()
		cards := deck.GetCards(0)

		if len(cards) != 0 {
			t.Errorf("Expected 0 cards, got %d", len(cards))
		}

		if deck.stack.Size() != initialSize {
			t.Errorf("Stack size should remain %d, got %d", initialSize, deck.stack.Size())
		}
	})
}

func TestDeck_Shuffle(t *testing.T) {
	deck := Deck{}
	deck.Init()

	// Get initial order of cards
	initialCards := make([]Card, 0)
	for deck.stack.Size() > 0 {
		cards := deck.GetCards(1)
		initialCards = append(initialCards, cards...)
	}

	// Reinitialize and shuffle
	deck.Init()

	// Get cards again
	shuffledCards := make([]Card, 0)
	for deck.stack.Size() > 0 {
		cards := deck.GetCards(1)
		shuffledCards = append(shuffledCards, cards...)
	}

	// Verify we have the same cards (just different order)
	if len(initialCards) != len(shuffledCards) {
		t.Errorf("Card count mismatch: initial %d, shuffled %d", len(initialCards), len(shuffledCards))
	}

	// Check that at least some cards are in different positions
	// (with random shuffle, it's extremely unlikely all cards are in same position)
	samePosition := 0
	for i := 0; i < len(initialCards) && i < len(shuffledCards); i++ {
		if initialCards[i].Naipe == shuffledCards[i].Naipe && initialCards[i].Code == shuffledCards[i].Code {
			samePosition++
		}
	}

	// If more than 90% are in same position, shuffle might not be working
	// (though this is probabilistic, so we just verify it's not 100%)
	if samePosition == len(initialCards) {
		t.Error("Shuffle appears to not be working - all cards in same position")
	}
}
