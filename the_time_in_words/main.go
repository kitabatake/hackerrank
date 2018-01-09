package main

import (
	"fmt"
)

var numeralsWords = map[int]string {
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
	10: "ten",
	11: "eleven",
	12: "twenty",
	13: "thirteen",
	14: "fourteen",
	15: "quarter",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	21: "twenty one",
	22: "twenty two",
	23: "twenty three",
	24: "twenty four",
	25: "twenty five",
	26: "twenty six",
	27: "twenty seven",
	28: "twenty eight",
	29: "twenty nine",
	30: "half",
}

func main() {
	var H, M int
	fmt.Scanf("%d", &H)
	fmt.Scanf("%d", &M)

	words := getMinutes(M)
	if words == "" {
		words = getHours(H, M)
	} else {
		words += " " + getHours(H, M)
	}

	fmt.Println(words)
}

func getHours(H int, M int) string {
	if M > 30 {
		H++
	}
	if H == 13 {
		H = 1
	}

	words := numeralsWords[H]

	if M == 0 {
		words += " o' clock"
	}
	return words
}

func getMinutes(M int) string {
	if M == 0 {
		return ""
	}

	var w, mc, last string
	if M <= 30 {
		w = numeralsWords[M]
		mc = getMinutesClause(w)
		last = "past"
	} else {
		w = numeralsWords[60 - M]
		mc = getMinutesClause(w)
		last = "to"
	}

	if mc == "" {
		return w + " " + last
	} else {
		return w + " " + getMinutesClause(w) + " " + last
	}
}

func getMinutesClause(m string) string {
	if m == "one" {
		return "minute"
	} else if m == "quarter" || m == "half" {
		return ""
	} else {
		return "minutes"
	}
}


