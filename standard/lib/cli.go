package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Current version of GO is " + runtime.Version())
	fmt.Printf("runtime os version %v", runtime.GOOS)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hey %v", text)
}
