package deck

import "fmt"

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
