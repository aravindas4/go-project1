package example

import (
	"testing"

	"github.com/matryer/is"
)

func TestHello(t *testing.T) {
	is := is.New(t)

	output := "hello world"
	is.True(output == hello())
}
