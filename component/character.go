package component

import "math/rand"

type char struct {
	Clan
	Clue Clan
	Rank uint32
}

type Character *char

type Clan string

const (
	BLUE_CLAN Clan = "b"
	RED_CLAN  Clan = "r"
	SEC_CLAN  Clan = "s"
)

var (
	B1 = chary(BLUE_CLAN, BLUE_CLAN, 1) // elder
	B2 = chary(BLUE_CLAN, BLUE_CLAN, 2) // assassin
	B3 = chary(BLUE_CLAN, RED_CLAN, 3)  // harlequin
	B4 = chary(BLUE_CLAN, BLUE_CLAN, 4) // alchemist
	B5 = chary(BLUE_CLAN, BLUE_CLAN, 5) // mentalist
	B6 = chary(BLUE_CLAN, BLUE_CLAN, 6) // guardian
	B7 = chary(BLUE_CLAN, BLUE_CLAN, 7) // berserker
	B8 = chary(BLUE_CLAN, BLUE_CLAN, 8) // mage
	B9 = chary(BLUE_CLAN, BLUE_CLAN, 9) // courtesan

	R1 = chary(RED_CLAN, RED_CLAN, 1)  // elder
	R2 = chary(RED_CLAN, RED_CLAN, 2)  // assassin
	R3 = chary(RED_CLAN, BLUE_CLAN, 3) // harlequin
	R4 = chary(RED_CLAN, RED_CLAN, 4)  // alchemist
	R5 = chary(RED_CLAN, RED_CLAN, 5)  // mentalist
	R6 = chary(RED_CLAN, RED_CLAN, 6)  // guardian
	R7 = chary(RED_CLAN, RED_CLAN, 7)  // berserker
	R8 = chary(RED_CLAN, RED_CLAN, 8)  // mage
	R9 = chary(RED_CLAN, RED_CLAN, 9)  // courtesan

	BS = chary(SEC_CLAN, BLUE_CLAN, 0) // blue clue inquisitor
	RS = chary(SEC_CLAN, RED_CLAN, 0)  // red clue inquisitor

	CharMap = map[Clan][]Character{
		BLUE_CLAN: {B1, B2, B3, B4, B5, B6, B7, B8, B9},
		RED_CLAN:  {R1, R2, R3, R4, R5, R6, R7, R8, R9},
		SEC_CLAN:  {BS, RS},
	}
)

func chary(clan Clan, clue Clan, rank uint32) Character {
	return Character(&char{Clan: clan, Clue: clue, Rank: rank})
}

func RandCharN(r *rand.Rand, clan Clan, n int) []Character {
	base := CharMap[clan]
	picked := r.Perm(len(base))[:n]

	result := make([]Character, 0, n)
	for _, i := range picked {
		result = append(result, base[i])
	}
	return result
}
