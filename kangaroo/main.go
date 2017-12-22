package main

import "fmt"

type kangaroo struct {
	x int
	v int
}

func main()  {
	faster, slower := inputProcess()
	ans := judge(faster, slower)

	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func judge(faster kangaroo, slower kangaroo) bool {
	if faster.v == slower.v && faster.x != slower.x {
		return false
	}
	if faster.x > slower.x {
		return false
	}

	for faster.x <= slower.x {
		if faster.x == slower.x {
			return true
		}
		faster.x += faster.v
		slower.x += slower.v
	}

	return false
}

func inputProcess() (faster, slower kangaroo) {
	var x1, x2, v1, v2 int
	fmt.Scanf("%d", &x1)
	fmt.Scanf("%d", &v1)
	fmt.Scanf("%d", &x2)
	fmt.Scanf("%d", &v2)

	if v1 > v2 {
		faster = kangaroo{x: x1, v: v1}
		slower = kangaroo{x: x2, v: v2}
	} else {
		faster = kangaroo{x: x2, v: v2}
		slower = kangaroo{x: x1, v: v1}
	}
	return
}
