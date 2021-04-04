package main

import (
	"fmt"
)

func main() {
	var first int
	var second int

	fmt.Scanf("%d %d", &first, &second)
	fmt.Println("Sum is: ", first+second)
	var name string
	fmt.Sscanf("Hello Go", "Hello %s", &name)
	fmt.Println(name)
	var source string
	var destination string
	fmt.Scanln(&source, &destination)
	fmt.Println(source, destination)

	output := fmt.Sprintf("%s -> %s", "hello", "go")
	fmt.Println(output)
}
