package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	bigIntAns := factorial(N)
	fmt.Println(bigIntAns.String())
}

func factorial(n int) *big.Int {
	ret := big.NewInt(1)
	for i := int64(n); i > 1; i-- {
		ret.Mul(ret, big.NewInt(i))
	}
	return ret
}
