package main

import "fmt"

func main() {
	k, numbers := processInput()
	fmt.Println(solve(numbers, k))
}

func solve(numbers []int, k int) int {
	remaindersCount := map[int]int{}
	for _, n := range numbers {
		remaindersCount[n % k]++
	}

	ret := 0
	if remaindersCount[0] > 0 {
		ret++
	}
	if k % 2 == 0 && remaindersCount[k/2] > 0 {
		ret++
	}

	for i := 1; float64(i) < float64(k)/2; i++ {
		if remaindersCount[i] > remaindersCount[k - i] {
			ret += remaindersCount[i]
		} else {
			ret += remaindersCount[k - i]
		}
	}
	return ret
}

func processInput() (int, []int) {
	var k, n int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}
	return k, numbers
}