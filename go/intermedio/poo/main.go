package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func NewEmployee(id *int, name *string) *Employee {
	return &Employee{
		id:   *id,
		name: *name,
	}
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
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// 2
	e2 := Employee{
		id:   1,
		name: "Johnny",
	}
	fmt.Printf("%v\n", e2)

	// 3
	e3 := new(Employee)
	e3.id = 1
	e3.name = "Johnny"
	fmt.Printf("%v\n", e3)

	// 4
	name := "Johnny"
	id := 24
	e4 := NewEmployee(&id, &name)
	fmt.Printf("%v\n", e4)
}
