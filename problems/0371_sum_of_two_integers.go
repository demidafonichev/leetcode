package main

import (
	"fmt"
	"strconv"
)

func getSum(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	res, extra := 0, false

	for i := 0; a != 0 || b != 0; i++ { // i++
		ai := (a & 1) != 0
		bi := (b & 1) != 0
		fmt.Println("ai, bi: ", ai, bi)
		// 1010
		// 0000
		if ai && bi || ai && extra || bi && extra {
			res |= 1 << i
			extra = true
		} else if ai || bi || extra {
			res |= 1 << i
			extra = false
		}
		a >>= 1
		b >>= 1
		fmt.Println("a, b: ", a, b)
	}
	fmt.Println(strconv.FormatInt(int64(res), 2))
	return res
}

func main() {
	// fmt.Println("1 + 2 =", getSum(1, 2))
	// fmt.Println("3 + 4 =", getSum(3, 4))
	fmt.Println("0 + 10 =", getSum(0, 10))
}
