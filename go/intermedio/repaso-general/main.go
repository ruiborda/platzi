package main

import (
	"fmt"
	"strconv"
	"time"
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

	// apuntadores
	g := 25
	fmt.Println(g)
	p := &g
	fmt.Println(p)
	fmt.Println(*p)
	
	// goRoutines
	c := make(chan bool)
	go doSomething(c)
	<-c
}

func doSomething(c chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Printf("Done\n")
	c <- true
}
