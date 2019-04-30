package main

import (
	"fmt"
)

func calk() int {
	// refer to https://qiita.com/koyo-miyamura/items/17ddf9aa70fb1479e6eb
	var (
		a, b, c, fee, result int
	)
	fmt.Scan(&a, &b, &c, &fee)
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				if i*500+j*100+k*50 == fee {
					result++
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(calk())
}
