package main

import "fmt"

func main() {
	var S string
	var N int
	fmt.Scanf("%s", &S)
	fmt.Scanf("%d", &N)

	var weight rune
	availableWeights := availableWeights(S)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &weight)
		if _, ok := availableWeights[weight]; ok {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func availableWeights(s string) map[rune]bool {
	weights := map[rune]bool{}
	weightBuff := rune(0)
	prevC := rune(-1)
	for _, c := range s {
		c -= 96
		if c == prevC {
			weightBuff += c
		} else {
			weightBuff = c
		}

		weights[weightBuff] = true
		prevC = c
	}
	return weights
}
