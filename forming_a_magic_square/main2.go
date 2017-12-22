package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// この解法のポイント。
// squareの状態は有限（約4億）なので、与えられたsquareから最小コストで作成できるsquareから順にチェックしていき、
// 一度チェックしたものは省くことで、状態の数の爆発を防ぐことができる。（と思う。）

type square [][]int

var N = 3

type queueElement struct {
	square square
	cost   int
}

func main() {
	// process input to square
	s := processInput()

	// init queue it contains square and cost is consumed to create current square.
	queue := list.New()
	queue.PushBack(queueElement{
		square: s, cost: 0,
	})

	cache := map[string]bool{}

	ans := 0
	for e := queue.Front(); e != nil; e = e.Next() {
		qe := e.Value.(queueElement)

		hash := squareToHash(qe.square)
		if _, ok := cache[hash]; ok {
			continue
		}
		cache[hash] = true

		if isMagicSquare(qe.square) {
			ans = qe.cost
			break
		}

		converted := createSquaresWithMinimalCost(qe.square)
		for _, square := range converted {
			queue.PushBack(queueElement{
				square: square,
				cost:   qe.cost + 1,
			})
		}
	}

	fmt.Println(ans)
}

func squareToHash(s square) string {
	hash := ""
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			hash += strconv.Itoa(s[i][j])
		}
	}
	return hash
}

func deepClone(s square) square {
	clone := make(square, len(s))
	for i := 0; i < N; i++ {
		clone[i] = make([]int, N)
		for j := 0; j < N; j++ {
			clone[i][j] = s[i][j]
		}
	}
	return clone
}

func createSquaresWithMinimalCost(s square) []square {
	var squares []square
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			tmp := deepClone(s)
			tmp[i][j] += 1
			if tmp[i][j] < 1 || tmp[i][j] > 9 {
				continue
			}
			squares = append(squares, tmp)

			tmp = deepClone(s)
			tmp[i][j] -= 1
			if tmp[i][j] < 1 || tmp[i][j] > 9 {
				continue
			}
			squares = append(squares, tmp)
		}
	}
	return squares
}

func processInput() square {
	ret := make([][]int, N)

	for i := 0; i < N; i++ {
		ret[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scanf("%d", &ret[i][j])
		}
	}
	return ret
}

// necessary to implements.
// 1. isMagicSquare
// 2. creates Square

func isMagicSquare(s square) bool {
	criteria := s[0][0] + s[0][1] + s[0][2]
	var colSum, rowSum int
	for i := 0; i < N; i++ {
		colSum = 0
		rowSum = 0
		for j := 0; j < N; j++ {
			colSum += s[i][j]
			rowSum += s[j][i]
		}
		if criteria != colSum || criteria != rowSum {
			return false
		}
	}

	var crossSumA, crossSumB int
	for i := 0; i < N; i++ {
		crossSumA += s[i][i]
		crossSumB += s[i][N-i-1]
	}
	if criteria != crossSumB || criteria != crossSumA {
		return false
	}
	return true
}
