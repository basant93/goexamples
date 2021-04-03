package main

import (
	"fmt"
	"github.com/basant93/gobyexamples/cli/buildInstall/hello"
)

//go help run
// compiles and run the main package comprising go source code files. It does not leave build artifacts.
//go help build

func main() {
	fmt.Println("Hello")
	hello.SayHello()
}
