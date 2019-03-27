package main

import (
	"errors"
	"fmt"
)

// SimpleFunction is a simple function that takes in two strings
// concatenates them into one string and returns the result
func SimpleFunction(firstName string, lastName string) string {
	return firstName + " " + lastName
}

// SimpleFuncWithError takes in the same two strings as before
// but checks to see if the first name is greater than 1 character
// in length. If it isn't it returns an error
func SimpleFuncWithError(firstName string, lastName string) (string, error) {
	if len(firstName) <= 1 {
		return "", errors.New("First Name Must Be At Least 2 letters long")
	}

	// if the first name passes the check, we return the
	// concatenated names and a nil error value
	return firstName + " " + lastName, nil
}

// addOne is a function that returns
// a function which in turn returns an int value
func addOne() func() int {
	var x int
	return func() int {
		x++
		return x + 1
	}
}

func main() {
	fmt.Println("Go Functions Tutorial")

	// Test our first simple function which just returns
	// one value which is our concatenated name
	fullName := SimpleFunction("Elliot", "Forbes")
	fmt.Println(fullName)

	// Test our second function which returns 2 values,
	// our name and an error if there is any
	fullName2, err := SimpleFuncWithError("Elliot", "Forbes")
	if err != nil {
		fmt.Println("We Handle the error if there is one")
	}
	fmt.Println(fullName2)

	// Let's test our our function that returns a
	// function which returns an int.
	myFunc := addOne()
	fmt.Println(myFunc()) // prints out 2
	fmt.Println(myFunc()) // prints out 3
	fmt.Println(myFunc()) // prints out 4
}
