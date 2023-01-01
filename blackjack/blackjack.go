package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	score := ParseCard(card1) + ParseCard(card2)
	dealerHand := ParseCard(dealerCard)

	var decision string
	switch {
	case score < 12:
		decision = smallHand(score, dealerHand)

	case score >= 12 && score <= 20:
		decision = mediumHand(score, dealerHand)

	case score > 20:
		decision = largeHand(score, dealerHand)

	}

	return decision
}

// smallHand returns the decision for scores below 12.
func smallHand(score int, dealerCard int) string {
	return "H"
}

// mediumHand returns the decision for scores 12 to 20.
func mediumHand(score int, dealerCard int) string {
	var decision string

	switch score {
	case 17, 18, 19, 20:
		decision = "S"
	case 12, 13, 14, 15, 16:
		if dealerCard >= 7 {
			decision = "H"
		} else {
			decision = "S"
		}
	}

	return decision
}

// largeHand returns the decision for scores above 20.
func largeHand(score int, dealerCard int) string {
	var decision string

	switch score {
	case 22:
		decision = "P"
	case 21:
		if dealerCard != 11 && dealerCard != 10 {
			decision = "W"
		} else {
			decision = "S"
		}
	}

	return decision
}
