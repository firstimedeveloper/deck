package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Five, Suit: Club})
	fmt.Println(Card{Rank: Ace, Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Five of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(Shuffle)
	//exp := Card{Rank: Ace, Suit: Spade}
	// if cards[0] != exp {
	// 	t.Error("Expected Ace of Spades as first card. Recieved: ", cards[0])
	// }
	for _, card := range cards {
		fmt.Println(card)
	}

}
