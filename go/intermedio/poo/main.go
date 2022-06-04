package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (employee *Employee) setId(id int) {
	employee.id = id
}

func (employee *Employee) setName(name string) {
	employee.name = name
}

func (employee *Employee) getId() int {
	return employee.id
}

func (employee *Employee) getName() string {
	return employee.name
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)
	e.setId(1)
	e.setName("Jhon")

	fmt.Printf("%v\n", e)
}
