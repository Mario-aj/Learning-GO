package main

import "fmt"

func main () {
	const firstName = "Mario"
	const middleName = "Alfredo"
	const lastName = "Jorge"

	fullName := fmt.Sprintf("%s %s %s", firstName, middleName, lastName)

	fmt.Printf(fullName)
}