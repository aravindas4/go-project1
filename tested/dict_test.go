package tested

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"some": "thing"}

	t.Run("The key exists in dict", func(t *testing.T) {

		got, _ := dict.Search("some")
		want := "thing"

		assertStrings(t, got, want)
	})

	t.Run("The key does not exist in dict", func(t *testing.T) {

		_, err := dict.Search("some1")
		want := ErrKeyNotFound

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	dict := Dictionary{"key": "value"}
	t.Run("New Word", func(t *testing.T) {
		key := "some"
		value := "thing"
		err := dict.Add("some", "thing")
		assertError(t, err, nil)
		assertAddDefinitions(t, dict, key, value)
	})

	t.Run("Existing Word", func(t *testing.T) {
		key := "key"
		value := "value1"
		err := dict.Add(key, value)

		assertError(t, err, ErrKeyExists)
		assertAddDefinitions(t, dict, key, "value")
	})
}

func TestUpdate(t *testing.T) {
	dict := Dictionary{"key": "value"}

	t.Run("Key exists", func(t *testing.T) {
		key := "key"
		value := "value1"
		err := dict.Update(key, value)
		assertError(t, err, nil)
		assertAddDefinitions(t, dict, key, value)
	})

	t.Run("Key not found", func(t *testing.T) {
		key := "keyq"
		value := "value1"
		err := dict.Update(key, value)
		assertError(t, err, ErrKeyNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q, But Want %q", got, want)
	}
}

func TestDelete(t *testing.T) {
	dict := Dictionary{"key": "value"}

	t.Run("Existing Key", func(t *testing.T) {
		err := dict.Delete("key")
		assertError(t, err, nil)
	})

	t.Run("New Key", func(t *testing.T) {
		err := dict.Delete("key1")
		assertError(t, err, ErrKeyNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got error %q, but want %q", got, want)
	}
}

func assertAddDefinitions(t testing.TB, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search(key)
	if err != nil {
		t.Fatal("Should not get error")
	}
	assertStrings(t, got, value)
}
