package main

import "fmt"

func main() {
	var s,t,a,b,m,n int
	fmt.Scanf("%d", &s)
	fmt.Scanf("%d", &t)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)

	var d int
	apples := 0
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &d)
		if inRange(s, t, a + d) {
			apples++
		}
	}
	fmt.Println(apples)

	oranges := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &d)
		if inRange(s, t, b + d) {
			oranges++
		}
	}
	fmt.Println(oranges)
}

func inRange(s int, e int, p int) bool {
	return p >= s && p <= e
}
