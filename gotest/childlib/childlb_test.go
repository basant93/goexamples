package childlib

import (
	"testing"
)

func Test_Subtraction(t *testing.T) {
	t.Run("Go sub", func(t *testing.T) {
		if 5-2 != 3 {
			t.Fail()
		}
	})
}

//Run : go test gotest/childlib/childlb_test.go
