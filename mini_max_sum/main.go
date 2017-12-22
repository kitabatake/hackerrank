package main

import "fmt"

var N = 5

func main() {
	integers := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &integers[i])
	}

	var min, max, sum uint32

	for pass := 0; pass < N; pass ++ {
		sum = 0
		for i := 0; i < N; i++ {
			if i == pass {
				continue
			}
			sum += uint32(integers[i])
		}
		if sum > max {
			max = sum
		}
		if min == 0 || sum < min {
			min = sum
		}
	}

	fmt.Println(min, max)
}
