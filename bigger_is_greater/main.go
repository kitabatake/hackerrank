package main

import (
	"fmt"
	"errors"
)

func main() {
	strs := processInput()
	for _, str := range strs {
		rearranged := rearrangeGreaterSmallest(str)
		if rearranged == str {
			fmt.Println("no answer")
		} else {
			fmt.Println(rearranged)
		}
	}
}

func rearrangeGreaterSmallest(str string) string {
	if len(str) < 2 {
		return str
	}

	return traverseFromBackward(str[:len(str) - 1], string(str[len(str) - 1]))
}

// 対象の文字列を2つのパートに分けて、
// 後半部分の文字の中で、前半部分の最後の文字より大きい文字があれば、
// その文字を前半部分の最後の文字と入れ替えて、後半部分をより辞書的に小さい形にソートする。
// 大きい文字がなければ、対象文字を後半部分に移動させて、再帰的に処理していく。
// そうすると、後半部分は自然と降順にソートされるので、後半部分を昇順にソートするときは単純に逆順にすればいい。
func traverseFromBackward(first string, second string) string {
	if len(first) == 0 {
		return second
	}

	// find new initial
	target := rune(first[len(first) - 1])
	greaterIndex, err := findGreaterIndex(target, second)
	if err != nil {
		return traverseFromBackward(first[:len(first) - 1], string(target) + second)
	}
	return first[:len(first) - 1] + rearrageTargetAndSecond(target, greaterIndex, second)
}

func rearrageTargetAndSecond(target rune, replaceIndex int, second string) string {
	newSecond := second[:replaceIndex] + string(target) + second[replaceIndex + 1:]
	ret := string(second[replaceIndex])
	for i := len(second) -1; i >= 0; i-- {
		ret += string(newSecond[i])
	}
	return ret
}

func findGreaterIndex(target rune, str string) (int, error) {
	ret := -1
	for i, c := range str {
		if c > target {
			ret = i
		}
	}
	if ret == -1 {
		return -1, errors.New("not found")
	} else {
		return ret, nil
	}
}

func processInput() []string {
	var n int
	fmt.Scanf("%d", &n)
	ret := make([]string, n)
	for i := 0; i  < n; i++ {
		fmt.Scanf("%s", &ret[i])
	}
	return ret
}