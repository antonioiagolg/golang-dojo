package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Know word", func(t *testing.T) {

		dictionary := Dictionary{"test": "this is just a test"}

		got, err := dictionary.Search("test")
		want := "this is just a test"
		checkString(t, got, want)
		checkErr(t, err, nil)
	})

	t.Run("Unkown word", func(t *testing.T) {
		d := Dictionary{"test": "this is just a test"}

		got, err := d.Search("something")
		want := ""
		wantErr := ErrorUnkonwWord
		checkString(t, got, want)
		checkErr(t, err, wantErr)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add a word", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "this is just a test")
		got, err := dict.Search("test")
		want := "this is just a test"

		if err != nil {
			t.Errorf("Want to find word, got error %v", err)
		}

		checkString(t, got, want)
	})

	t.Run("Add existing word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add("test", "this is just a test")
		checkErr(t, err, nil)
		err = dict.Add("test", "this is just a test")
		checkErr(t, err, ErrorAlreadyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update definition", func(t *testing.T) {
		dict := Dictionary{"test": "test definition"}
		want := "new definition"
		dict.Update("test", want)
		got, err := dict.Search("test")

		checkErr(t, err, nil)
		checkString(t, got, want)
	})

	t.Run("New word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update("test", "test definition")
		checkErr(t, err, ErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete definition", func(t *testing.T) {
		dict := Dictionary{"test": "test definition"}
		dict.Delete("test")
		_, err := dict.Search("test")
		checkErr(t, err, ErrorUnkonwWord)
	})

	t.Run("Delete definition does not exist", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Delete("test")
		checkErr(t, err, ErrorDeleteWordNotExists)
	})

}

func checkString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, but want %q", got, want)
	}
}

func checkErr(t testing.TB, got, want error) {
	t.Helper()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}
