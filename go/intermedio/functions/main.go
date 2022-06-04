package main

import "fmt"

func printNames(names ...string) {
	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
}

func getValues(x int) (first int, second int) {
	first = x
	second = x * 2
	return
}

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Printf("%d\n", y)
	// goRoutine
	c := make(chan bool)
	go func() {
		fmt.Printf("%d\n", x)
		c <- true
	}()
	<-c

	// variadic
	printNames("John", "Jane", "Jack")

	// named return values
	first, second := getValues(5)
	fmt.Printf("%d %d\n", first, second)
}
