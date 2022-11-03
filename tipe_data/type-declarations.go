package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpIkhwan NoKTP = "1717171717171717"
	var marriedStatus Married = true

	fmt.Println(noKtpIkhwan)
	fmt.Println(marriedStatus)
}
