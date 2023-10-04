package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Ten, Suit: Diamond})
	fmt.Println(Card{Rank: Seven, Suit: Spade})
	fmt.Println(Card{Rank: Three, Suit: Heart})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Ten of Diamonds
	// Seven of Spades
	// Three of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 52 {
		t.Error("Wrong number of cards in the deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("expected Ace of Spades as first card, received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("expected Ace of Spades as first card, received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	numJokers := 3
	cards := New(Jokers(numJokers))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}
	if count != numJokers {
		t.Errorf("Expected %d Jokers, got %d", numJokers, count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}	
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all Twos and Threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	numDecks := 3
	cards := New(Deck(numDecks))	
	if len(cards) != 13*4*numDecks {
		t.Errorf("Expected %d cards, got %d", 13*4*numDecks, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected first card to be %v, got %v", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected second card to be %v, got %v", second, cards[1])
	}


}