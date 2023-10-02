package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Ten, Suit: Diamond})
	fmt.Println(Card{Rank: Seven, Suit: Spade})
	fmt.Println(Card{Rank: Three, Suit: Heart})
	fmt.Println(Card{Suit: Joker})
	//Output:
	//Ace of Hearts
	//Two of Clubs
	//Ten of Diamonds
	//Seven of Spades
	//Three of Hearts
	//Joker
}

func TestNew(t *testing.T) {
	cards := New()
	//13 ranks * 4 suits
	if len(cards) != 52 {
		t.Error("Wrong number of cards in the deck")	
	}
}