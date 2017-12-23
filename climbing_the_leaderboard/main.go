// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
package main

import "fmt"

func main() {
	playerScores , aliceScores := inputProcess()
	currentRank := len(playerScores) + 1
	currentIndex := len(playerScores) - 1

	for _, score := range aliceScores {
		// calculate current score's rank and next score's index
		currentRank, currentIndex = calculateRankAndNextIndex(playerScores, score, currentIndex, currentRank)
		fmt.Println(currentRank)
	}
}

func calculateRankAndNextIndex(scores []int, score int, index int, rank int) (int, int) {
	if index < 0 {
		return 1, -1
	} else if score < scores[index] {
		return rank, index
	} else {
		return calculateRankAndNextIndex(scores, score, index - 1, rank - 1)
	}
}

func inputProcess() (playerScores []int, aliceScores []int) {
	var playersCount, levelsCount, score int
	fmt.Scanf("%d", &playersCount)

	prevScore := 1000000001
	for i := 0; i < playersCount; i++ {
		fmt.Scanf("%d", &score)
		if score != prevScore {
			playerScores = append(playerScores, score)
			prevScore = score
		}
	}

	fmt.Scanf("%d", &levelsCount)
	aliceScores = make([]int, levelsCount)
	for i := 0; i < levelsCount; i++ {
		fmt.Scanf("%d", &aliceScores[i])
	}

	return
}