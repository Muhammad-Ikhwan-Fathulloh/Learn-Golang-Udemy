package main

import "fmt"

func main() {
	const name = "Muhammad Ikhwan Fathulloh"
	fmt.Println(name)

	firstName := "Muhammad"
	fmt.Println(firstName)

	const friendName = "Ikhwan"
	fmt.Println(friendName)

	const age = 21
	fmt.Println(age)

	const (
		first_Name  string = "Muhammad"
		middle_Name        = "Ikhwan"
		last_Name          = "Fathulloh"
	)

	fmt.Println(first_Name)
	fmt.Println(middle_Name)
	fmt.Println(last_Name)
}
