package main

import (
	"fmt"
	"strconv"
)

/* Parent class */
type Person struct {
	name string
	age  int
}

/* Parent class */
type Employee struct {
	id int
}

/* Child class*/
type FullTimeEmployee struct {
	Person
	Employee
	salaryRate int
} /* Child class*/
type TempEmployee struct {
	Person
	Employee
	beginDate int
	endDate   int
}

/*Polimofirsm is trought interface*/
type PrintInfo interface {
	getMessage() string
}

func (e *Employee) setId(id int) {
	e.id = id
}

func (p *Person) setName(name string) {
	p.name = name
}
func (tempEmployee *TempEmployee) setDates(endDate int, beginDate int) {
	tempEmployee.beginDate = beginDate
	tempEmployee.endDate = endDate
}

func (fullTimeEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee is "
}
func (tempEmployee TempEmployee) getMessage() string {
	return "Temp time employee end date is " + strconv.Itoa(tempEmployee.endDate)
}

func getMessage(p PrintInfo) string {
	return p.getMessage()
}

func main() {
	employee := FullTimeEmployee{}
	tempEmployee := TempEmployee{}
	fmt.Println(employee)
	employee.setId(2)
	employee.setName("Not Pepito")
	tempEmployee.setId(22)
	tempEmployee.setName("Is Pepito")
	tempEmployee.setDates(202405, 200310)
	fmt.Println(employee.getMessage(), employee.name)
	fmt.Println(getMessage(tempEmployee))
	fmt.Println(getMessage(employee), employee.name)
}
