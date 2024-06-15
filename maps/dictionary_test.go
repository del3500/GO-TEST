package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		_, got := dictionary.Search("test")
		assertError(t, got, nil)
	})

	t.Run("missing key", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("Add unique word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "sample word"
		definition := "sample definition"

		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Add duplicate word", func(t *testing.T) {
		word := "test"
		definition := "test def"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrDuplicate)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Update("test2", "new definition")

		assertError(t, err, ErrNotFound)
	})

}

func TestDelete(t *testing.T) {

	t.Run("delete known word", func(t *testing.T) {
		word := "test"
		definition := "this is just test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)
		assertDelete(t, dictionary, word)

	})

	t.Run("delete unknown word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertDelete(t *testing.T, d Dictionary, word string) {
	t.Helper()
	_, err := d.Search(word)
	if err == nil {
		t.Fatal("word should already be deleted")
	}
}

func assertDefinition(t *testing.T, d Dictionary, word, definition string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("Should find the added word", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, got, err error) {
	t.Helper()
	if got != err {
		t.Errorf("got %q, want %q", got, err)
	}
}
