package gocover

import "testing"

func TestAdder(t *testing.T) {
	if adder(2, 3) != 5 {
		t.Fail()
	}
}

//Run test with coverage report
//go test -cover github.com/basant93/goexamples/gotest/gocover/
// to generate a cover profile file
//go test -coverprofile cover.out github.com/basant93/goexamples/gotest/gocover/
//to visualize the coverage report
//go tool cover -html=cover.out
//analyse the number of times the statement is covered
//go test -covermode count -coverprofile cover.out github.com/basant93/goexamples/gotest/gocover/
//scale of the coverage
