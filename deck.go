//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	NUMOFSUITS = 4
	NUMOFRANKS = 13
)

//Rank is of type int, used to enumerate the 13 Ranks
type Rank int

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

//Suit is of type int, this is used to enumerate the 4 types of suites
type Suit int

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
	//Red //Diamond, Heart //TODO: add a mechanism for filtering out a color
	//Black //Spade, Club
)

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

//Card is a struct that has 2 values, S which is of type Suite, and Num which is of type int
type Card struct {
	Suit
	Rank
}

type SortOption func(*SortOptions)
type SortOptions struct{}

//New returns a deck of cards which is an array of the struct Card
func New(opts ...func([]Card) []Card) []Card {
	var deck []Card
	for suit := 0; suit < NUMOFSUITS; suit++ {
		for num := 1; num <= NUMOFRANKS; num++ {
			deck = append(deck, Card{
				Suit(suit),
				Rank(num),
			})
		}
	}

	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck
}

//MultipleDecks makes a single deck composed of multiple decks.
func MultipleDecks(n int) func([]Card) []Card {

	return func(deck []Card) []Card {
		for i := 1; i < n; i++ {
			for _, c := range deck {
				deck = append(deck, c)
			}
		}
		return deck
	}
}

// func ReverseSort(cards []Card) []Card {
// 	return sort.Reverse(cards)
// }

//FilterRank filters out specific Ranks.
//Both int and the name of the rank can be used
//Ex Ace or 1, 11 or Jack etc. can be used.
func FilterRank(rank ...Rank) func([]Card) []Card {
	return func(card []Card) []Card {
		i := 0
		for _, r := range rank { //one
			i = 0
			for _, c := range card {
				if r != c.Rank {
					card[i] = c
					i++
				}
			}
			card = card[:i]
		}
		return card[:i]
	}
}

//FilterSuit filters out specific suits.
//Both int and the name of the suit can be used
//However, the name of the suit should be used to avoid confusion
//Ex Both FilterSuit(Heart) and FilterSuit(3) removes all Heart cards
func FilterSuit(suit ...Suit) func([]Card) []Card {
	return func(card []Card) []Card {
		i := 0
		for _, s := range suit {
			i = 0
			for _, c := range card {
				if s != c.Suit {
					card[i] = c
					i++
				}
			}
			card = card[:i]
		}
		return card[:i]
	}
}

//FilterSpecific filters out specific cards
//ex FilterSpecific(Card{Heart, Ace}) filters out Ace of Hearts
func FilterSpecific(card ...Card) func([]Card) []Card {
	return func(deck []Card) []Card {
		i := 0
		for _, c := range card {
			i = 0
			for _, d := range deck {
				if c != d {
					deck[i] = d
					i++
				}
			}
			deck = deck[:i]
		}
		return deck[:i]
	}
}

//Jokers is a function that adds an arbritary number of Joker cards
//to the end of the deck.
//If You want to a shuffled deck, make sure to add Jokers(n) before Shuffle
//as a parameter to the New() function.
func Jokers(n int) func([]Card) []Card {

	return func(c []Card) []Card {
		for i := 0; i < n; i++ {
			c = append(c, Card{
				Suit: Joker,
			})
		}
		return c
	}
}

//Shuffle method shuffles the deck ([]Card) in a random order.
//The order will be different everytime its run.
func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var newDeck []Card
	randomNum := r.Perm(len(cards))
	for i := 0; i < len(cards); i++ {
		newDeck = append(newDeck, cards[randomNum[i]])
	}
	return newDeck
}

//DefaultSort is the default way most just-bought cards come in.
//The order is Spade, Diamond, Club, and Heart.
//Ex. Spade Ace-King, Diamond Ace-King ... in that order.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Less function compares the absRank of two cards, returns true if it's in order
//(Meaning the first card is smaller than the second one)
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(NUMOFRANKS) + int(c.Rank)
}
