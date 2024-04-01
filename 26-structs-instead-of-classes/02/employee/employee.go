package employee

import (
	"fmt"
)

// In other OOP languages like java, this problem can be solved by using constructors.
// A valid object can be created by using a parameterized constructor.

// Go doesn’t support constructors.
// If the zero value of a type is not usable,
// it is the job of the programmer to unexport the type to prevent access from other packages
// and also to provide a function named NewT(parameters) which initializes the type T with the required values.
// It is a convention in Go to name a function that creates a value of type T to NewT(parameters).
// This will act as a constructor.
// If the package defines only one type, then it’s a convention
// in Go to name this function just New(parameters) instead of NewT(parameters).

// Let’s make changes to the program we wrote so that every time an employee is created, it is usable.

// We have made the starting letter e of Employee struct to lower case
// By doing so we have successfully unexported the employee struct and prevented access from other packages
// It’s a good practice to make all fields of an unexported struct to be unexported too unless there is a specific need to export them
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

// Now since employee is unexported, it’s not possible to create values of type Employee from other packages
// Hence we are providing an exported New function which takes the required parameters as input and returns a newly created employee.
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
