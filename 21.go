package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func PrintPlayer(playerCards [9]int8) {
	var score int8 = 0
	SlowPrint("У вас:", 20000, false)
	for i := 0; i < len(playerCards); i++ {
		if playerCards[i] == 1 {
			SlowPrint(" ", 20000, false)
			fmt.Print(playerCards[i])
			switch i {
			case 0:
				score += 2
				SlowPrint(" валет", 20000, false)
			case 1:
				score += 3
				SlowPrint(" дама", 20000, false)
			case 2:
				score += 4
				SlowPrint(" король", 20000, false)
			case 3:
				score += 6
				SlowPrint(" шестерка", 20000, false)
			case 4:
				score += 7
				SlowPrint(" семерка", 20000, false)
			case 5:
				score += 8
				SlowPrint(" восьмерка", 20000, false)
			case 6:
				score += 9
				SlowPrint(" девятка", 20000, false)
			case 7:
				score += 10
				SlowPrint(" десятка", 20000, false)
			case 8:
				score += 11
				SlowPrint(" туз", 20000, false)
			}
		} else if playerCards[i] > 1 {
			SlowPrint(" ", 20000, false)
			fmt.Print(playerCards[i])
			switch i {
			case 0:
				score += int8(2 * playerCards[i])
				SlowPrint(" вальта", 20000, false)
			case 1:
				score += int8(3 * playerCards[i])
				SlowPrint(" дамы", 20000, false)
			case 2:
				score += int8(4 * playerCards[i])
				SlowPrint(" короля", 20000, false)
			case 3:
				score += int8(6 * playerCards[i])
				SlowPrint(" шестерки", 20000, false)
			case 4:
				score += int8(7 * playerCards[i])
				SlowPrint(" семерки", 20000, false)
			case 5:
				score += int8(8 * playerCards[i])
				SlowPrint(" восьмерки", 20000, false)
			case 6:
				score += int8(9 * playerCards[i])
				SlowPrint(" девятки", 20000, false)
			case 7:
				score += int8(10 * playerCards[i])
				SlowPrint(" десятки", 20000, false)
			case 8:
				score += int8(11 * playerCards[i])
				SlowPrint(" туза", 20000, false)
			}
		}
	}
	SlowPrint(". Сумма очков ", 20000, false)
	fmt.Print(score)
	SlowPrint(".", 20000, true)
}

func GiveCard(user int8, counterCard int8, deck [36]int8, userCards [9]int8) (int8, int8, [36]int8, [9]int8) {
	if deck[counterCard] == -1 {
		panic("В колоде закончились карты!")
	} else if deck[counterCard] < 3 {
		user += deck[counterCard] + 2
	} else if deck[counterCard] > 2 {
		user += deck[counterCard] + 3
	}
	userCards[deck[counterCard]] += 1
	deck[counterCard] = -1
	counterCard += 1
	return user, counterCard, deck, userCards
}

func StirDeck() [36]int8 {
	var CountDeck [9]int8
	var Deck [36]int8
	var card int8
	for count := 0; count < 36; {
		card = int8(rand.Int31n(9))
		for CountDeck[card] > 3 {
			card = int8(rand.Int31n(9))
		}
		CountDeck[card] += 1
		Deck[count] = card
		count++
	}
	return Deck
}

func play() {
	var take string
	var playerCards, dealerCards [9]int8
	var dealer, player, counterCard int8
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
