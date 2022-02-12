package tested

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Aravinda", "")
		want := "Hello, Aravinda"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Hello, world', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Kannada", func(t *testing.T) {
		got := Hello("Aravinda", "Kannada")
		want := "Aravinda Avare,"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Tulu", func(t *testing.T) {
		got := Hello("Aravinda", "Tulu")
		want := "Aravinda Mere,"

		assertCorrectMessage(t, got, want)
	})
}
