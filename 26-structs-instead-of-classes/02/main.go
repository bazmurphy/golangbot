package main

import "oop/employee"

// The program we wrote above looks alright but there is a subtle issue in it.
// Let’s see what happens when we define the employee struct with zero values.
// Replace the contents of main.go with the following code
// The only change we have made is creating a zero value Employee in line no. 6. This program will output
//   has 0 leaves remaining
// As you can see, the variable created with the zero value of Employee is unusable.
// It doesn’t have a valid first name, last name and also doesn’t have valid leave details.

// func main() {
// 	var e employee.Employee
// 	e.LeavesRemaining()
// }

//   has 0 leaves remaining

func main() {
	// We have created a new employee by passing the required parameters to the New function.
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}

// Sam Adolf has 10 leaves remaining

// Although Go doesn’t support classes, structs can effectively be used instead of classes and methods of signature New(parameters) can be used in the place of constructors
