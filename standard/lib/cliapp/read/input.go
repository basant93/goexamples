package main

import (
	"fmt"
)

func main() {
	var address string
	fmt.Println("Enter your address")
	inputs, _ := fmt.Scanf("%s", &address)

	switch inputs {
	case 0:
		fmt.Println("You must enter a address")
	case 1:
		fmt.Println("Address entered:", address)
	}

}
