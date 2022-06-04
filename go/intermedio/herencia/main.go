package main

import "fmt"

type PrintInfo interface {
	GetInfo() string
}

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

func (fullTimeEmployee FullTimeEmployee) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Id: %d", fullTimeEmployee.Name, fullTimeEmployee.Age, fullTimeEmployee.Id)
}

func GetInfo(employee PrintInfo) string {
	return employee.GetInfo()
}

func main() {
	employee := FullTimeEmployee{}
	employee.Name = "John"
	employee.Age = 30
	employee.Id = 1
	fmt.Printf("%v\n", employee)
	fmt.Printf("%v\n", GetInfo(employee))
}
