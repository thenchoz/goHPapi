package goHPapi

import (
	"context"
	"testing"
)

func Test_Fetch_Book(t *testing.T) {
	api := New()
	_, err := api.FetchBook(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Books(t *testing.T) {
	api := New()
	_, err := api.FetchBooks(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Character(t *testing.T) {
	api := New()
	_, err := api.FetchCharacter(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Characters(t *testing.T) {
	api := New()
	_, err := api.FetchCharacters(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_House(t *testing.T) {
	api := New()
	_, err := api.FetchHouse(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Houses(t *testing.T) {
	api := New()
	_, err := api.FetchHouses(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Spell(t *testing.T) {
	api := New()
	_, err := api.FetchSpell(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Spells(t *testing.T) {
	api := New()
	_, err := api.FetchSpells(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Quote(t *testing.T) {
	api := New()
	_, err := api.FetchQuote(context.Background())

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Fetch_Param(t *testing.T) {
	testMax := 5

	api := New()
	api.SetMax(testMax)

	books, err := api.FetchBooks(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	if len(books) != testMax {
		t.Fatalf("Wrong number of results")
	}
}
