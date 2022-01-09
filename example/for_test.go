package example

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestFor(t *testing.T) {
	is := is.New(t)

	output := for_caller(3)
	fmt.Printf(output)

	is.Equal(output, "gogogo")
}

func Test_run_3_times(t *testing.T) {
	is := is.New(t)

	output := run_3_times(27)

	is.Equal(output, 30)
}
