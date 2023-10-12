package model

import "math/rand"

type StandardDeck struct {
	cards Cards
}

func NewStandardDeck() *StandardDeck {
	d := StandardDeck{}
	d.cards = []Card{
		AceHeart,
		TwoHeart,
		ThreeHeart,
		FourHeart,
		FiveHeart,
		SixHeart,
		SevenHeart,
		EightHeart,
		NineHeart,
		TenHeart,
		JackHeart,
		QueenHeart,
		KingHeart,
		AceDiamond,
		TwoDiamond,
		ThreeDiamond,
		FourDiamond,
		FiveDiamond,
		SixDiamond,
		SevenDiamond,
		EightDiamond,
		NineDiamond,
		TenDiamond,
		JackDiamond,
		QueenDiamond,
		KingDiamond,
		AceClub,
		TwoClub,
		ThreeClub,
		FourClub,
		FiveClub,
		SixClub,
		SevenClub,
		EightClub,
		NineClub,
		TenClub,
		JackClub,
		QueenClub,
		KingClub,
		AceSpade,
		TwoSpade,
		ThreeSpade,
		FourSpade,
		FiveSpade,
		SixSpade,
		SevenSpade,
		EightSpade,
		NineSpade,
		TenSpade,
		JackSpade,
		QueenSpade,
		KingSpade,
	}
	d.Shuffle()
	return &d
}

func (d *StandardDeck) Cards() *Cards {
	return &d.cards
}

func (d *StandardDeck) Shuffle() {
	rand.Shuffle(52, func(i, j int) {
		temp := d.cards[i]
		d.cards[i] = d.cards[j]
		d.cards[j] = temp
	})
}
