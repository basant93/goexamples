package profiling

import (
	"fmt"
	"testing"
)

func messageReader(greeter, greet string) string {
	message := fmt.Sprintf("%v, %v", greeter, greet)
	return message
}
