package main

import (
	"flag"
	"fmt"
)

var (
	name           = flag.String("name", "", "Customer name")
	gender         = flag.String("gender", "", "Customer Gender")
	mobile         = flag.String("mobile", "", "Customer mobile number")
	billAmount     = flag.Float64("billAmount", 0.0, "Bill amount to be payed")
	itemsPurchased = flag.String("itemsPurchased", "", "Items purchased by customer")
)

func main() {
	flag.Parse()
	check := true
	switch {
	case *name == "":
		fmt.Println("Please enter customer name")
		check = false
	case *gender == "":
		fmt.Println("Please enter customer gender")
		check = false
	case *mobile == "":
		fmt.Println("Please eter customer mobile number")
		check = false
	case *billAmount == 0.0:
		fmt.Println("Please enter bill amount")
		check = false
	case *itemsPurchased == "":
		fmt.Println("Please enter the items purchased by customer")
		check = false

	}
	if !check {
		fmt.Println("Please enter the required deatails")
	}
	fmt.Println("Customer Info:")
	fmt.Printf("Name: %v \n Gender: %v \n Mobile: %v \n Bill Amount: %v \n Items Purchased: %v \n", *name, *gender, *mobile, *billAmount, *itemsPurchased)

}
