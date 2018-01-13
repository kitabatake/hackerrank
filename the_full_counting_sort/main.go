package main

import (
	"fmt"
	"sort"
)

var N int

type Element struct {
	number int
	str string
	order int
}

type ElementsByNumber []Element
func (a ElementsByNumber) Len() int { return len(a) }
func (a ElementsByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ElementsByNumber) Less(i, j int) bool { return a[i].number < a[j].number }

func main() {
	fmt.Scanf("%d", &N)
	elements := make(ElementsByNumber, N)
	var str string
	var number int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &number)
		fmt.Scanf("%s", &str)
		if i < N / 2 {
			str = "-"
		}
		elements[i] = Element{number:number, str:str, order: i}
	}
	sort.Stable(elements)

	for i, ele := range elements {
		fmt.Print(ele.str)
		if i == N - 1 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}
