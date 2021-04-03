package hello

import "fmt"

func SayHello() {
	fmt.Println("Hello Go")
}

//go build cli/buildInstall/build/hello/hello.go
//This package is not a package. go build the source file and threw the result.
//If we run go build on main package, it saves the artifacts on the disk
//go build cli/buildInstall/build/build.go
//Run : ./build
//by default libraries are build and thrown away and binaries are build and artifacts are saved
//if you want to save use : -o flag
// -i flag : install the dependencies
