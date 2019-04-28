package main

import (
	"fmt"
	"strconv"
)

func calk() int {
	var arg string
	fmt.Scan(&arg)
	var res int
	for i := 0; i < len(arg); i++ {
		i, e := strconv.Atoi(string([]rune(arg)[i : i+1]))
		if e != nil {
			panic(e)
		}
		res += i
	}
	return res
}

func main() {
	fmt.Println(calk())
}
