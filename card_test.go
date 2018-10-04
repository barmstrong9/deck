package deck

import (
	"testing"
	"fmt"
)

func ExampleCard(){
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Rank: Eight, Suit: Club})
	fmt.Println(Card{Rank: Ten, Suit: Spade})
	fmt.Println(Card{Suit: Joker})
	//Output:
	//Ace of Hearts
	//Jack of Diamonds
	//Eight of Clubs
	//Ten of Spades
	//Joker

}

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 13*4{
		t.Error("Wrong Number of Cards")
	}
}