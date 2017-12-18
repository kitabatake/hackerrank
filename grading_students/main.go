package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	var grade int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &grade)
		fmt.Println(roundGrade(grade))
	}
}

func roundGrade(original int) int {
	if original < 38 {
		return original
	}
	if original % 5 >= 3 {
		return original + (5 - original % 5)
	}
	return original
}
