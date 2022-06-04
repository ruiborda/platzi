package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Printf("%d\n", x)
	fmt.Printf("%d\n", y)

	parseInt, err := strconv.ParseInt("789", 0, 64)
	if err == nil {
		fmt.Printf("%d\n", parseInt)
	} else {
		fmt.Printf("%s\n", err.Error())
	}

	m := make([]int, 5, 6)
	m2 := map[string]int{}
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", m2)

	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Printf("%d: %d\n", i, v)
	}

}
