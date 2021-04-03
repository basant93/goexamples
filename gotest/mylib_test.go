package mylib

import "testing"

//identify resource as test add _test suffix in file name
//if you run go build it exclueds test files by deafault
//if we run go test, it includes test files in build
//save anciliary data in "tfestdata" dir, it is ignored by test

func Test_BasicsCheck(t *testing.T) {
	t.Parallel() //go knows these test do not rely on each other so it can run in parallel.
	t.Run("failed GO add check", func(t *testing.T) {
		if 1+1 != 2 {
			t.Fail()
		}

	})

	t.Run("Go concat check", func(t *testing.T) {
		if "Hello"+"GO" != "HelloGO" {
			t.Fail()
		}
	})
}

func Test_MoreBasics(t *testing.T) {
	t.Run("Add div test", func(t *testing.T) {
		if 21/3 != 7 {
			t.Fail()
		}
	})
}

//go test gotest/...
//go test ./...
//execute all packages including child packages
//go test gotest/mylib_test.go
// -args flag pass the remainder of the command line to test binary
//go help testflag
/*
-cpu 1,2,4
Specify a list of GOMAXPROCS values for which the tests or
benchmarks should be executed. The default is the current value
of GOMAXPROCS.
*/
//List the test
//go test -list Basics ./...
/*
-list regexp
            List tests, benchmarks, or examples matching the regular expression.
            No tests, benchmarks or examples will be run. This will only
			list top-level tests. No subtest or subbenchmarks will be shown.

 -count n
            Run each test and benchmark n times (default 1).
            If -cpu is set, run n times for each GOMAXPROCS value.
			Examples are always run once.
 -cover
            Enable coverage analysis.
            Note that because coverage works by annotating the source
            code before compilation, compilation and test failures with
            coverage enabled may report line numbers that don't correspond
			to the original sources.

*/
