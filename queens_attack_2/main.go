package main

import (
	"fmt"
)

func main() {
	n, queens, obstables := processInput()
	obstablesMap := createObstablesMap(obstables)
	fmt.Println(calculateAttackablePositionCounts(n, queens, obstablesMap))
}

func calculateAttackablePositionCounts(n int, queens position, obstablesMap map[string]bool) int {
	directions := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	counts := 0
	for _, direction := range directions {
		tmp := position{
			x: queens.x + direction[0],
			y: queens.y + direction[1],
		}
		for tmp.x >= 1 && tmp.x <= n && tmp.y >= 1 && tmp.y <= n {
			if _, ok := obstablesMap[tmp.string()]; ok {
				break
			}
			tmp.x += direction[0]
			tmp.y += direction[1]
			counts++
		}
	}
	return counts
}

type position struct {
	x int
	y int
}

func (p *position) string() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func createObstablesMap(positions []position) map[string]bool {
	m := map[string]bool{}
	for _, p := range positions {
		m[p.string()] = true
	}
	return m
}

func processInput() (int, position, []position) {
	var n, k int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)
	var queens position
	fmt.Scanf("%d", &queens.y)
	fmt.Scanf("%d", &queens.x)

	obstacles := make([]position, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &obstacles[i].y)
		fmt.Scanf("%d", &obstacles[i].x)
	}
	return n, queens, obstacles
}
