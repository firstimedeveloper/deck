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

// func ReverseSort(cards []Card) []Card {
// 	return sort.Reverse(cards)
// }

func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var newDeck []Card
	randomNum := r.Perm(len(cards))
	for i := 0; i < len(cards); i++ {
		newDeck = append(newDeck, cards[randomNum[i]])
	}
	return newDeck
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(NUMOFRANKS) + int(c.Rank)
}
