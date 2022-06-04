package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}
type Employee struct {
	Id uint32
}

type FullTimeEmployee struct {
	Employee
	Person
}

func main() {
	employee := FullTimeEmployee{}
	employee.Name = "John"
	employee.Age = 30
	employee.Id = 1
	fmt.Printf("%v\n", employee)
}
