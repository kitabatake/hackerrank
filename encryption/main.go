package main

import (
	"fmt"
	"math"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	sqrt := math.Sqrt(float64(len(S)))
	columns := int(math.Ceil(sqrt))
	rows := int(math.Floor(sqrt))

	if columns * rows < len(S) {
		rows++
	}

	arranged := make([]string, rows)
	for i, c := range S {
		arranged[i / columns] = arranged[i / columns] + string(c)
	}

	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			if c < len(arranged[r]) {
				fmt.Printf("%c", arranged[r][c])
			}
		}
		if c != columns - 1{
			fmt.Print(" ")
		}
	}
}
