//go help testflag
//
package profiling

import "testing"

func TestMessageReader(t *testing.T) {
	if messageReader("Hello", "GO") != "Hello,Go" {
		t.Fail()
	}
}

//-blockprofile gives you about number of goroutines waiting
//-cpuprofile, -memprofile
//-trace : write an execution trace to the specified file before testing
//generate a memory profile
//go test -memprofile mem.out -benchmem github.com/basant93/goexamples/gotest/gocover/
//to visualize memprofile install graphviz
//brew install graphviz
//genrate a memprofile with memprofilerate as 1
//go test -memprofile mem.out -memprofilerate 1 github.com/basant93/goexamples/gotest/gocover/
//go tool pprof -web gocover.test mem.out
//for cpuprofiling
//go test -cpuprofile cpu.out -count 10000 github.com/basant93/goexamples/gotest/gocover/
//go tool pprof -web gocover.test cpu.out
