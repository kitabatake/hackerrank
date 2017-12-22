package _main

import (
	"fmt"
)


type square [][]int
var N = 3

func main() {
	s := processInput()

	cost := 0
	for ; cost < 72; cost++ {
		if canConvertMagic(s, cost) {
			break;
		}
	}

	fmt.Println(cost)
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

func canConvertMagic(s square, cost int) bool {
	if cost == 0 {
		return isMagicSquare(s)
	}

	var tmp square

	for c := cost; c > 0; c-- {
		for i := 0; i < N; i++ {
			for j:= 0; j < N; j++ {
				tmp = deepClone(s)
				tmp[i][j] += c
				if tmp[i][j] < 1 || tmp[i][j] > 9 {
					continue
				}
				if canConvertMagic(tmp, cost - c) {
					return true
				}
			}
		}
	}

	return false
}

func isMagicSquare(s square) bool {
	return false
	fmt.Printf("%v\n", s)
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
		crossSumB += s[N - i - 1][N - i - 1]
	}
	if criteria != crossSumB || criteria != crossSumA {
		return false
	}

	fmt.Println("it's magic!")
	fmt.Printf("%#v\n", s)
	return true
}