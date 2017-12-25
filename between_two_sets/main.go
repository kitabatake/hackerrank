package main

import "fmt"

func main() {
	A, B := inputProcess()
	ans := 0
	for _, n := range extractCommonFactors(B) {
		if isCommonMultiple(n, A) {
			ans++
		}
	}
	fmt.Println(ans)
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func extractCommonFactors(group []int) []int {
	var factors []int
	var isFactor bool

	min, _ := MinMax(group)

	for i := 1; i <= min; i++ {
		isFactor = true
		for _, n := range group {
			if n % i != 0 {
				isFactor = false
			}
		}
		if isFactor {
			factors = append(factors, i)
		}
	}

	return factors
}

func isCommonMultiple(n int, group []int) bool {
	for _, m := range group {
		if n % m != 0 {
			return false
		}
	}
	return true
}

func inputProcess() (a []int, b []int) {
	var n, m int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)

	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}
	return
}