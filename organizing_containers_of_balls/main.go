// Approach:
//	This problem's importance things is whether each container's capacity can fit each ball type counts or not.
//  Container's status and swap operation are'nt important.
//  Hence to focus container's capacities and each ball type counts.

package main

import "fmt"

func main() {
	queries := processInput()

	for _, query := range queries {
		capacities := calcContainersCapacities(query)
		ballCounts := calcBallCounts(query)

		//fmt.Println("query is ", query)
		//fmt.Println("capacities is ", capacities)
		//fmt.Println("ballCounts is ", ballCounts)
		fit := canFit(capacities, ballCounts)
		if fit {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}

func canFit(capacities []int64, ballCounts []int64) bool {
	if len(capacities) == 0 {
		return true
	}

	capacityIndex := -1
	ballCountIndex := -1

	loop:
	for i, capacity := range capacities {
		for j, ballCount := range ballCounts {
			if capacity == ballCount {
				capacityIndex = i
				ballCountIndex = j
				break loop
			}
		}
	}

	if capacityIndex != -1 {
		return canFit(delete(capacities, capacityIndex), delete(ballCounts, ballCountIndex))
	} else {
		return false
	}
}

func calcBallCounts(query [][]int64) []int64 {
	n := len(query)
	counts := make([]int64, n)
	for c := 0; c < n; c ++ {
		count := int64(0)
		for r := 0; r < n; r++ {
			count += query[r][c]
		}
		counts[c] += count
	}
	return counts
}

func calcContainersCapacities(query [][]int64) []int64 {
	n := len(query)
	capacities := make([]int64, n)
	for r := 0; r < n; r ++ {
		cap := int64(0)
		for c := 0; c < n; c++ {
			cap += query[r][c]
		}
		capacities[r] = cap
	}
	return capacities
}

func processInput() [][][]int64 {
	var n, q int
	fmt.Scanf("%d", &q)

	var queries [][][]int64
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &n)
		query := make([][]int64, n)
		for r := 0; r < n; r++ {
			query[r] = make([]int64, n)
			for c := 0; c < n; c++ {
				fmt.Scanf("%d", &query[r][c])
			}
		}
		queries = append(queries, query)
	}
	return queries
}

func delete(s []int64, i int) []int64 {
	s = append(s[:i], s[i+1:]...)
	n := make([]int64, len(s))
	copy(n, s)
	return n
}