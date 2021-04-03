package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) == 1 && args[0] == "help" {
		fmt.Println("Usage : ./app 10 20 30")
	}
	var total_wage float32
	for _, v := range args {
		val, _ := strconv.ParseFloat(v, 32)
		total_wage += float32(val)

	}
	fmt.Printf("Total wage for today is %v \n", total_wage)

}
