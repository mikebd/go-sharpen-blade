package model

type Cards []Card

type Card struct {
	rank Rank
	suit Suit
}

var UnknownCard = Card{UnknownRank, UnknownSuit}
var AceHeart = Card{Ace, Heart}
var TwoHeart = Card{Two, Heart}
var ThreeHeart = Card{Three, Heart}
var FourHeart = Card{Four, Heart}
var FiveHeart = Card{Five, Heart}
var SixHeart = Card{Six, Heart}
var SevenHeart = Card{Seven, Heart}
var EightHeart = Card{Eight, Heart}
var NineHeart = Card{Nine, Heart}
var TenHeart = Card{Ten, Heart}
var JackHeart = Card{Jack, Heart}
var QueenHeart = Card{Queen, Heart}
var KingHeart = Card{King, Heart}
var AceDiamond = Card{Ace, Diamond}
var TwoDiamond = Card{Two, Diamond}
var ThreeDiamond = Card{Three, Diamond}
var FourDiamond = Card{Four, Diamond}
var FiveDiamond = Card{Five, Diamond}
var SixDiamond = Card{Six, Diamond}
var SevenDiamond = Card{Seven, Diamond}
var EightDiamond = Card{Eight, Diamond}
var NineDiamond = Card{Nine, Diamond}
var TenDiamond = Card{Ten, Diamond}
var JackDiamond = Card{Jack, Diamond}
var QueenDiamond = Card{Queen, Diamond}
var KingDiamond = Card{King, Diamond}
var AceClub = Card{Ace, Club}
var TwoClub = Card{Two, Club}
var ThreeClub = Card{Three, Club}
var FourClub = Card{Four, Club}
var FiveClub = Card{Five, Club}
var SixClub = Card{Six, Club}
var SevenClub = Card{Seven, Club}
var EightClub = Card{Eight, Club}
var NineClub = Card{Nine, Club}
var TenClub = Card{Ten, Club}
var JackClub = Card{Jack, Club}
var QueenClub = Card{Queen, Club}
var KingClub = Card{King, Club}
var AceSpade = Card{Ace, Spade}
var TwoSpade = Card{Two, Spade}
var ThreeSpade = Card{Three, Spade}
var FourSpade = Card{Four, Spade}
var FiveSpade = Card{Five, Spade}
var SixSpade = Card{Six, Spade}
var SevenSpade = Card{Seven, Spade}
var EightSpade = Card{Eight, Spade}
var NineSpade = Card{Nine, Spade}
var TenSpade = Card{Ten, Spade}
var JackSpade = Card{Jack, Spade}
var QueenSpade = Card{Queen, Spade}
var KingSpade = Card{King, Spade}

func (c Card) String() string {
	switch c.suit {
	case UnknownSuit:
		switch c.rank {
		case UnknownRank:
			return " ??"
		case Ace:
			return " A?"
		case Two:
			return " 2?"
		case Three:
			return " 3?"
		case Four:
			return " 4?"
		case Five:
			return " 5?"
		case Six:
			return " 6?"
		case Seven:
			return " 7?"
		case Eight:
			return " 8?"
		case Nine:
			return " 9?"
		case Ten:
			return "10?"
		case Jack:
			return " J?"
		case Queen:
			return " Q?"
		case King:
			return " K?"
		}
	case Heart:
		switch c.rank {
		case UnknownRank:
			return " ?♥"
		case Ace:
			return " A♥"
		case Two:
			return " 2♥"
		case Three:
			return " 3♥"
		case Four:
			return " 4♥"
		case Five:
			return " 5♥"
		case Six:
			return " 6♥"
		case Seven:
			return " 7♥"
		case Eight:
			return " 8♥"
		case Nine:
			return " 9♥"
		case Ten:
			return "10♥"
		case Jack:
			return " J♥"
		case Queen:
			return " Q♥"
		case King:
			return " K♥"
		}
	case Diamond:
		switch c.rank {
		case UnknownRank:
			return " ?♦"
		case Ace:
			return " A♦"
		case Two:
			return " 2♦"
		case Three:
			return " 3♦"
		case Four:
			return " 4♦"
		case Five:
			return " 5♦"
		case Six:
			return " 6♦"
		case Seven:
			return " 7♦"
		case Eight:
			return " 8♦"
		case Nine:
			return " 9♦"
		case Ten:
			return "10♦"
		case Jack:
			return " J♦"
		case Queen:
			return " Q♦"
		case King:
			return " K♦"
		}
	case Club:
		switch c.rank {
		case UnknownRank:
			return " ?♣"
		case Ace:
			return " A♣"
		case Two:
			return " 2♣"
		case Three:
			return " 3♣"
		case Four:
			return " 4♣"
		case Five:
			return " 5♣"
		case Six:
			return " 6♣"
		case Seven:
			return " 7♣"
		case Eight:
			return " 8♣"
		case Nine:
			return " 9♣"
		case Ten:
			return "10♣"
		case Jack:
			return " J♣"
		case Queen:
			return " Q♣"
		case King:
			return " K♣"
		}
	case Spade:
		switch c.rank {
		case UnknownRank:
			return " ?♠"
		case Ace:
			return " A♠"
		case Two:
			return " 2♠"
		case Three:
			return " 3♠"
		case Four:
			return " 4♠"
		case Five:
			return " 5♠"
		case Six:
			return " 6♠"
		case Seven:
			return " 7♠"
		case Eight:
			return " 8♠"
		case Nine:
			return " 9♠"
		case Ten:
			return "10♠"
		case Jack:
			return " J♠"
		case Queen:
			return " Q♠"
		case King:
			return " K♠"
		}
	}

	return " ??"
}
