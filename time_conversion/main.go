package main

import (
	"fmt"
	"strconv"
)

func main() {
	hour, minute, second, ap := processInput()
	militaryHour := calculateMilitaryHour(hour, ap)

	fmt.Printf("%02d:%02d:%02d\n", militaryHour, minute, second)
}

func calculateMilitaryHour(hour int, ap string) int {
	if hour == 12 {
		if ap == "AM" {
			return 0
		} else {
			return 12
		}
	} else {
		if ap == "PM" {
			return hour + 12
		} else {
			return hour
		}
	}
}

func processInput() (int, int, int, string) {
	var input string
	fmt.Scanf("%s", &input)
	hour, _ := strconv.Atoi(input[:2])
	minute, _ := strconv.Atoi(input[3:5])
	second, _ := strconv.Atoi(input[6:8])
	ap := input[8:]
	return hour, minute, second, ap
}
