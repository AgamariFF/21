package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type card struct {
	suit    int
	meaning int
}

func SlowPrint(s string, p int, nl bool) {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
		time.Sleep(time.Duration(p) * time.Microsecond)
	}
	if nl {
		fmt.Println()
	}
}

func PrintPlayer(playerCards [36]card) {
	var score, count int
	var chek = false
	SlowPrint("Ваши карты:", 20000, true)
	for i := 0; i < len(playerCards); i++ {
		if playerCards[i].meaning == -1 {
			break
		} else {
			if playerCards[i].meaning < 3 {
				score += playerCards[i].meaning + 2
			} else if playerCards[i].meaning > 2 {
				score += playerCards[i].meaning + 3
			}
			count++
		}
	}
	SlowPrint(strings.Repeat(" ___  ", count), 7000, true)
	var str string
	for i := 0; i < len(playerCards); i++ {
		switch playerCards[i].meaning {
		case -1:
			{
				fmt.Println()
				chek = true
				break
			}
		case 0:
			str = "|В  | "
		case 1:
			str = "|Д  | "
		case 2:
			str = "|К  | "
		case 3:
			str = "|6  | "
		case 4:
			str = "|7  | "
		case 5:
			str = "|8  | "
		case 6:
			str = "|9  | "
		case 7:
			str = "|10 | "
		case 8:
			str = "|Т  | "
		}
		if chek {
			break
		}
		SlowPrint(str, 7000, false)
	}
	chek = false
	var suit string
	for i := 0; i < len(playerCards); i++ {
		switch playerCards[i].suit {
		case -1:
			{
				fmt.Println()
				chek = true
				break
			}
		case 0:
			suit = "♥"
		case 1:
			suit = "♦"
		case 2:
			suit = "♣"
		case 3:
			suit = "♠"
		}
		if chek {
			break
		}
		SlowPrint("| "+suit+" | ", 7000, false)
	}
	chek = false
	for i := 0; i < len(playerCards); i++ {
		switch playerCards[i].meaning {
		case -1:
			{
				fmt.Println()
				chek = true
				break
			}
		case 0:
			str = "|__В| "
		case 1:
			str = "|__Д| "
		case 2:
			str = "|__К| "
		case 3:
			str = "|__6| "
		case 4:
			str = "|__7| "
		case 5:
			str = "|__8| "
		case 6:
			str = "|__9| "
		case 7:
			str = "|_10| "
		case 8:
			str = "|__Т| "
		}
		if chek {
			break
		}
		SlowPrint(str, 7000, false)
	}

	SlowPrint("Сумма очков ", 20000, false)
	fmt.Print(score)
	SlowPrint(".", 20000, true)
}

func GiveCard(user int, counterCard int, deck [36]card, userCards [36]card) (int, int, [36]card, [36]card) {
	if deck[counterCard].meaning < 3 {
		user += deck[counterCard].meaning + 2
	} else if deck[counterCard].meaning > 2 {
		user += deck[counterCard].meaning + 3
	}
	for i := 0; i < 36; i++ {
		if userCards[i].meaning == -1 {
			userCards[i].meaning = deck[counterCard].meaning
			userCards[i].suit = deck[counterCard].suit
			break
		}
	}
	counterCard += 1
	return user, counterCard, deck, userCards
}

func StirDeck() [36]card {
	var chek bool
	var theCard card
	var Deck [36]card
	for i := 0; i < 36; i++ {
		Deck[i].meaning = -1
		Deck[i].suit = -1
	}
	for count := 0; count < 36; {
		chek = false
		theCard.meaning = int(rand.Int31n(9))
		theCard.suit = int(rand.Int31n(4))
		for i := 0; i < 36; i++ {
			if Deck[i] == theCard {
				chek = true
				break
			}
		}
		if chek {
			continue
		}
		Deck[count] = theCard
		count++
	}
	return Deck
}

func play() {
	var take string
	var playerCards, dealerCards [36]card
	for i := 0; i < 36; i++ {
		playerCards[i].meaning = -1
		playerCards[i].suit = -1
	}
	for i := 0; i < 36; i++ {
		dealerCards[i].meaning = -1
		dealerCards[i].suit = -1
	}
	var dealer, player, counterCard int
	deck := StirDeck()
	dealer, counterCard, deck, dealerCards = GiveCard(dealer, counterCard, deck, dealerCards)
	player, counterCard, deck, playerCards = GiveCard(player, counterCard, deck, playerCards)
	PrintPlayer(playerCards)
	SlowPrint("Взять еще карту? (y/n)", 20000, true)
	fmt.Scan(&take)
	for player < 21 && string(take) == "y" {
		player, counterCard, deck, playerCards = GiveCard(player, counterCard, deck, playerCards)
		PrintPlayer(playerCards)
		if player < 21 {
			SlowPrint("Взять еще карту? (y/n)", 20000, true)
			fmt.Scan(&take)
		}
	}
	SlowPrint("Дилер берёт карты.", 20000, true)
	time.Sleep(1000 * time.Millisecond)
	for dealer < 17 {
		dealer, counterCard, deck, dealerCards = GiveCard(dealer, counterCard, deck, dealerCards)
	}
	SlowPrint("У дилера ", 20000, false)
	fmt.Print(dealer)
	if dealer%10 < 1 || dealer%10 > 4 {
		SlowPrint(" очков. ", 20000, false)
	} else if dealer == 21 {
		SlowPrint(" очко. ", 20000, false)
	} else {
		SlowPrint(" очка. ", 20000, false)
	}
	if player > 21 || (dealer >= player && dealer <= 21) {
		SlowPrint("Вы проиграли.", 20000, true)
	} else {
		SlowPrint("Вы выйграли!", 20000, true)
	}
}

func main() {
	var game string = "y"
	for game == string("y") {
		play()
		SlowPrint("Сыграть еще раз? (y/n)", 20000, true)
		fmt.Scan(&game)
	}
}
