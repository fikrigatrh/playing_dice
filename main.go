package main

import (
	"fmt"
	"math/rand"
)

func main() {

	numPlayers := 0
	dice := 0

	fmt.Print("masukan jumlah pemain: ")
	fmt.Scanf("%d", &numPlayers)
	fmt.Print("")
	fmt.Print("masukan jumlah dadu: ")
	fmt.Scanf("%d", &dice)

	playerDice := make([][]int, numPlayers)
	playerPoints := make([]int, numPlayers)
	playerActive := make([]bool, numPlayers)

	for i := 0; i < numPlayers; i++ {
		playerDice[i] = rollDice(dice)
		playerActive[i] = true
	}

	turn := 1

	for {
		fmt.Printf("Giliran %d lempar dadu:\n", turn)

		for i := 0; i < numPlayers; i++ {
			if playerActive[i] {
				fmt.Printf("Pemain #%d (%d): %v\n", i, playerPoints[i], playerDice[i])
			}
		}

		var oneDice []int
		var point int
		fmt.Println("Setelah evaluasi:")
		for i := 0; i < numPlayers; i++ {
			if playerActive[i] {
				point, playerDice[i], oneDice = evaluateDice(playerDice[i])
				playerPoints[i] += point
				if len(oneDice) > 0 {
					if i == numPlayers-1 && len(playerDice[i]) != 0 {
						playerDice[0] = append(playerDice[0], oneDice...)
					}

					j := 0
					j = i + 1
					if j >= numPlayers {
						continue
					}
					i = j
					if len(playerDice[i]) != 0 {
						playerDice[i] = append(playerDice[i], oneDice...)
					}
				}

			}
		}
		for i := 0; i < numPlayers; i++ {
			fmt.Printf("Pemain #%d (%d): %v\n", i, playerPoints[i], playerDice[i])
		}

		//var playerInput int
		for i := 0; i < numPlayers; i++ {
			if playerActive[i] {
				fmt.Printf("Pemain #%d, tekan Enter untuk melanjutkan lemparan...", i)
				playerDice[i] = rollDice(len(playerDice[i]))
				fmt.Scanln()
			}
		}

		activeCount := 0
		activePlayer := -1
		score := playerPoints[0]
		winner := 0
		for i := 0; i < numPlayers; i++ {
			if len(playerDice[i]) != 0 {
				activeCount++
				activePlayer = i
			}
			if playerPoints[i] > score {
				score = playerPoints[i]
				winner = i
			}
		}

		if activeCount == 1 {
			fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", activePlayer)
			fmt.Printf("Game dimenangkan oleh pemain  #%d  karena memiliki poin lebih banyak dari pemain lainnya..\n", winner)
			break
		}

		turn++
	}
}

func rollDice(numDice int) []int {
	dice := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		dice[i] = rand.Intn(6) + 1
	}
	return dice
}

func evaluateDice(dice []int) (points int, newDice []int, oneDice []int) {
	for _, value := range dice {
		if value == 6 {
			points++
			continue
		}
		if value == 1 {
			oneDice = append(oneDice, value)
		}
		if value != 1 {
			newDice = append(newDice, value)
		}
	}
	return
}
