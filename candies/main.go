// https://www.hackerrank.com/challenges/candies/problem
// ratingが低い子供から最適なキャンディ数が決まっていく（rating:1の子供は絶対1つ）
// なので、ratingが低い子供順にソートして、既に配られている隣接子供からその子供のキャンディ数を順々に決定していく

package main

import (
	"fmt"
	"sort"
)

var N int

type Child struct {
	rating int
	candyNum int
	index int
}

type ByRating []Child
func (a ByRating) Len() int { return len(a) }
func (a ByRating) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRating) Less(i, j int) bool { return a[i].rating < a[j].rating }


func main() {
	fmt.Scanf("%d", &N)
	children := make([]Child, N)
	var rating int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &rating)
		children[i] = Child{rating: rating, index: i}
	}

	byRating := make(ByRating, N)
	copy(byRating, children)
	sort.Sort(byRating)
	ans := 0

	for _, child := range byRating {
		children[child.index].candyNum = calcCandyNum(children, child.index)
		ans += children[child.index].candyNum
	}
	fmt.Println(ans)
}

// 隣接した子供が既にキャンディが配られている場合、ratingから配るキャンディ数を決定する
// ratingが低い子供から呼び出される前提。
func calcCandyNum(children []Child, index int) int {
	dis := []int {-1, 1}
	ret := 1
	for _, di := range dis {
		ci := index + di
		tmp := 1
		if ci < 0 || ci >= len(children) {
			continue
		}
		if children[ci].candyNum != 0 {
			if children[index].rating > children[ci].rating {
				tmp = children[ci].candyNum + 1
			}
		}
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}
