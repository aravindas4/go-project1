package example

import (
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func Test_switch_caller(t *testing.T) {
	is := is.New(t)
	var output interface{}

	// bool
	output = switch_caller(true)
	is.Equal(output, reflect.TypeOf(false))

	// integer
	output = switch_caller(1)
	is.Equal(output, reflect.TypeOf(2))

	// String
	output = switch_caller("aravinda")
	is.Equal(output, reflect.TypeOf("some"))

	//float
	output = switch_caller(1.24)
	is.Equal(output, reflect.TypeOf(1.2323))
}
