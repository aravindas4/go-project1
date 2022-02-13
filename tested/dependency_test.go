package tested

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Aravinda")

	got := buffer.String()
	want := "Hello, Aravinda"

	if got != want {
		t.Errorf("Got %q, But want %q", got, want)
	}
}
