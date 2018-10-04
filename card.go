//go:generate stringer -type=Suit,Rank

package deck

import (
	"time"
	"math/rand"
	"sort"
	"fmt"
)

//Suit gives each suit an integer value
type Suit uint8

const(
	//Spade is the first suit given an integer value of 0
	Spade Suit = iota
	//Diamond has an integer value of 3
	Diamond
	//Club has an integer value of 2
	Club
	//Heart has an integer value of 1
	Heart
	//Joker has an int value of 4
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

//Rank gives each card a value such as Jack, Queen, 2 etc.
type Rank uint8
//Gives each Rank an integer value iota increments by +1.
const(
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

const(
	minRank = Ace
	maxRank = King
)

//Card assigns a value and suit to each card
type Card struct{
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker{
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
//NewDeck sets up the deck in order from Ace to King going Spades, Diamond, Clubs, Hearts.
func NewDeck(opts ...func([]Card) []Card) []Card{
	var cards []Card
	for _, suit := range suits{
		for rank := minRank; rank <= maxRank; rank++{
			cards = append(cards, Card{Suit:suit, Rank:rank})
		}
	}
	for _, opt := range opts{
		cards = opt(cards)
	}
	return cards
}

func defaultSort(cards []Card) []Card{
	sort.Slice(cards, Less(cards))
	return cards
}

func customSort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card{
	return func(cards []Card) []Card{
		sort.Slice(cards, less(cards))
		return cards
	}
}
//Less sorts cards based on there rank and suit, going spades, diamonds, clubs, hearts.
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool{
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int{
	return int(c.Suit)*int(maxRank)+int(c.Rank)
}
//Shuffle randomly rearranges the cards
func Shuffle(cards []Card) []Card{
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm{
	ret[i] = cards[j]	
	}
	return ret
}
//Jokers adds jokers into the game if needed
func Jokers(n int) func([]Card) []Card{
	return func(cards []Card) []Card{
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}
//Filter allows the removal of certain ranks from the deck.
func Filter(f func(card Card)bool) func([]Card) []Card {
	return func(cards []Card) []Card{
		var ret []Card
		for _, c := range cards{
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}
//MultiDeck allows for more than one deck of cards at once
func MultiDeck(n int) func([]Card) []Card{
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}