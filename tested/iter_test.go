package tested

import "testing"

func TestIter(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("Expected %q, But got %q", expected, repeated)
	}
}
