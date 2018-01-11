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

type Children []Child

func (c Children) Len() int {
	return len(c)
}

func (c Children) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Children) Less(i, j int) bool {
	return c[i].rating < c[j].rating
}

func main() {
	fmt.Scanf("%d", &N)
	children := make(Children, N)
	sortedChildren := make(Children, N)
	var rating int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &rating)
		children[i] = Child{rating: rating, index: i}
		sortedChildren[i] = Child{rating: rating, index: i}
	}

	sort.Sort(sortedChildren)
	ans := 0
	for _, child := range sortedChildren {
		children[child.index].candyNum = calcCandyNum(children, child.index)
		ans += children[child.index].candyNum
	}
	fmt.Println(ans)
}

// 隣接した子供が既にキャンディが配られている場合、ratingから配るキャンディ数を決定する
// ratingが低い子供から呼び出される前提。
func calcCandyNum(children Children, index int) int {
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
