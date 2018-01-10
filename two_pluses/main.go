package main

import (
	"fmt"
)

func main() {
	G := processInput()
	fmt.Println(dfs(G, 0, []int{}))
}

func printGrid(G [][]rune)  {
	for r := 0; r < len(G); r++ {
		fmt.Println(string(G[r]))
	}
}

func dfs(G [][]rune, depth int, plusAreas []int) int {
	if depth == 2 {
		return plusAreas[0] * plusAreas[1]
	}

	max := 0
	for r := 0; r < len(G); r++ {
		for c := 0; c < len(G[0]); c++ {
			plusLengths := getPlusLengths(G, r, c)
			for _, plusLength := range plusLengths {
				if plusLength > 0 {
					areas := dfs(turnBad(G, r, c, plusLength), depth + 1, append(plusAreas, (plusLength * 2 - 1)))
					if areas > max {
						max = areas
					}
				}
			}
		}
	}
	return max
}

func getPlusLengths(G [][]rune, r int, c int) []int {
	if G[r][c] == 'B' {
		return []int{}
	}

	directions := [][]int {
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	ret := []int {1}
	current := 1
	loop:
	for {
		for i := 0; i < 4; i++ {
			cc := c + (directions[i][0] * current)
			cr := r + (directions[i][1] * current)
			if cc < 0 || cc >= len(G[0]) || cr < 0 || cr >= len(G) {
				break loop
			}
			if G[cr][cc] == 'B' {
				break loop
			}
		}
		current++
		ret = append(ret, 2 * (current - 1) + 1)
	}
	return ret
}

func turnBad(G [][]rune, r int, c int, length int) [][]rune {
	ret := make([][]rune, len(G))
	for i := 0; i < len(G); i++ {
		ret[i] = make([]rune, len(G[0]))
		for j := 0; j < len(G[0]); j++ {
			ret[i][j] = G[i][j]
		}
	}

	ret[r][c] = 'B'
	directions := [][]int {
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	eachLength := length/2
	for _, direction := range directions {
		for l := 1; l <= eachLength; l++ {
			cr := r + (direction[1] * l)
			cc := c + (direction[0] * l)
			ret[cr][cc] = 'B'
		}
	}
	return ret
}

func processInput() [][]rune {
	var R, C int
	fmt.Scanf("%d", &R)
	fmt.Scanf("%d", &C)

	G := make([][]rune, R)
	var s string
	for r := 0; r < R; r++ {
		fmt.Scanf("%s", &s)
		G[r] = []rune(s)
	}
	return G
}
