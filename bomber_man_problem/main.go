package main

import (
	"fmt"
)

func main() {
	initialState, N := processInput()
	outputState(generateState(initialState, N))
}

func generateState(initialState []string, N int) []string {
	if N == 1 {
		return initialState
	} else if N % 2 == 0 {
		C := len(initialState[0])
		R := len(initialState)
		tmp := make([]string, R)
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				tmp[r] += "O"
			}
		}
		return tmp
	} else if N % 4 == 3 {
		return bombed(initialState)
	} else {
		return bombed(bombed(initialState))
	}
}

func bombed(initialState []string) []string {
	C := len(initialState[0])
	R := len(initialState)
	state := make([]string, R)

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if isDetonate(r, c, initialState) {
				state[r] += "."
			} else {
				state[r] += "O"
			}
		}
	}
	return state
}

func isDetonate(r int, c int, state []string) bool {
	directions := [][]int{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}

	C := len(state[0])
	R := len(state)

	if state[r][c] == 'O' {
		return true
	}

	flag := false
	for _, direction := range directions {
		dc := c + direction[0]
		dr := r + direction[1]

		if dc < 0 || dc >= C || dr < 0 || dr >= R {
			continue
		}

		if state[dr][dc] == 'O' {
			flag = true
			break
		}
	}
	return flag
}

func outputState(state []string)  {
	for _, row := range state {
		fmt.Println(row)
	}
}

func processInput() ([]string, int) {
	var R, C, N int
	fmt.Scanf("%d", &R)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &N)
	state := make([]string, R)
	for r := 0; r < R; r++ {
		fmt.Scanf("%s", &state[r])
	}

	return state, N
}