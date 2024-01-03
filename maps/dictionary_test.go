package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"
		assertStrings(t, expected, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("untest")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new_word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing_word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)
		dupErr := dictionary.Add(word, "this is a duplicate")
		assertError(t, dupErr, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update_existing_word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		updatedDef := "this an updated test"

		err := dictionary.Update(word, updatedDef)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, updatedDef)
	})

	t.Run("update_unexisting_word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		updatedDef := "this an updated test"

		err := dictionary.Update(word, updatedDef)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "word"
	dictionary := Dictionary{word: "definition of the word"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertError(t, err, ErrNotFound)

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	def, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	assertStrings(t, definition, def)

}

func assertStrings(t testing.TB, expected, got string) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func assertError(t testing.TB, got, err error) {
	t.Helper()
	if err != got {
		t.Errorf("expected %q but got %q.", err, got)
	}

}
