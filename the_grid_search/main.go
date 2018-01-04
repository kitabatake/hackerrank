package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		G, P := processInput()
		find := find(G, P)
		if find {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func find(G []string, P []string) bool {
	pc := len(P[0])
	pr := len(P)
	gc := len(G[0])
	gr := len(G)

	for r := pr - 1; r < gr; r++ {
		for c := pc - 1; c < gc; c++ {
			target := extractPattern(G, pr, pc, c, r)
			if isSamePattern(target, P) {
				return true
			}
		}
	}
	return false
}

func isSamePattern(p1 []string, p2 []string) bool {
	r := len(p1)
	for i := 0; i < r; i++ {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}

func extractPattern(G []string, r int, c int, right int, bottom int) []string {
	pattern := make([]string, r)
	for i := 0; i < r; i++ {
		pattern[i] = G[bottom - r + 1 + i][right - c + 1:right + 1]
	}
	return pattern
}

func processInput() ([]string, []string) {
	var r, c int
	fmt.Scanf("%d", &r)
	fmt.Scanf("%d", &c)
	G := make([]string, r)
	for i := 0; i < r; i ++ {
		fmt.Scanf("%s", &G[i])
	}

	fmt.Scanf("%d", &r)
	fmt.Scanf("%d", &c)
	P := make([]string, r)
	for i := 0; i < r; i ++ {
		fmt.Scanf("%s", &P[i])
	}
	return G, P
}